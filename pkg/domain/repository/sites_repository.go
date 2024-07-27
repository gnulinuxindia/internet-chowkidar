package repository

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	// "github.com/gnulinuxindia/internet-chowkidar/ent/sites"
)

type SitesRepository interface {
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]*ent.Sites, error)
}

type sitesRepositoryImpl struct {
	db *ent.Client
}

func (s *sitesRepositoryImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]*ent.Sites, error) {
	s.db.Sites.Query().Where(
		// sites
	).
	Limit(params.Limit.Or(50)).
	All(ctx)

	return nil, nil
}
