// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/ent/isps"
)

// IspsCreate is the builder for creating a Isps entity.
type IspsCreate struct {
	config
	mutation *IspsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *IspsCreate) SetCreatedAt(t time.Time) *IspsCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IspsCreate) SetNillableCreatedAt(t *time.Time) *IspsCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *IspsCreate) SetUpdatedAt(t time.Time) *IspsCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *IspsCreate) SetNillableUpdatedAt(t *time.Time) *IspsCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetLatitude sets the "latitude" field.
func (ic *IspsCreate) SetLatitude(f float64) *IspsCreate {
	ic.mutation.SetLatitude(f)
	return ic
}

// SetLongitude sets the "longitude" field.
func (ic *IspsCreate) SetLongitude(f float64) *IspsCreate {
	ic.mutation.SetLongitude(f)
	return ic
}

// SetName sets the "name" field.
func (ic *IspsCreate) SetName(s string) *IspsCreate {
	ic.mutation.SetName(s)
	return ic
}

// AddIspBlockIDs adds the "isp_blocks" edge to the Blocks entity by IDs.
func (ic *IspsCreate) AddIspBlockIDs(ids ...int) *IspsCreate {
	ic.mutation.AddIspBlockIDs(ids...)
	return ic
}

// AddIspBlocks adds the "isp_blocks" edges to the Blocks entity.
func (ic *IspsCreate) AddIspBlocks(b ...*Blocks) *IspsCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ic.AddIspBlockIDs(ids...)
}

// Mutation returns the IspsMutation object of the builder.
func (ic *IspsCreate) Mutation() *IspsMutation {
	return ic.mutation
}

// Save creates the Isps in the database.
func (ic *IspsCreate) Save(ctx context.Context) (*Isps, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IspsCreate) SaveX(ctx context.Context) *Isps {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IspsCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IspsCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IspsCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := isps.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := isps.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IspsCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Isps.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Isps.updated_at"`)}
	}
	if _, ok := ic.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Isps.latitude"`)}
	}
	if _, ok := ic.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Isps.longitude"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Isps.name"`)}
	}
	return nil
}

func (ic *IspsCreate) sqlSave(ctx context.Context) (*Isps, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IspsCreate) createSpec() (*Isps, *sqlgraph.CreateSpec) {
	var (
		_node = &Isps{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(isps.Table, sqlgraph.NewFieldSpec(isps.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(isps.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(isps.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.Latitude(); ok {
		_spec.SetField(isps.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := ic.mutation.Longitude(); ok {
		_spec.SetField(isps.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(isps.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := ic.mutation.IspBlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   isps.IspBlocksTable,
			Columns: []string{isps.IspBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blocks.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IspsCreateBulk is the builder for creating many Isps entities in bulk.
type IspsCreateBulk struct {
	config
	err      error
	builders []*IspsCreate
}

// Save creates the Isps entities in the database.
func (icb *IspsCreateBulk) Save(ctx context.Context) ([]*Isps, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Isps, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IspsMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IspsCreateBulk) SaveX(ctx context.Context) []*Isps {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IspsCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IspsCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
