package service

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
)

type CategoriesService interface {
	GetAllCategories(ctx context.Context) ([]genapi.Category, error)
	CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*genapi.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type categoriesServiceImpl struct {
	categoriesRepository repository.CategoriesRepository
}

func (c *categoriesServiceImpl) GetAllCategories(ctx context.Context) ([]genapi.Category, error) {
	categories, err := c.categoriesRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	var res []genapi.Category
	for _, category := range categories {
		res = append(res, genapi.Category{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	return res, nil
}

func (c *categoriesServiceImpl) CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*genapi.Category, error) {
	// Check if category already exists
	existing, err := c.categoriesRepository.GetCategoryByName(ctx, req.Name)
	if err == nil && existing != nil {
		// Return existing category instead of creating a duplicate
		return &genapi.Category{
			ID:        existing.ID,
			Name:      existing.Name,
			CreatedAt: existing.CreatedAt,
			UpdatedAt: existing.UpdatedAt,
		}, nil
	}

	category, err := c.categoriesRepository.CreateCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &genapi.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (c *categoriesServiceImpl) DeleteCategory(ctx context.Context, id int) error {
	return c.categoriesRepository.DeleteCategory(ctx, id)
}
