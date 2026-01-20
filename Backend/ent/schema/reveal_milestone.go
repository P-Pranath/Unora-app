// Package schema contains the Ent schema definitions
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RevealMilestone holds the schema definition for the RevealMilestone entity (master data).
type RevealMilestone struct {
	ent.Schema
}

// Fields of the RevealMilestone.
func (RevealMilestone) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.Int("reveal_number").
			Min(1).
			Max(3),

		field.Int("day_required").
			Min(1).
			Max(15),

		field.Enum("reveal_type").
			Values("personality", "values", "lifestyle"),

		field.String("title").
			MaxLen(100).
			NotEmpty(),

		field.Text("description").
			Optional().
			Nillable(),

		field.String("icon_name").
			MaxLen(50).
			Optional().
			Nillable(),

		field.Int("credit_cost").
			Default(0),

		field.Bool("is_active").
			Default(true),
	}
}

// Indexes of the RevealMilestone.
func (RevealMilestone) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("reveal_number").Unique(),
		index.Fields("day_required"),
	}
}

// Edges of the RevealMilestone.
func (RevealMilestone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reveals", Reveal.Type),
	}
}
