package repository

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
)

type CategoriesRepository interface {
	GetAllCategories(ctx context.Context) ([]*ent.Categories, error)
	CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*ent.Categories, error)
}

type categoriesRepositoryImpl struct {
	db *ent.Client
}


func (c *categoriesRepositoryImpl) GetAllCategories(ctx context.Context) ([]*ent.Categories, error) {
	return c.db.Categories.Query().All(ctx)
}

func (c *categoriesRepositoryImpl) CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*ent.Categories, error) {
	return c.db.Categories.Create().SetName(req.Name).Save(ctx)
}

