package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Sites struct {
	ent.Schema
}

// Fields of the Sites.
func (Sites) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain").Unique(),
	}
}

// Edges of the Sites.
func (Sites) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blocks", Blocks.Type),
		edge.To("categories", Categories.Type).
			Through("sites_categories", SitesCategories.Type),
	}
}

func (Sites) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Sites) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("domain").Unique(),
	}
}
