package repository

import (
	"context"
	"log/slog"
	"slices"
	"strings"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/ent/categories"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sites"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sitesuggestions"
	"github.com/go-errors/errors"
)

type SitesRepository interface {
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*ent.Sites, error)
	CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*ent.SiteSuggestions, error)
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
	GetAllSiteSuggestions(ctx context.Context, params genapi.ListSiteSuggestionsParams) ([]genapi.SiteSuggestion, error)
	GetSiteByDomain(ctx context.Context, domain string) (*ent.Sites, error)
	GetSiteBlocksByID(ctx context.Context, id int) (map[int][]*ent.Blocks, error)
	GetSiteByID(ctx context.Context, id int) (*ent.Sites, error)
	GetSiteSuggestionByID(ctx context.Context, id int) (*ent.SiteSuggestions, error)
}

type sitesRepositoryImpl struct {
	db *ent.Client
}

func (s *sitesRepositoryImpl) CreateSite(ctx context.Context, req *genapi.SiteInput) (*ent.Sites, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		slog.Error("failed to start transaction", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	// find existing categories
	categories, err := tx.Categories.Query().Where(
		categories.NameIn(req.Categories...),
	).All(ctx)
	if err != nil {
		slog.Error("failed to get categories", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	// check if all categories are found
	if len(categories) != len(req.Categories) {
		slog.Warn("categories not found, will attempt to create automatically", "categories", req.Categories)

		// find the missing categories
		missingCategories := []string{}
		for _, category := range req.Categories {
			found := false
			for _, c := range categories {
				if c.Name == category {
					found = true
					break
				}
			}
			if !found {
				missingCategories = append(missingCategories, category)
			}
		}

		// create the missing categories
		newCategories, err := tx.Categories.MapCreateBulk(missingCategories, func(cc *ent.CategoriesCreate, i int) {
			cc.SetName(missingCategories[i])
		}).Save(ctx)
		if err != nil {
			slog.Error("failed to create missing categories", "error", err)
			return nil, rollback(tx, errors.Wrap(err, 0))
		}

		// append the new categories to the existing categories
		categories = append(categories, newCategories...)
	}

	// create the site
	var ping_url string
	if req.PingURL.Set {
		ping_url = req.PingURL.Value
	} else {
		ping_url = req.Domain
	}
	site, err := tx.Sites.Create().
		SetDomain(req.Domain).
		SetPingURL(ping_url).
		AddCategories(categories...).
		Save(ctx)
	if err != nil {
		slog.Error("failed to create site", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	return site, tx.Commit()
}

func (s *sitesRepositoryImpl) CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*ent.SiteSuggestions, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		slog.Error("failed to start transaction", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	// create the site
	var ping_url string
	if req.PingURL.Set {
		ping_url = req.PingURL.Value
	} else {
		ping_url = req.Domain
	}

	categories := strings.Join(req.Categories, ",")

	suggestion, err := tx.SiteSuggestions.Create().
		SetDomain(req.Domain).
		SetPingURL(ping_url).
		SetCategories(categories).
		SetReason(req.Reason).
		Save(ctx)
	if err != nil {
		slog.Error("failed to create site suggestion", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	return suggestion, tx.Commit()
}

func (s *sitesRepositoryImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		slog.Error("failed to start transaction", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	// get all sites in the database
	query := tx.Sites.Query().WithCategories().WithBlocks()
	if params.Category.Set {
		catNames := strings.Split(params.Category.Or(""), ",")
		for i, cat := range catNames {
			catNames[i] = strings.TrimSpace(cat)
		}

		predicates := make([]predicate.Categories, len(catNames))
		for i, cat := range catNames {
			predicates[i] = categories.NameEQ(cat)
		}
		query = query.Where(
			sites.HasCategoriesWith(categories.NameIn(catNames...)),
		)
	}

	if params.Order.Set {
		if params.Order.Value == genapi.ListSitesOrderAsc {
			query = query.Order(ent.Asc(params.Sort.Or("id")))
		} else {
			query = query.Order(ent.Desc(params.Sort.Or("id")))
		}
	}

	dbSites, err := query.Limit(params.Limit.Or(50)).Offset(params.Offset.Or(0)).All(ctx)
	if err != nil {
		slog.Error("failed to get sites", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	var filteredSites []*ent.Sites
	if params.Category.Set {
		catNames := strings.Split(params.Category.Or(""), ",")
		for i, cat := range catNames {
			catNames[i] = strings.TrimSpace(cat)
		}

		for _, s := range dbSites {
			cats, err := s.QueryCategories().Where(categories.NameIn(catNames...)).All(ctx)
			if err != nil {
				slog.Error("failed to get categories", "error", err)
				return nil, rollback(tx, errors.Wrap(err, 0))
			}
			if len(cats) == len(catNames) {
				filteredSites = append(filteredSites, s)
			}
		}
	} else {
		filteredSites = dbSites
	}

	// create a map of sites
	sites := map[string]*genapi.Site{}
	for _, site := range filteredSites {
		// convert the categories to a slice of strings
		c := make([]string, len(site.Edges.Categories))
		for i, category := range site.Edges.Categories {
			c[i] = category.Name
		}

		// map domain to site struct
		sites[site.Domain] = &genapi.Site{
			ID:         site.ID,
			Domain:     site.Domain,
			PingURL:    site.PingURL,
			Categories: c,
			CreatedAt:  site.CreatedAt,
			UpdatedAt:  site.UpdatedAt,
		}
	}

	for _, site := range filteredSites {
		for _, block := range site.Edges.Blocks {
			if _, ok := sites[site.Domain]; ok {
				// Update the existing site
				// Add the block and unblock reports
				if block.Blocked {
					sites[site.Domain].BlockReports += 1
				} else {
					sites[site.Domain].UnblockReports += 1
				}

				// Update the last reported at
				if sites[site.Domain].LastReportedAt.Before(block.LastReportedAt) {
					sites[site.Domain].LastReportedAt = block.LastReportedAt
				}
			}
		}
	}

	// convert the map to a slice
	result := make([]genapi.Site, 0, len(sites))
	for _, site := range sites {
		result = append(result, *site)
	}

	return result, nil
}

func (s *sitesRepositoryImpl) GetAllSiteSuggestions(ctx context.Context, params genapi.ListSiteSuggestionsParams) ([]genapi.SiteSuggestion, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		slog.Error("failed to start transaction", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	// get all sites in the database
	query := tx.SiteSuggestions.Query()

	if params.Order.Set {
		if params.Order.Value == genapi.ListSiteSuggestionsOrderAsc {
			query = query.Order(ent.Asc(params.Sort.Or("id")))
		} else {
			query = query.Order(ent.Desc(params.Sort.Or("id")))
		}
	}

	dbSites, err := query.Limit(params.Limit.Or(50)).Offset(params.Offset.Or(0)).All(ctx)

	if err != nil {
		slog.Error("failed to get site suggestions", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	var filteredSiteSuggestions []*ent.SiteSuggestions
	if params.Category.Set {
		catNames := strings.Split(params.Category.Or(""), ",")
		for i, cat := range catNames {
			catNames[i] = strings.TrimSpace(cat)
		}

		for _, s := range dbSites {
			siteCats := strings.Split(s.Categories, ",")
			addSite := true
			for _, cat := range catNames {
				if !slices.Contains(siteCats, cat) {
					addSite = false
					break
				}
			}
			if addSite {
				filteredSiteSuggestions = append(filteredSiteSuggestions, s)
			}
		}
	} else {
		filteredSiteSuggestions = dbSites
	}

	// create a map of sites
	suggestions := map[string]*genapi.SiteSuggestion{}
	for _, suggestion := range filteredSiteSuggestions {
		// map domain to site struct
		suggestions[suggestion.Domain] = &genapi.SiteSuggestion{
			ID:            suggestion.ID,
			Domain:        suggestion.Domain,
			PingURL:       suggestion.PingURL,
			Categories:    strings.Split(suggestion.Categories, ","),
			Reason:        suggestion.Reason,
			Status:        genapi.SiteSuggestionStatus(suggestion.Status),
			ResolveReason: suggestion.ResolveReason,
			ResolvedAt:    suggestion.ResolvedAt,
			CreatedAt:     suggestion.CreatedAt,
			UpdatedAt:     suggestion.UpdatedAt,
		}
	}

	// convert the map to a slice
	result := make([]genapi.SiteSuggestion, 0, len(suggestions))
	for _, suggestion := range suggestions {
		result = append(result, *suggestion)
	}

	return result, nil
}

func (s *sitesRepositoryImpl) GetSiteByDomain(ctx context.Context, domain string) (*ent.Sites, error) {
	db := s.getDb(ctx)

	return db.Sites.Query().
		Where(sites.DomainEQ(domain)).
		First(ctx)
}

func (s *sitesRepositoryImpl) GetSiteBlocksByID(ctx context.Context, id int) (map[int][]*ent.Blocks, error) {
	db := s.getDb(ctx)

	blocks, err := db.Blocks.Query().
		Where(blocks.HasSiteWith(sites.ID(id))).
		WithSite().
		WithIsp().
		All(ctx)
	if err != nil {
		slog.Error("failed to get blocks", "error", err)
		return nil, errors.Wrap(err, 0)
	}
	if len(blocks) == 0 {
		return nil, nil
	}

	isps := make(map[int][]*ent.Blocks)
	// Ignore nilderef, ISP ID is a mandatory field so it should always exist
	for _, block := range blocks {
		isps[block.IspID] = append(isps[block.IspID], block)
	}

	return isps, nil
}

func (s *sitesRepositoryImpl) GetSiteByID(ctx context.Context, id int) (*ent.Sites, error) {
	db := s.getDb(ctx)

	return db.Sites.Query().WithCategories().Where(sites.ID(id)).First(ctx)
}

func (s *sitesRepositoryImpl) GetSiteSuggestionByID(ctx context.Context, id int) (*ent.SiteSuggestions, error) {
	db := s.getDb(ctx)

	return db.SiteSuggestions.Query().Where(sitesuggestions.ID(id)).First(ctx)
}

func (s *sitesRepositoryImpl) getDb(ctx context.Context) *ent.Client {
	db := ent.FromContext(ctx)
	if db == nil {
		return s.db
	}
	return db
}
