// Code generated by ent, DO NOT EDIT.

package sitescategories

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sitescategories type in the database.
	Label = "sites_categories"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSitesID holds the string denoting the sites_id field in the database.
	FieldSitesID = "sites_id"
	// FieldCategoriesID holds the string denoting the categories_id field in the database.
	FieldCategoriesID = "categories_id"
	// EdgeSites holds the string denoting the sites edge name in mutations.
	EdgeSites = "sites"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// Table holds the table name of the sitescategories in the database.
	Table = "sites_categories"
	// SitesTable is the table that holds the sites relation/edge.
	SitesTable = "sites_categories"
	// SitesInverseTable is the table name for the Sites entity.
	// It exists in this package in order to avoid circular dependency with the "sites" package.
	SitesInverseTable = "sites"
	// SitesColumn is the table column denoting the sites relation/edge.
	SitesColumn = "sites_id"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "sites_categories"
	// CategoriesInverseTable is the table name for the Categories entity.
	// It exists in this package in order to avoid circular dependency with the "categories" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "categories_id"
)

// Columns holds all SQL columns for sitescategories fields.
var Columns = []string{
	FieldID,
	FieldSitesID,
	FieldCategoriesID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the SitesCategories queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySitesID orders the results by the sites_id field.
func BySitesID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSitesID, opts...).ToFunc()
}

// ByCategoriesID orders the results by the categories_id field.
func ByCategoriesID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoriesID, opts...).ToFunc()
}

// BySitesField orders the results by sites field.
func BySitesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSitesStep(), sql.OrderByField(field, opts...))
	}
}

// ByCategoriesField orders the results by categories field.
func ByCategoriesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), sql.OrderByField(field, opts...))
	}
}
func newSitesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SitesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SitesTable, SitesColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CategoriesTable, CategoriesColumn),
	)
}