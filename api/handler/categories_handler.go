package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
)

type CategoryHandler interface {
	ListCategories(ctx context.Context) ([]genapi.Category, error)
	CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*genapi.Category, error)
}

type categoryHandlerImpl struct {
}

func (c *categoryHandlerImpl) ListCategories(ctx context.Context) ([]genapi.Category, error) {
	panic("not implemented")
}

func (c *categoryHandlerImpl) CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*genapi.Category, error) {
	panic("not implemented")
}
