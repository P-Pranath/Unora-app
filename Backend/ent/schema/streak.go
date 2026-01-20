// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Streak holds the schema definition for the Streak entity.
type Streak struct {
	ent.Schema
}

// Fields of the Streak.
func (Streak) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("connection_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("streak_state").
			Values("active", "at_risk", "payment_window", "reset", "completed", "terminated").
			Default("active"),

		field.Int("current_day").
			Default(1).
			Min(1).
			Max(15),

		field.Int("reset_count").
			Default(0),

		// Breaker tracking
		field.String("breaker_user_id").
			MaxLen(36).
			Optional().
			Nillable(),

		// Recovery window
		field.Time("recovery_deadline_at").
			Optional().
			Nillable(),
		field.String("recovery_payment_id").
			MaxLen(36).
			Optional().
			Nillable(),

		// AI scoring
		field.Float("streak_health_score").
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("completed_at").
			Optional().
			Nillable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Streak.
func (Streak) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("connection_id"),
		index.Fields("streak_state"),
		index.Fields("streak_state", "updated_at"),
		index.Fields("deleted_at"),
	}
}

// Edges of the Streak.
func (Streak) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connection", Connection.Type).
			Ref("streak").
			Field("connection_id").
			Unique().
			Required(),
		edge.From("breaker", User.Type).
			Ref("broken_streaks").
			Field("breaker_user_id").
			Unique(),
		edge.To("check_ins", CheckIn.Type),
		edge.To("nudges", Nudge.Type),
	}
}
