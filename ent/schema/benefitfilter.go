package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BenefitFilter holds the schema definition for the BenefitFilter entity.
type BenefitFilter struct {
	ent.Schema
}

// Fields of the BenefitFilter.
func (BenefitFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Int("benefit_id"),
		field.Int("filter_id"),
		field.String("value").
			Nillable().
			Optional(),
		field.String("from").
			Nillable().
			Optional(),
		field.String("to").
			Nillable().
			Optional(),
	}
}

// Edges of the BenefitFilter.
func (BenefitFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("benefit", Benefit.Type).
			Ref("benefit_filters").
			Field("benefit_id").
			Required().
			Unique(),
		edge.From("filter", Filter.Type).
			Ref("benefit_filters").
			Field("filter_id").
			Required().
			Unique(),
	}
}
