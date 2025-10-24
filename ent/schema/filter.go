package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/citizenkz/core/services/filter/consts"
)

// Filter holds the schema definition for the Filter entity.
type Filter struct {
	ent.Schema
}

// Fields of the Filter.
func (Filter) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("hint").
			Nillable().
			Optional(),
		field.Enum("type").
			Values(
				consts.DateRange.String(),
				consts.NumberRange.String(),
				consts.StringRange.String(),
			),
		field.JSON("values", []string{}),
	}
}

// Edges of the Filter.
func (Filter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_filters", UserFilter.Type),
		edge.To("benefit_filters", BenefitFilter.Type),
		edge.To("child_filters", ChildFilter.Type),
	}
}
