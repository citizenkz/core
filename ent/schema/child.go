package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Child holds the schema definition for the Child entity.
type Child struct {
	ent.Schema
}

// Fields of the Child.
func (Child) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("last_name"),
		field.Time("birth_date"),
		field.Int("user_id"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Child.
func (Child) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("children").
			Field("user_id").
			Unique().
			Required(),
		edge.To("child_filters", ChildFilter.Type),
	}
}
