// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Reveal holds the schema definition for the Reveal entity (user's unlocked reveals).
type Reveal struct {
	ent.Schema
}

// Fields of the Reveal.
func (Reveal) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("connection_id").
			MaxLen(36).
			NotEmpty(),

		field.String("milestone_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("unlock_method").
			Values("earned", "purchased", "gifted").
			Default("earned"),

		field.Enum("reveal_status").
			Values("locked", "unlocked", "viewed").
			Default("locked"),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("unlocked_at").
			Optional().
			Nillable(),
		field.Time("viewed_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Reveal.
func (Reveal) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("connection_id", "milestone_id").Unique(),
		index.Fields("reveal_status"),
	}
}

// Edges of the Reveal.
func (Reveal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connection", Connection.Type).
			Ref("reveals").
			Field("connection_id").
			Unique().
			Required(),
		edge.From("milestone", RevealMilestone.Type).
			Ref("reveals").
			Field("milestone_id").
			Unique().
			Required(),
		edge.To("content", RevealContent.Type).
			Unique(),
	}
}
