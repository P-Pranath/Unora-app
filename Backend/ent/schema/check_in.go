// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CheckIn holds the schema definition for the CheckIn entity.
type CheckIn struct {
	ent.Schema
}

// Fields of the CheckIn.
func (CheckIn) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("streak_id").
			MaxLen(36).
			NotEmpty(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		field.Int("day_number").
			Min(1).
			Max(15),

		// Check-in type
		field.Enum("check_in_type").
			Values("manual", "nudge_response", "auto").
			Default("manual"),

		// Optional event data (what activity they did)
		field.JSON("event_data", map[string]interface{}{}).
			Optional(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Indexes of the CheckIn.
func (CheckIn) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("streak_id", "day_number", "user_id").Unique(),
		index.Fields("user_id", "created_at"),
	}
}

// Edges of the CheckIn.
func (CheckIn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("streak", Streak.Type).
			Ref("check_ins").
			Field("streak_id").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("check_ins").
			Field("user_id").
			Unique().
			Required(),
	}
}
