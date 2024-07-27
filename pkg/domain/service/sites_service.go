package service

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
)

type SitesService interface {
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
}

type sitesServiceImpl struct {
	repo repository.SitesRepository
}

func (s *sitesServiceImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	// TODO: add validation and business logic etc
	return s.repo.GetAllSites(ctx, params)
}