// Package schema contains the Ent schema definitions
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// HobbyOption holds the schema definition for the HobbyOption entity (master list).
type HobbyOption struct {
	ent.Schema
}

// Fields of the HobbyOption.
func (HobbyOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("name").
			MaxLen(100).
			NotEmpty(),

		field.String("category").
			MaxLen(50).
			NotEmpty(),

		field.String("icon_name").
			MaxLen(50).
			Optional().
			Nillable(),

		field.JSON("micro_descriptions", []string{}).
			Optional(),

		field.Bool("is_active").
			Default(true),

		field.Int("sort_order").
			Default(0),
	}
}

// Indexes of the HobbyOption.
func (HobbyOption) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
		index.Fields("is_active"),
	}
}

// Edges of the HobbyOption.
func (HobbyOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hobbies", Hobby.Type),
	}
}
