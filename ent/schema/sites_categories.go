package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SitesCategories struct {
	ent.Schema
}

// Fields of the SitesCategories.
func (SitesCategories) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sites_id"),
		field.Int("categories_id"),
	}
}

// Edges of the SitesCategories.
func (SitesCategories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sites", Sites.Type).
			Required().
			Unique().
			Field("sites_id"),
		edge.To("categories", Categories.Type).
			Required().
			Unique().
			Field("categories_id"),
	}
}
