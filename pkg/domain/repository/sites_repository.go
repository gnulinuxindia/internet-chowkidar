package repository

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/go-errors/errors"
)

type SitesRepository interface {
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
}

type sitesRepositoryImpl struct {
	db *ent.Client
}

func (s *sitesRepositoryImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	blocks, err := s.db.Blocks.Query().
		WithIsp().
		WithSite().
		Limit(params.Limit.Or(50)).
		Offset(params.Offset.Or(0)).
		All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

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
