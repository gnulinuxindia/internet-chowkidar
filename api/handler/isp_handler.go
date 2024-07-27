package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
)

type IspHandler interface {
	ListISPs(ctx context.Context) ([]genapi.ISP, error)
	CreateISP(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error)
}

type ispHandlerImpl struct {
}


func (i *ispHandlerImpl) ListISPs(ctx context.Context) ([]genapi.ISP, error) {
	panic("not implemented")
}

func (i *ispHandlerImpl) CreateISP(ctx context.Context, req *genapi.ISPInput) (*genapi.ISP, error) {
	panic("not implemented")
}
