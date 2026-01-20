// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Nudge holds the schema definition for the Nudge entity.
type Nudge struct {
	ent.Schema
}

// Fields of the Nudge.
func (Nudge) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("streak_id").
			MaxLen(36).
			NotEmpty(),

		field.String("sender_user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("receiver_user_id").
			MaxLen(36).
			NotEmpty(),

		field.Int("day_number").
			Min(1).
			Max(15),

		field.Enum("nudge_status").
			Values("sent", "seen", "responded", "expired").
			Default("sent"),

		// Optional message
		field.String("message").
			MaxLen(200).
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("seen_at").
			Optional().
			Nillable(),
		field.Time("responded_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Nudge.
func (Nudge) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("streak_id", "day_number"),
		index.Fields("receiver_user_id", "nudge_status"),
		index.Fields("sender_user_id"),
	}
}

// Edges of the Nudge.
func (Nudge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("streak", Streak.Type).
			Ref("nudges").
			Field("streak_id").
			Unique().
			Required(),
		edge.From("sender", User.Type).
			Ref("sent_nudges").
			Field("sender_user_id").
			Unique().
			Required(),
		edge.From("receiver", User.Type).
			Ref("received_nudges").
			Field("receiver_user_id").
			Unique().
			Required(),
	}
}
