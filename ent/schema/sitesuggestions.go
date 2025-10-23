package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SiteSuggestions holds the schema definition for the SiteSuggestions entity.
type SiteSuggestions struct {
	ent.Schema
}

// Fields of the SiteSuggestions.
func (SiteSuggestions) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain").Unique(),
		field.String("ping_url").Unique(),
		field.String("categories").Optional(),
		field.String("reason"),
		field.Enum("status").Values("pending", "accepted", "rejected").Default("pending"),
		field.String("resolve_reason").Optional(),
		field.Time("resolved_at").Optional(),
	}
}

// Edges of the SiteSuggestions.
func (SiteSuggestions) Edges() []ent.Edge {
	return nil
}

func (SiteSuggestions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (SiteSuggestions) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("domain").Unique(),
	}
}
