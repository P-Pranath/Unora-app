// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserBlock holds the schema definition for the UserBlock entity.
type UserBlock struct {
	ent.Schema
}

// Fields of the UserBlock.
func (UserBlock) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("blocker_user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("blocked_user_id").
			MaxLen(36).
			NotEmpty(),

		field.Text("reason").
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("unblocked_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the UserBlock.
func (UserBlock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("blocker_user_id", "blocked_user_id").Unique(),
		index.Fields("blocker_user_id", "unblocked_at"),
		index.Fields("blocked_user_id"),
	}
}

// Edges of the UserBlock.
func (UserBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blocker", User.Type).
			Ref("blocks_given").
			Field("blocker_user_id").
			Unique().
			Required(),
		edge.From("blocked", User.Type).
			Ref("blocks_received").
			Field("blocked_user_id").
			Unique().
			Required(),
	}
}
