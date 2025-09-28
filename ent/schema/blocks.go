package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Blocks struct {
	ent.Schema
}

// Fields of the Blocks.
func (Blocks) Fields() []ent.Field {
	return []ent.Field{
		field.Int("site_id"),
		field.Int("isp_id"),
		field.Int("client_id"),
		field.Bool("blocked"),
		field.Time("last_reported_at"),
	}
}

// Edges of the Blocks.
func (Blocks) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("site", Sites.Type).
			Ref("blocks").
			Field("site_id").
			Required().
			Unique(),
		edge.From("isp", Isps.Type).
			Ref("isp_blocks").
			Field("isp_id").
			Required().
			Unique(),
	}
}

func (Blocks) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Blocks) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("site_id", "isp_id").Unique(),
	}
}
