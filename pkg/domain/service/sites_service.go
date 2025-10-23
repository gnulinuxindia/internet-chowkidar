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
	GetAllSites(ctx context.Context, params genapi.ListSitesParams) ([]genapi.Site, error)
	GetSite(ctx context.Context, params genapi.GetSiteParams) (*genapi.SiteDetails, error)
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
