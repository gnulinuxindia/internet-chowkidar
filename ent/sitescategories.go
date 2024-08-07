// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gnulinuxindia/internet-chowkidar/ent/categories"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sites"
	"github.com/gnulinuxindia/internet-chowkidar/ent/sitescategories"
)

// SitesCategories is the model entity for the SitesCategories schema.
type SitesCategories struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SitesID holds the value of the "sites_id" field.
	SitesID int `json:"sites_id,omitempty"`
	// CategoriesID holds the value of the "categories_id" field.
	CategoriesID int `json:"categories_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SitesCategoriesQuery when eager-loading is set.
	Edges        SitesCategoriesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SitesCategoriesEdges holds the relations/edges for other nodes in the graph.
type SitesCategoriesEdges struct {
	// Sites holds the value of the sites edge.
	Sites *Sites `json:"sites,omitempty"`
	// Categories holds the value of the categories edge.
	Categories *Categories `json:"categories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SitesOrErr returns the Sites value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SitesCategoriesEdges) SitesOrErr() (*Sites, error) {
	if e.Sites != nil {
		return e.Sites, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: sites.Label}
	}
	return nil, &NotLoadedError{edge: "sites"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SitesCategoriesEdges) CategoriesOrErr() (*Categories, error) {
	if e.Categories != nil {
		return e.Categories, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: categories.Label}
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SitesCategories) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sitescategories.FieldID, sitescategories.FieldSitesID, sitescategories.FieldCategoriesID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SitesCategories fields.
func (sc *SitesCategories) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sitescategories.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case sitescategories.FieldSitesID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sites_id", values[i])
			} else if value.Valid {
				sc.SitesID = int(value.Int64)
			}
		case sitescategories.FieldCategoriesID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field categories_id", values[i])
			} else if value.Valid {
				sc.CategoriesID = int(value.Int64)
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SitesCategories.
// This includes values selected through modifiers, order, etc.
func (sc *SitesCategories) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// QuerySites queries the "sites" edge of the SitesCategories entity.
func (sc *SitesCategories) QuerySites() *SitesQuery {
	return NewSitesCategoriesClient(sc.config).QuerySites(sc)
}

// QueryCategories queries the "categories" edge of the SitesCategories entity.
func (sc *SitesCategories) QueryCategories() *CategoriesQuery {
	return NewSitesCategoriesClient(sc.config).QueryCategories(sc)
}

// Update returns a builder for updating this SitesCategories.
// Note that you need to call SitesCategories.Unwrap() before calling this method if this SitesCategories
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SitesCategories) Update() *SitesCategoriesUpdateOne {
	return NewSitesCategoriesClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the SitesCategories entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SitesCategories) Unwrap() *SitesCategories {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SitesCategories is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SitesCategories) String() string {
	var builder strings.Builder
	builder.WriteString("SitesCategories(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("sites_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.SitesID))
	builder.WriteString(", ")
	builder.WriteString("categories_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.CategoriesID))
	builder.WriteByte(')')
	return builder.String()
}

// SitesCategoriesSlice is a parsable slice of SitesCategories.
type SitesCategoriesSlice []*SitesCategories
