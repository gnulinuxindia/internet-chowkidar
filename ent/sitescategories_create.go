// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/categories"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sites"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sitescategories"
)

// SitesCategoriesCreate is the builder for creating a SitesCategories entity.
type SitesCategoriesCreate struct {
	config
	mutation *SitesCategoriesMutation
	hooks    []Hook
}

// SetSitesID sets the "sites_id" field.
func (scc *SitesCategoriesCreate) SetSitesID(i int) *SitesCategoriesCreate {
	scc.mutation.SetSitesID(i)
	return scc
}

// SetCategoriesID sets the "categories_id" field.
func (scc *SitesCategoriesCreate) SetCategoriesID(i int) *SitesCategoriesCreate {
	scc.mutation.SetCategoriesID(i)
	return scc
}

// SetSites sets the "sites" edge to the Sites entity.
func (scc *SitesCategoriesCreate) SetSites(s *Sites) *SitesCategoriesCreate {
	return scc.SetSitesID(s.ID)
}

// SetCategories sets the "categories" edge to the Categories entity.
func (scc *SitesCategoriesCreate) SetCategories(c *Categories) *SitesCategoriesCreate {
	return scc.SetCategoriesID(c.ID)
}

// Mutation returns the SitesCategoriesMutation object of the builder.
func (scc *SitesCategoriesCreate) Mutation() *SitesCategoriesMutation {
	return scc.mutation
}

// Save creates the SitesCategories in the database.
func (scc *SitesCategoriesCreate) Save(ctx context.Context) (*SitesCategories, error) {
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SitesCategoriesCreate) SaveX(ctx context.Context) *SitesCategories {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SitesCategoriesCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SitesCategoriesCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SitesCategoriesCreate) check() error {
	if _, ok := scc.mutation.SitesID(); !ok {
		return &ValidationError{Name: "sites_id", err: errors.New(`ent: missing required field "SitesCategories.sites_id"`)}
	}
	if _, ok := scc.mutation.CategoriesID(); !ok {
		return &ValidationError{Name: "categories_id", err: errors.New(`ent: missing required field "SitesCategories.categories_id"`)}
	}
	if _, ok := scc.mutation.SitesID(); !ok {
		return &ValidationError{Name: "sites", err: errors.New(`ent: missing required edge "SitesCategories.sites"`)}
	}
	if _, ok := scc.mutation.CategoriesID(); !ok {
		return &ValidationError{Name: "categories", err: errors.New(`ent: missing required edge "SitesCategories.categories"`)}
	}
	return nil
}

func (scc *SitesCategoriesCreate) sqlSave(ctx context.Context) (*SitesCategories, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *SitesCategoriesCreate) createSpec() (*SitesCategories, *sqlgraph.CreateSpec) {
	var (
		_node = &SitesCategories{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(sitescategories.Table, sqlgraph.NewFieldSpec(sitescategories.FieldID, field.TypeInt))
	)
	if nodes := scc.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sitescategories.SitesTable,
			Columns: []string{sitescategories.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sites.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SitesID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sitescategories.CategoriesTable,
			Columns: []string{sitescategories.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(categories.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoriesID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SitesCategoriesCreateBulk is the builder for creating many SitesCategories entities in bulk.
type SitesCategoriesCreateBulk struct {
	config
	err      error
	builders []*SitesCategoriesCreate
}

// Save creates the SitesCategories entities in the database.
func (sccb *SitesCategoriesCreateBulk) Save(ctx context.Context) ([]*SitesCategories, error) {
	if sccb.err != nil {
		return nil, sccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SitesCategories, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SitesCategoriesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SitesCategoriesCreateBulk) SaveX(ctx context.Context) []*SitesCategories {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SitesCategoriesCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SitesCategoriesCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
