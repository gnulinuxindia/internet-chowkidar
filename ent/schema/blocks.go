package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Blocks struct {
	ent.Schema
}

// Fields of the Blocks.
func (Blocks) Fields() []ent.Field {
	return []ent.Field{
		field.Int("block_reports").Default(0),
		field.Int("unblock_reports").Default(0),
		field.Time("last_reported_at"),
	}
}

// Edges of the Blocks.
func (Blocks) Edges() []ent.Edge {
	return []ent.Edge{
		
	}
}

func (Blocks) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Blocks) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ip").Unique(),
	}
}
