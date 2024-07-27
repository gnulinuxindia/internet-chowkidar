package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
)

type SitesHandler interface {
	ListSites(ctx context.Context) ([]genapi.Site, error)
	ListSiteSuggestions(ctx context.Context) ([]genapi.SiteSuggestion, error)
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.Site, error)
	CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error)
}

type sitesHandlerImpl struct {

}

func (s *sitesHandlerImpl) ListSites(ctx context.Context) ([]genapi.Site, error) {
	panic("not implemented")
}

func (s *sitesHandlerImpl) ListSiteSuggestions(ctx context.Context) ([]genapi.SiteSuggestion, error) {
	panic("not implemented")
}

func (s *sitesHandlerImpl) CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.Site, error) {
	panic("not implemented")
}

func (s *sitesHandlerImpl) CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error) {
	panic("not implemented")
}
