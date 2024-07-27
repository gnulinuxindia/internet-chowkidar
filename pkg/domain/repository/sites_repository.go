package repository

import (
	"context"
	"log/slog"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/ent/categories"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sites"
	"github.com/go-errors/errors"
)

type SitesRepository interface {
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*ent.Sites, error)
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
	GetSiteByDomain(ctx context.Context, domain string) (*ent.Sites, error)
	GetSiteByID(ctx context.Context, id int) (*ent.Sites, error)
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
	site, err := tx.Sites.Create().
		SetDomain(req.Domain).
		AddCategories(categories...).
		Save(ctx)
	if err != nil {
		slog.Error("failed to create site", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	return site, tx.Commit()
}

func (s *sitesRepositoryImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		slog.Error("failed to start transaction", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	query := tx.Blocks.Query().
		WithIsp().
		WithSite().
		Limit(params.Limit.Or(50)).
		Offset(params.Offset.Or(0))

	if params.Order.Set {
		if params.Order.Value == genapi.ListSitesOrderAsc {
			query = query.Order(ent.Asc(params.Sort.Or("id")))
		} else {
			query = query.Order(ent.Desc(params.Sort.Or("id")))
		}
	}

	blocks, err := query.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	slog.Debug("blocks", "blocks", blocks)

	// get all sites in the database
	dbSites, err := tx.Sites.Query().WithCategories().All(ctx)
	if err != nil {
		slog.Error("failed to get sites", "error", err)
		return nil, rollback(tx, errors.Wrap(err, 0))
	}

	// create a map of sites
	sites := map[string]*genapi.Site{}
	for _, dbSite := range dbSites {

		// convert the categories to a slice of strings
		c := make([]string, len(dbSite.Edges.Categories))
		for i, category := range dbSite.Edges.Categories {
			c[i] = category.Name
		}

		sites[dbSite.Domain] = &genapi.Site{
			ID:         dbSite.ID,
			Domain:     dbSite.Domain,
			Categories: c,
			CreatedAt:  dbSite.CreatedAt,
			UpdatedAt:  dbSite.UpdatedAt,
		}
	}
			

	for _, block := range blocks {
		site := block.Edges.Site

		if _, ok := sites[site.Domain]; !ok {

			// Get the site categories
			categories, err := site.QueryCategories().All(ctx)
			if err != nil {
				// no need to be a fatal error, just log it
				slog.Error("failed to get site categories", "error", err)
			}

			c := make([]string, len(categories))
			for i, category := range categories {
				c[i] = category.Name
			}

			// Add the site to the map
			sites[site.Domain] = &genapi.Site{
				ID:             site.ID,
				Domain:         site.Domain,
				BlockReports:   block.BlockReports,
				UnblockReports: block.UnblockReports,
				LastReportedAt: block.LastReportedAt,
				Categories:     c,
				CreatedAt:      site.CreatedAt,
				UpdatedAt:      site.UpdatedAt,
			}
		} else {
			// Update the existing site
			// Add the block and unblock reports
			sites[site.Domain].BlockReports += block.BlockReports
			sites[site.Domain].UnblockReports += block.UnblockReports

			// Update the last reported at
			if sites[site.Domain].LastReportedAt.Before(block.LastReportedAt) {
				sites[site.Domain].LastReportedAt = block.LastReportedAt
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

func (s *sitesRepositoryImpl) GetSiteByDomain(ctx context.Context, domain string) (*ent.Sites, error) {
	return s.db.Sites.Query().
		Where(sites.DomainEQ(domain)).
		First(ctx)
}

func (s *sitesRepositoryImpl) GetSiteByID(ctx context.Context, id int) (*ent.Sites, error) {
	return s.db.Sites.Get(ctx, id)
}
