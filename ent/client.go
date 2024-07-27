// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/gnulinuxindia/internet-chowkidar/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/gnulinuxindia/internet-chowkidar/ent/counter"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Counter is the client for interacting with the Counter builders.
	Counter *CounterClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Counter = NewCounterClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Counter: NewCounterClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Counter: NewCounterClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Counter.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Counter.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Counter.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CounterMutation:
		return c.Counter.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CounterClient is a client for the Counter schema.
type CounterClient struct {
	config
}

// NewCounterClient returns a client for the Counter from the given config.
func NewCounterClient(c config) *CounterClient {
	return &CounterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `counter.Hooks(f(g(h())))`.
func (c *CounterClient) Use(hooks ...Hook) {
	c.hooks.Counter = append(c.hooks.Counter, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `counter.Intercept(f(g(h())))`.
func (c *CounterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Counter = append(c.inters.Counter, interceptors...)
}

// Create returns a builder for creating a Counter entity.
func (c *CounterClient) Create() *CounterCreate {
	mutation := newCounterMutation(c.config, OpCreate)
	return &CounterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Counter entities.
func (c *CounterClient) CreateBulk(builders ...*CounterCreate) *CounterCreateBulk {
	return &CounterCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CounterClient) MapCreateBulk(slice any, setFunc func(*CounterCreate, int)) *CounterCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CounterCreateBulk{err: fmt.Errorf("calling to CounterClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CounterCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CounterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Counter.
func (c *CounterClient) Update() *CounterUpdate {
	mutation := newCounterMutation(c.config, OpUpdate)
	return &CounterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CounterClient) UpdateOne(co *Counter) *CounterUpdateOne {
	mutation := newCounterMutation(c.config, OpUpdateOne, withCounter(co))
	return &CounterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CounterClient) UpdateOneID(id int) *CounterUpdateOne {
	mutation := newCounterMutation(c.config, OpUpdateOne, withCounterID(id))
	return &CounterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Counter.
func (c *CounterClient) Delete() *CounterDelete {
	mutation := newCounterMutation(c.config, OpDelete)
	return &CounterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CounterClient) DeleteOne(co *Counter) *CounterDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CounterClient) DeleteOneID(id int) *CounterDeleteOne {
	builder := c.Delete().Where(counter.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CounterDeleteOne{builder}
}

// Query returns a query builder for Counter.
func (c *CounterClient) Query() *CounterQuery {
	return &CounterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCounter},
		inters: c.Interceptors(),
	}
}

// Get returns a Counter entity by its id.
func (c *CounterClient) Get(ctx context.Context, id int) (*Counter, error) {
	return c.Query().Where(counter.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CounterClient) GetX(ctx context.Context, id int) *Counter {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CounterClient) Hooks() []Hook {
	return c.hooks.Counter
}

// Interceptors returns the client interceptors.
func (c *CounterClient) Interceptors() []Interceptor {
	return c.inters.Counter
}

func (c *CounterClient) mutate(ctx context.Context, m *CounterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CounterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CounterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CounterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CounterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Counter mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Counter []ent.Hook
	}
	inters struct {
		Counter []ent.Interceptor
	}
)
