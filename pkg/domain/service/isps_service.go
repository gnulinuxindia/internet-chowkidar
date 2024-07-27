package service

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
)

type IspService interface {
	CreateIsp(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error)
	GetAllIsps(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error)
}

type ispServiceImpl struct {
	repo repository.IspRepository
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
