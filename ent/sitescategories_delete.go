// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sitescategories"
)

// SitesCategoriesDelete is the builder for deleting a SitesCategories entity.
type SitesCategoriesDelete struct {
	config
	hooks    []Hook
	mutation *SitesCategoriesMutation
}

// Where appends a list predicates to the SitesCategoriesDelete builder.
func (scd *SitesCategoriesDelete) Where(ps ...predicate.SitesCategories) *SitesCategoriesDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SitesCategoriesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SitesCategoriesDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SitesCategoriesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sitescategories.Table, sqlgraph.NewFieldSpec(sitescategories.FieldID, field.TypeInt))
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	scd.mutation.done = true
	return affected, err
}

// SitesCategoriesDeleteOne is the builder for deleting a single SitesCategories entity.
type SitesCategoriesDeleteOne struct {
	scd *SitesCategoriesDelete
}

// Where appends a list predicates to the SitesCategoriesDelete builder.
func (scdo *SitesCategoriesDeleteOne) Where(ps ...predicate.SitesCategories) *SitesCategoriesDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *SitesCategoriesDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sitescategories.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SitesCategoriesDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}
