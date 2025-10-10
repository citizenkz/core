package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/citizenkz/core/utils/gen"
	"github.com/google/uuid"
)

// Attempt holds the schema definition for the Attempt entity.
type Attempt struct {
	ent.Schema
}

// Fields of the Attempt.
func (Attempt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default((func() uuid.UUID)(gen.UUID())),
		field.String("otp").
			NotEmpty(),
		field.String("email").
			NotEmpty(),
	}
}

// Edges of the Attempt.
func (Attempt) Edges() []ent.Edge {
	return []ent.Edge{}
}
