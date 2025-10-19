package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Benefit holds the schema definition for the Benefit entity.
type Benefit struct {
	ent.Schema
}

// Fields of the Benefit.
func (Benefit) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Text("content"),
		field.String("bonus"),
		field.String("video_url").
			Nillable().
			Optional(),
		field.String("source_url").
			Nillable().
			Optional(),
	}
}

// Edges of the Benefit.
func (Benefit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("benefit_filters", BenefitFilter.Type),
		edge.To("benefit_categories", BenefitCategory.Type),
	}
}
