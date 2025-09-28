package repository

import (
	"context"
	"log/slog"
	"time"

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
	GetBlocksForISP(ctx context.Context, id int) (map[int][]*ent.Blocks, error)
}

type ispRepositoryImpl struct {
	db *ent.Client
}

func (i *ispRepositoryImpl) CreateISP(ctx context.Context, isp *genapi.ISPInput) (*ent.Isps, error) {
	tx, err := i.db.Tx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	defer tx.Rollback()

	existing, err := tx.Isps.Query().Where(
		isps.Name(isp.Name),
		isps.LatitudeEQ(float64(isp.Latitude)),
		isps.LongitudeEQ(float64(isp.Longitude)),
	).First(ctx)
	var entErr *ent.NotFoundError
	if errors.As(err, &entErr) {
		newIsp, err := i.db.Isps.Create().
			SetName(isp.Name).
			SetLatitude(float64(isp.Latitude)).
			SetLongitude(float64(isp.Longitude)).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}
		return newIsp, tx.Commit()
	} else if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return existing, tx.Commit()
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
			ID:             genapi.NewOptInt(isp.ID),
			Name:           genapi.NewOptString(isp.Name),
			Latitude:       genapi.NewOptFloat32(float32(isp.Latitude)),
			Longitude:      genapi.NewOptFloat32(float32(isp.Longitude)),
			BlockReports:   genapi.NewOptInt(0),
			UnblockReports: genapi.NewOptInt(0),
		}

		for _, block := range isp.Edges.IspBlocks {
			if block.Blocked {
				apiIsp.BlockReports = genapi.NewOptInt(apiIsp.GetBlockReports().Or(0) + 1)
			} else {
				apiIsp.UnblockReports = genapi.NewOptInt(apiIsp.GetUnblockReports().Or(0) + 1)
			}
			if apiIsp.LastReportedAt.Or(time.Time{}).Before(block.LastReportedAt) {
				apiIsp.LastReportedAt = genapi.NewOptDateTime(block.LastReportedAt)
			}
		}

		res = append(res, apiIsp)
	}
	return res, nil
}

func (i *ispRepositoryImpl) GetISPByID(ctx context.Context, id int) (*ent.Isps, error) {
	db := i.getDb(ctx)
	return db.Isps.Query().Where(isps.ID(id)).Only(ctx)
}

func (i *ispRepositoryImpl) GetBlocksForISP(ctx context.Context, id int) (map[int][]*ent.Blocks, error) {
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
	if len(blocks) == 0 {
		return nil, nil
	}

	sites := make(map[int][]*ent.Blocks)
	for _, block := range blocks {
		// Ignore nilderef, Site ID is a mandatory field so it should always exist
		sites[block.SiteID] = append(sites[block.SiteID], block)
	}

	return sites, nil
}

func (i *ispRepositoryImpl) getDb(ctx context.Context) *ent.Client {
	db := ent.FromContext(ctx)
	if db == nil {
		db = i.db
	}
	return db
}
