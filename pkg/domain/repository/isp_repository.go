package repository

import (
	"context"
	"log/slog"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/ent/isps"
	"github.com/go-errors/errors"
)

type IspRepository interface {
	CreateISP(ctx context.Context, isp *genapi.ISPInput) (*ent.Isps, error)
	GetAllISPs(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error)
	GetISPByID(ctx context.Context, id int) (*ent.Isps, error)
	GetBlocksForISP(ctx context.Context, id int) ([]*ent.Blocks, error)
}

type ispRepositoryImpl struct {
	db *ent.Client
}

func (i *ispRepositoryImpl) CreateISP(ctx context.Context, isp *genapi.ISPInput) (*ent.Isps, error) {
	return i.db.Isps.Create().
		SetName(isp.Name).
		SetLatitude(float64(isp.Latitude)).
		SetLongitude(float64(isp.Longitude)).
		Save(ctx)
}

func (i *ispRepositoryImpl) GetAllISPs(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error) {
	query := i.db.Isps.Query().
		Limit(params.Limit.Or(50)).
		Offset(params.Offset.Or(0)).
		WithIspBlocks()

	if params.Order.Set {
		if params.Order.Value == genapi.ListISPsOrderAsc {
			query = query.Order(ent.Asc(params.Sort.Or("id")))
		} else {
			query = query.Order(ent.Desc(params.Sort.Or("id")))
		}
	}

	isps, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	var res []genapi.ISP
	for _, isp := range isps {
		apiIsp := genapi.ISP{
			ID:        genapi.NewOptInt(isp.ID),
			Name:      genapi.NewOptString(isp.Name),
			Latitude:  genapi.NewOptFloat32(float32(isp.Latitude)),
			Longitude: genapi.NewOptFloat32(float32(isp.Longitude)),
			BlockReports: genapi.NewOptInt(0),
			UnblockReports: genapi.NewOptInt(0),
		}

		for _, block := range isp.Edges.IspBlocks {
			apiIsp.BlockReports = genapi.NewOptInt(apiIsp.GetBlockReports().Or(0) + block.BlockReports)
			apiIsp.UnblockReports = genapi.NewOptInt(apiIsp.GetUnblockReports().Or(0) + block.UnblockReports)
		}

		res = append(res, apiIsp)
	}
	return res, nil
}

func (i *ispRepositoryImpl) GetISPByID(ctx context.Context, id int) (*ent.Isps, error) {
	db := i.getDb(ctx)
	return db.Isps.Query().Where(isps.ID(id)).Only(ctx)
}

func (i *ispRepositoryImpl) GetBlocksForISP(ctx context.Context, id int) ([]*ent.Blocks, error) {
	db := i.getDb(ctx)
	blocks, err := db.Blocks.Query().
		Where(blocks.HasIspWith(isps.ID(id))).
		WithSite().
		WithIsp().
		All(ctx)
	if err != nil {
		slog.Error("failed to get blocks for isp", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	return blocks, nil
}

func (i *ispRepositoryImpl) getDb(ctx context.Context) *ent.Client {
	db := ent.FromContext(ctx)
	if db == nil {
		db = i.db
	}
	return db
}
