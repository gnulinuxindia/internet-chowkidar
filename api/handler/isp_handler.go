package handler

import (
	"context"
	"errors"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
)

type IspHandler interface {
	ListISPs(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error)
	CreateISP(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error)
}

type ispHandlerImpl struct {
	ispService service.IspService
}

func (i *ispHandlerImpl) ListISPs(ctx context.Context, params genapi.ListISPsParams) ([]genapi.ISP, error) {
	return i.ispService.GetAllIsps(ctx, params)
}

func (i *ispHandlerImpl) CreateISP(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	isp, err := i.ispService.CreateIsp(ctx, req)
	if err != nil {
		return nil, err
	}

	return isp, nil
}
