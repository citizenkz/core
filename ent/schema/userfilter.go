package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserFilter holds the schema definition for the UserFilter entity.
type UserFilter struct {
	ent.Schema
}

// Fields of the UserFilter.
func (UserFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("filter_id"),
		field.String("value"),
	}
}

// Edges of the UserFilter.
func (UserFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_filters").
			Field("user_id").
			Required().
			Unique(),
		edge.From("filter", Filter.Type).
			Ref("user_filters").
			Field("filter_id").
			Required().
			Unique(),
	}
}
