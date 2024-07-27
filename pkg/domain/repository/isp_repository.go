package repository

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
)

type IspRepository interface {
	CreateISP(ctx context.Context, isp *genapi.ISPInput) (*ent.Isps, error)
	GetAllISPs(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error)
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
	isps, err := i.db.Isps.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var res []genapi.ISP
	for _, isp := range isps {
		res = append(res, genapi.ISP{
			ID:        genapi.NewOptInt(isp.ID),
			Name:      genapi.NewOptString(isp.Name),
			Latitude:  genapi.NewOptFloat32(float32(isp.Latitude)),
			Longitude: genapi.NewOptFloat32(float32(isp.Longitude)),
		})
	}
	return res, nil
}
