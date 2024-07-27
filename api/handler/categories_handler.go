package handler

import (
	"context"
	"errors"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
)

type CategoryHandler interface {
	ListCategories(ctx context.Context) ([]genapi.Category, error)
	CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*genapi.Category, error)
}

type categoryHandlerImpl struct {
	categoriesService service.CategoriesService
}

func (c *categoryHandlerImpl) ListCategories(ctx context.Context) ([]genapi.Category, error) {
	categories, err := c.categoriesService.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryHandlerImpl) CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*genapi.Category, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	category, err := c.categoriesService.CreateCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return category, nil
}
