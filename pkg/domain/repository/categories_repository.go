package repository

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/ent/categories"
)

type CategoriesRepository interface {
	GetAllCategories(ctx context.Context) ([]*ent.Categories, error)
	CreateCategory(ctx context.Context, req *genapi.CreateCategoryReq) (*ent.Categories, error)
	GetCategoryByName(ctx context.Context, name string) (*ent.Categories, error)
	DeleteCategory(ctx context.Context, id int) error
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

func (c *categoriesRepositoryImpl) GetCategoryByName(ctx context.Context, name string) (*ent.Categories, error) {
	return c.db.Categories.Query().Where(categories.Name(name)).Only(ctx)
}

func (c *categoriesRepositoryImpl) DeleteCategory(ctx context.Context, id int) error {
	return c.db.Categories.DeleteOneID(id).Exec(ctx)
}
