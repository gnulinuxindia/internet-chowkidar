// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/ent/isps"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
)

// IspsQuery is the builder for querying Isps entities.
type IspsQuery struct {
	config
	ctx           *QueryContext
	order         []isps.OrderOption
	inters        []Interceptor
	predicates    []predicate.Isps
	withIspBlocks *BlocksQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IspsQuery builder.
func (iq *IspsQuery) Where(ps ...predicate.Isps) *IspsQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IspsQuery) Limit(limit int) *IspsQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IspsQuery) Offset(offset int) *IspsQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IspsQuery) Unique(unique bool) *IspsQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IspsQuery) Order(o ...isps.OrderOption) *IspsQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryIspBlocks chains the current query on the "isp_blocks" edge.
func (iq *IspsQuery) QueryIspBlocks() *BlocksQuery {
	query := (&BlocksClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(isps.Table, isps.FieldID, selector),
			sqlgraph.To(blocks.Table, blocks.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, isps.IspBlocksTable, isps.IspBlocksColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Isps entity from the query.
// Returns a *NotFoundError when no Isps was found.
func (iq *IspsQuery) First(ctx context.Context) (*Isps, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{isps.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IspsQuery) FirstX(ctx context.Context) *Isps {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Isps ID from the query.
// Returns a *NotFoundError when no Isps ID was found.
func (iq *IspsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{isps.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IspsQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Isps entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Isps entity is found.
// Returns a *NotFoundError when no Isps entities are found.
func (iq *IspsQuery) Only(ctx context.Context) (*Isps, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{isps.Label}
	default:
		return nil, &NotSingularError{isps.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IspsQuery) OnlyX(ctx context.Context) *Isps {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Isps ID in the query.
// Returns a *NotSingularError when more than one Isps ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IspsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{isps.Label}
	default:
		err = &NotSingularError{isps.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IspsQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IspsSlice.
func (iq *IspsQuery) All(ctx context.Context) ([]*Isps, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Isps, *IspsQuery]()
	return withInterceptors[[]*Isps](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IspsQuery) AllX(ctx context.Context) []*Isps {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Isps IDs.
func (iq *IspsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(isps.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IspsQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IspsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IspsQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IspsQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IspsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IspsQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IspsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IspsQuery) Clone() *IspsQuery {
	if iq == nil {
		return nil
	}
	return &IspsQuery{
		config:        iq.config,
		ctx:           iq.ctx.Clone(),
		order:         append([]isps.OrderOption{}, iq.order...),
		inters:        append([]Interceptor{}, iq.inters...),
		predicates:    append([]predicate.Isps{}, iq.predicates...),
		withIspBlocks: iq.withIspBlocks.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithIspBlocks tells the query-builder to eager-load the nodes that are connected to
// the "isp_blocks" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IspsQuery) WithIspBlocks(opts ...func(*BlocksQuery)) *IspsQuery {
	query := (&BlocksClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withIspBlocks = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Isps.Query().
//		GroupBy(isps.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *IspsQuery) GroupBy(field string, fields ...string) *IspsGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IspsGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = isps.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Isps.Query().
//		Select(isps.FieldCreatedAt).
//		Scan(ctx, &v)
func (iq *IspsQuery) Select(fields ...string) *IspsSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IspsSelect{IspsQuery: iq}
	sbuild.label = isps.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IspsSelect configured with the given aggregations.
func (iq *IspsQuery) Aggregate(fns ...AggregateFunc) *IspsSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IspsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !isps.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IspsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Isps, error) {
	var (
		nodes       = []*Isps{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withIspBlocks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Isps).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Isps{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withIspBlocks; query != nil {
		if err := iq.loadIspBlocks(ctx, query, nodes,
			func(n *Isps) { n.Edges.IspBlocks = []*Blocks{} },
			func(n *Isps, e *Blocks) { n.Edges.IspBlocks = append(n.Edges.IspBlocks, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *IspsQuery) loadIspBlocks(ctx context.Context, query *BlocksQuery, nodes []*Isps, init func(*Isps), assign func(*Isps, *Blocks)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Isps)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(blocks.FieldIspID)
	}
	query.Where(predicate.Blocks(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(isps.IspBlocksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.IspID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "isp_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *IspsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IspsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(isps.Table, isps.Columns, sqlgraph.NewFieldSpec(isps.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, isps.FieldID)
		for i := range fields {
			if fields[i] != isps.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IspsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(isps.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = isps.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IspsGroupBy is the group-by builder for Isps entities.
type IspsGroupBy struct {
	selector
	build *IspsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IspsGroupBy) Aggregate(fns ...AggregateFunc) *IspsGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IspsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IspsQuery, *IspsGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IspsGroupBy) sqlScan(ctx context.Context, root *IspsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IspsSelect is the builder for selecting fields of Isps entities.
type IspsSelect struct {
	*IspsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IspsSelect) Aggregate(fns ...AggregateFunc) *IspsSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IspsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IspsQuery, *IspsSelect](ctx, is.IspsQuery, is, is.inters, v)
}

func (is *IspsSelect) sqlScan(ctx context.Context, root *IspsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}