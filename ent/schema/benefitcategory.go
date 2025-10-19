package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BenefitCategory holds the schema definition for the BenefitCategory entity.
type BenefitCategory struct {
	ent.Schema
}

// Fields of the BenefitCategory.
func (BenefitCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("benefit_id"),
		field.Int("category_id"),
	}
}

// Edges of the BenefitCategory.
func (BenefitCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("benefit", Benefit.Type).
			Ref("benefit_categories").
			Field("benefit_id").
			Required().
			Unique(),
		edge.From("category", Category.Type).
			Ref("benefit_categories").
			Field("category_id").
			Required().
			Unique(),
	}
}
