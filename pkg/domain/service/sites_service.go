package service

import (
	"context"
	"strings"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
	"github.com/go-errors/errors"
)

type SitesService interface {
	CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error)
	CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error)
	ResolveSiteSuggestion(ctx context.Context, req *genapi.ResolveSiteSuggestionInput, params genapi.ResolveSiteSuggestionParams) (*genapi.SiteSuggestion, error)
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
	GetAllSiteSuggestions(ctx context.Context, params genapi.ListSiteSuggestionsParams) ([]genapi.SiteSuggestion, error)
	GetSite(ctx context.Context, params genapi.GetSiteParams) (*genapi.SiteDetails, error)
	GetSiteSuggestion(ctx context.Context, params genapi.GetSiteSuggestionParams) (*genapi.SiteSuggestion, error)
}

type sitesServiceImpl struct {
	sitesRepository  repository.SitesRepository
	blocksRepository repository.BlocksRepository
	txHandler        repository.TxHandler
}

func (s *sitesServiceImpl) CreateSite(ctx context.Context, req *genapi.SiteInput) (*genapi.SiteCreate, error) {
	req.Categories = append(req.Categories, "all")

	site, err := s.sitesRepository.CreateSite(ctx, req)
	if err != nil {
		return nil, err
	}

	return &genapi.SiteCreate{
		ID:        site.ID,
		Domain:    site.Domain,
		PingURL:   site.PingURL,
		CreatedAt: site.CreatedAt,
		UpdatedAt: site.UpdatedAt,
	}, nil
}

func (s *sitesServiceImpl) ResolveSiteSuggestion(ctx context.Context, req *genapi.ResolveSiteSuggestionInput, params genapi.ResolveSiteSuggestionParams) (*genapi.SiteSuggestion, error) {
	site, err := s.sitesRepository.GetSiteSuggestionByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if req.Status == "accepted" {
		data := &genapi.SiteInput{
			Domain:     site.Domain,
			PingURL:    genapi.NewOptString(site.PingURL),
			Categories: strings.Split(site.Categories, ","),
		}
		if req.Domain.Set {
			data.Domain = req.Domain.Value
		}
		if req.PingURL.Set {
			data.PingURL = req.PingURL
		}
		if len(req.Categories) > 0 {
			data.Categories = req.Categories
		}
		_, err := s.CreateSite(ctx, data)
		if err != nil {
			return nil, err
		}
	}
	site2, err := s.sitesRepository.ResolveSiteSuggestion(ctx, req, params)
	if err != nil {
		return nil, err
	}
	return &genapi.SiteSuggestion{
		ID:            site2.ID,
		Domain:        site2.Domain,
		PingURL:       site2.PingURL,
		Categories:    strings.Split(site2.Categories, ","),
		Reason:        site2.Reason,
		Status:        genapi.SiteSuggestionStatus(site2.Status),
		ResolveReason: site2.ResolveReason,
		ResolvedAt:    site2.ResolvedAt,
		CreatedAt:     site2.CreatedAt,
		UpdatedAt:     site2.UpdatedAt,
	}, nil
}

func (s *sitesServiceImpl) CreateSiteSuggestion(ctx context.Context, req *genapi.SiteSuggestionInput) (*genapi.SiteSuggestion, error) {
	suggestion, err := s.sitesRepository.CreateSiteSuggestion(ctx, req)
	if err != nil {
		return nil, err
	}

	return &genapi.SiteSuggestion{
		ID:            suggestion.ID,
		Domain:        suggestion.Domain,
		Categories:    strings.Split(suggestion.Categories, ","),
		PingURL:       suggestion.PingURL,
		Reason:        suggestion.Reason,
		Status:        genapi.SiteSuggestionStatus(suggestion.Status),
		ResolveReason: suggestion.ResolveReason,
		CreatedAt:     suggestion.CreatedAt,
		UpdatedAt:     suggestion.UpdatedAt,
	}, nil
}

func (s *sitesServiceImpl) GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error) {
	// TODO: add validation and business logic etc
	return s.sitesRepository.GetAllSites(ctx, params)
}

func (s *sitesServiceImpl) GetAllSiteSuggestions(ctx context.Context, params genapi.ListSiteSuggestionsParams) ([]genapi.SiteSuggestion, error) {
	// TODO: add validation and business logic etc
	return s.sitesRepository.GetAllSiteSuggestions(ctx, params)
}

func (s *sitesServiceImpl) GetSite(ctx context.Context, params genapi.GetSiteParams) (*genapi.SiteDetails, error) {
	tx, err := s.txHandler.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	defer tx.Rollback()

	txCtx := ent.NewContext(ctx, tx.Client())

	blocks, err := s.sitesRepository.GetSiteBlocksByID(txCtx, params.ID)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	ds, err := s.sitesRepository.GetSiteByID(txCtx, params.ID)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	site := &genapi.SiteDetails{
		ID:         ds.ID,
		Domain:     ds.Domain,
		PingURL:    ds.PingURL,
		CreatedAt:  ds.CreatedAt,
		UpdatedAt:  ds.UpdatedAt,
		Categories: make([]string, len(ds.Edges.Categories)),
	}

	for i, category := range ds.Edges.Categories {
		site.Categories[i] = category.Name
	}

	for _, block := range blocks {
		blockReports, unblockReports := 0, 0
		blockIsp := block[0].Edges.Isp
		for _, block2 := range block {
			if block2.Blocked {
				site.BlockReports += 1
				blockReports += 1
			} else {
				site.UnblockReports += 1
				unblockReports += 1
			}
			if site.LastReportedAt.Before(block2.LastReportedAt) {
				site.LastReportedAt = block2.LastReportedAt
			}
		}
		isp := genapi.ISP{
			ID:             genapi.NewOptInt(blockIsp.ID),
			Name:           genapi.NewOptString(blockIsp.Name),
			Latitude:       genapi.NewOptFloat32(float32(blockIsp.Latitude)),
			Longitude:      genapi.NewOptFloat32(float32(blockIsp.Longitude)),
			BlockReports:   genapi.NewOptInt(blockReports),
			UnblockReports: genapi.NewOptInt(unblockReports),
			CreatedAt:      genapi.NewOptDateTime(blockIsp.CreatedAt),
			UpdatedAt:      genapi.NewOptDateTime(blockIsp.UpdatedAt),
		}

		site.BlockedByIsps = append(site.BlockedByIsps, isp)
	}

	return site, tx.Commit()
}

func (s *sitesServiceImpl) GetSiteSuggestion(ctx context.Context, params genapi.GetSiteSuggestionParams) (*genapi.SiteSuggestion, error) {
	tx, err := s.txHandler.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	defer tx.Rollback()

	txCtx := ent.NewContext(ctx, tx.Client())

	ds, err := s.sitesRepository.GetSiteSuggestionByID(txCtx, params.ID)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	site := &genapi.SiteSuggestion{
		ID:            ds.ID,
		Domain:        ds.Domain,
		PingURL:       ds.PingURL,
		Categories:    strings.Split(ds.Categories, ","),
		Reason:        ds.Reason,
		Status:        genapi.SiteSuggestionStatus(ds.Status),
		ResolveReason: ds.ResolveReason,
		ResolvedAt:    ds.ResolvedAt,
		CreatedAt:     ds.CreatedAt,
		UpdatedAt:     ds.UpdatedAt,
	}

	return site, tx.Commit()
}
