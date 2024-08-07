// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gnulinuxindia/internet-chowkidar/ent/counter"
)

// Counter is the model entity for the Counter schema.
type Counter struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Count holds the value of the "count" field.
	Count        int `json:"count,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Counter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case counter.FieldID, counter.FieldCount:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Counter fields.
func (c *Counter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case counter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case counter.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				c.Count = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Counter.
// This includes values selected through modifiers, order, etc.
func (c *Counter) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Counter.
// Note that you need to call Counter.Unwrap() before calling this method if this Counter
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Counter) Update() *CounterUpdateOne {
	return NewCounterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Counter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Counter) Unwrap() *Counter {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Counter is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Counter) String() string {
	var builder strings.Builder
	builder.WriteString("Counter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", c.Count))
	builder.WriteByte(')')
	return builder.String()
}

// Counters is a parsable slice of Counter.
type Counters []*Counter
