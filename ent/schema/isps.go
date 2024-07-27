package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Isps struct {
	ent.Schema
}

// Fields of the Isps.
func (Isps) Fields() []ent.Field {
	return []ent.Field{
		field.Float("latitude"),
		field.Float("longitude"),
		field.String("name"),
	}
}

// Edges of the Isps.
func (Isps) Edges() []ent.Edge {
	return nil
}

func (Isps) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
