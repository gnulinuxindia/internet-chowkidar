package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sites", Sites.Type).
			Ref("categories").
			Through("sites_categories", SitesCategories.Type),
	}
}

func (Categories) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Categories) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
