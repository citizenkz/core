package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("first_name").
            NotEmpty().
            MaxLen(100),

        field.String("last_name").
            NotEmpty().
            MaxLen(100),

        field.Time("birth_date").
            Optional(),

        field.String("email").
            Unique().
            NotEmpty(),

        field.String("password").
            Sensitive(),

		field.Time("created_at").
			Default(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("user_filters", UserFilter.Type),
	}
}
