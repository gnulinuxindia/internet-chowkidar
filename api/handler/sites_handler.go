package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
)

type SitesHandler interface {
	ListSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
	ListSiteSuggestions(ctx context.Context) ([]genapi.SiteSuggestion, error)
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error)
	CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error)
}

type sitesHandlerImpl struct {
	sitesService service.SitesService
}

func (s *sitesHandlerImpl) ListSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	return s.sitesService.GetAllSites(ctx, params)
}

func (s *sitesHandlerImpl) ListSiteSuggestions(ctx context.Context) ([]genapi.SiteSuggestion, error) {
	panic("not implemented")
}

func (s *sitesHandlerImpl) CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error) {
	return s.sitesService.CreateSite(ctx, req)
}

func (s *sitesHandlerImpl) CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error) {
	panic("not implemented")
}
