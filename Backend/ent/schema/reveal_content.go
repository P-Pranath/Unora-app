// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RevealContent holds the schema definition for the RevealContent entity.
type RevealContent struct {
	ent.Schema
}

// Fields of the RevealContent.
func (RevealContent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("reveal_id").
			MaxLen(36).
			NotEmpty(),

		// AI-generated content
		field.Text("ai_summary").
			Optional().
			Nillable(),

		field.Text("compatibility_insight").
			Optional().
			Nillable(),

		field.Text("conversation_starters").
			Optional().
			Nillable(),

		// Personality dimension scores (JSON)
		field.JSON("dimension_scores", map[string]interface{}{}).
			Optional(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Indexes of the RevealContent.
func (RevealContent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("reveal_id").Unique(),
	}
}

// Edges of the RevealContent.
func (RevealContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reveal", Reveal.Type).
			Ref("content").
			Field("reveal_id").
			Unique().
			Required(),
	}
}
