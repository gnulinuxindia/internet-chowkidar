package repository

import (
	"context"
	"log/slog"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
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
	return s.db.Sites.Create().
		SetDomain(req.Domain).
		Save(ctx)
}

func (s *sitesRepositoryImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	query := s.db.Blocks.Query().
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

	sites := map[string]*genapi.Site{}

	for _, block := range blocks {
		if _, ok := sites[block.Edges.Site.Domain]; !ok {
			// Add the site to the map
			sites[block.Edges.Site.Domain] = &genapi.Site{
				ID:             block.Edges.Site.ID,
				Domain:         block.Edges.Site.Domain,
				BlockReports:   block.BlockReports,
				UnblockReports: block.UnblockReports,
				LastReportedAt: block.LastReportedAt,
			}
		} else {
			// Update the existing site
			// Add the block and unblock reports
			sites[block.Edges.Site.Domain].BlockReports += block.BlockReports
			sites[block.Edges.Site.Domain].UnblockReports += block.UnblockReports

			// Update the last reported at
			if sites[block.Edges.Site.Domain].LastReportedAt.Before(block.LastReportedAt) {
				sites[block.Edges.Site.Domain].LastReportedAt = block.LastReportedAt
			}
		}
	}

	// convert the map to a slice
	result := make([]genapi.Site, 0, len(sites))
	for _, site := range sites {
		result = append(result, *site)
	}

	return nil, nil
}

func (s *sitesRepositoryImpl) GetSiteByDomain(ctx context.Context, domain string) (*ent.Sites, error) {
	return s.db.Sites.Query().
		Where(sites.DomainEQ(domain)).
		First(ctx)
}

func (s *sitesRepositoryImpl) GetSiteByID(ctx context.Context, id int) (*ent.Sites, error) {
	return s.db.Sites.Get(ctx, id)
}
