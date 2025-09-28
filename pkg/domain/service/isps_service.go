package service

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
	"github.com/go-errors/errors"
)

type IspService interface {
	CreateIsp(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error)
	GetAllIsps(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error)
	GetISP(ctx context.Context, params genapi.GetISPParams) (*genapi.ISPDetails, error)
}

type ispServiceImpl struct {
	repo      repository.IspRepository
	txHandler repository.TxHandler
}

func (i *ispServiceImpl) CreateIsp(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error) {
	isp, err := i.repo.CreateISP(ctx, req)
	if err != nil {
		return nil, err
	}

	return &genapi.ISP{
		ID:        genapi.NewOptInt(isp.ID),
		Name:      genapi.NewOptString(isp.Name),
		Latitude:  genapi.NewOptFloat32(float32(isp.Latitude)),
		Longitude: genapi.NewOptFloat32(float32(isp.Longitude)),
	}, nil
}

func (i *ispServiceImpl) GetAllIsps(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error) {
	isps, err := i.repo.GetAllISPs(ctx, params)
	if err != nil {
		return nil, err
	}

	return isps, nil
}

func (i *ispServiceImpl) GetISP(ctx context.Context, params genapi.GetISPParams) (*genapi.ISPDetails, error) {
	tx, err := i.txHandler.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	defer tx.Rollback()

	tCtx := ent.NewContext(ctx, tx.Client())

	isp, err := i.repo.GetISPByID(tCtx, params.ID)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	blocks, err := i.repo.GetBlocksForISP(tCtx, params.ID)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	details := &genapi.ISPDetails{
		ID:        isp.ID,
		Name:      isp.Name,
		Latitude:  float32(isp.Latitude),
		Longitude: float32(isp.Longitude),
		Blocks:    make([]genapi.ISPBlock, len(blocks)),
		CreatedAt: isp.CreatedAt,
		UpdatedAt: isp.UpdatedAt,
	}

	count := 0
	for i, block := range blocks {
		blockReports, unblockReports := 0, 0
		blockSites := block[0].Edges.Site
		lastReportedAt := block[0].LastReportedAt
		updatedAt := block[0].UpdatedAt
		createdAt := block[0].CreatedAt
		for _, block2 := range block {
			if block2.Blocked {
				blockReports += 1
			} else {
				unblockReports += 1
			}
			if lastReportedAt.Before(block2.LastReportedAt) {
				lastReportedAt = block2.LastReportedAt
			}
			if updatedAt.Before(block2.UpdatedAt) {
				updatedAt = block2.UpdatedAt
			}
			if createdAt.After(block2.CreatedAt) {
				createdAt = block2.CreatedAt
			}
		}
		details.Blocks[count] = genapi.ISPBlock{
			ID:             i,
			BlockReports:   blockReports,
			UnblockReports: unblockReports,
			LastReportedAt: lastReportedAt,
			SiteID:         blockSites.ID,
			Domain:         blockSites.Domain,
			CreatedAt:      block[0].CreatedAt,
			UpdatedAt:      block[0].UpdatedAt,
		}
		count += 1
	}

	return details, nil
}
