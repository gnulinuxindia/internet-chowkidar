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
	repo repository.IspRepository
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
		ID: isp.ID,
		Name: isp.Name,
		Latitude: float32(isp.Latitude),
		Longitude: float32(isp.Longitude),
		Blocks: make([]genapi.ISPBlock, len(blocks)),
		CreatedAt: isp.CreatedAt,
		UpdatedAt: isp.UpdatedAt,
	}

	for i, block := range blocks {
		details.Blocks[i] = genapi.ISPBlock{
			ID: block.ID,
			BlockReports: block.BlockReports,
			UnblockReports: block.UnblockReports,
			LastReportedAt: block.LastReportedAt,
			SiteID: block.Edges.Site.ID,
			Domain: block.Edges.Site.Domain,
			CreatedAt: block.CreatedAt,
			UpdatedAt: block.UpdatedAt,
		}
	}


	return details, nil
}