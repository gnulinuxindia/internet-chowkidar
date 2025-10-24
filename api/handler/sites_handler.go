package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
)

type SitesHandler interface {
	ListSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
	ListSiteSuggestions(ctx context.Context, params genapi.ListSiteSuggestionsParams) ([]genapi.SiteSuggestion, error)
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error)
	CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error)
	ResolveSiteSuggestion(ctx context.Context, req *genapi.ResolveSiteSuggestionInput, params genapi.ResolveSiteSuggestionParams) (*genapi.SiteSuggestion, error)
	GetSite(ctx context.Context, params genapi.GetSiteParams) (*genapi.SiteDetails, error)
	GetSiteSuggestion(ctx context.Context, params genapi.GetSiteSuggestionParams) (*genapi.SiteSuggestion, error)
}

type sitesHandlerImpl struct {
	sitesService service.SitesService
}

func (s *sitesHandlerImpl) ListSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	return s.sitesService.GetAllSites(ctx, params)
}

func (s *sitesHandlerImpl) ListSiteSuggestions(ctx context.Context, params genapi.ListSiteSuggestionsParams) ([]genapi.SiteSuggestion, error) {
	return s.sitesService.GetAllSiteSuggestions(ctx, params)
}

func (s *sitesHandlerImpl) CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error) {
	return s.sitesService.CreateSite(ctx, req)
}

func (s *sitesHandlerImpl) GetSite(ctx context.Context, params genapi.GetSiteParams) (*genapi.SiteDetails, error) {
	siteDetails, err := s.sitesService.GetSite(ctx, params)
	if err != nil {
		return nil, err
	}

	return siteDetails, nil
}

func (s *sitesHandlerImpl) GetSiteSuggestion(ctx context.Context, params genapi.GetSiteSuggestionParams) (*genapi.SiteSuggestion, error) {
	suggestionDetails, err := s.sitesService.GetSiteSuggestion(ctx, params)
	if err != nil {
		return nil, err
	}

	return suggestionDetails, nil
}

func (s *sitesHandlerImpl) ResolveSiteSuggestion(ctx context.Context, req *genapi.ResolveSiteSuggestionInput, params genapi.ResolveSiteSuggestionParams) (*genapi.SiteSuggestion, error) {
	suggestionDetails, err := s.sitesService.ResolveSiteSuggestion(ctx, req, params)
	if err != nil {
		return nil, err
	}

	return suggestionDetails, nil
}

func (s *sitesHandlerImpl) CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error) {
	return s.sitesService.CreateSiteSuggestion(ctx, req)
}
