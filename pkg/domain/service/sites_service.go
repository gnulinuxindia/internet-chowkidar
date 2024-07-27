package service

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
)

type SitesService interface {
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error)
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
}

type sitesServiceImpl struct {
	sitesRepository  repository.SitesRepository
	blocksRepository repository.BlocksRepository
}

func (s *sitesServiceImpl) CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error) {
	site, err := s.sitesRepository.CreateSite(ctx, req)
	if err != nil {
		return nil, err
	}

	return &genapi.SiteCreate{
		ID:        site.ID,
		Domain:    site.Domain,
		CreatedAt: site.CreatedAt,
		UpdatedAt: site.UpdatedAt,
	}, nil
}

func (s *sitesServiceImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	// TODO: add validation and business logic etc
	return s.sitesRepository.GetAllSites(ctx, params)
}
