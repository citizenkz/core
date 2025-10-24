package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChildFilter holds the schema definition for the ChildFilter entity.
type ChildFilter struct {
	ent.Schema
}

// Fields of the ChildFilter.
func (ChildFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Int("child_id"),
		field.Int("filter_id"),
		field.String("value"),
	}
}

// Edges of the ChildFilter.
func (ChildFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("child", Child.Type).
			Ref("child_filters").
			Field("child_id").
			Unique().
			Required(),
		edge.From("filter", Filter.Type).
			Ref("child_filters").
			Field("filter_id").
			Unique().
			Required(),
	}
}
