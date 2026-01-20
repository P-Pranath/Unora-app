// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Connection holds the schema definition for the Connection entity.
type Connection struct {
	ent.Schema
}

// Fields of the Connection.
func (Connection) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		// User A ID is always < User B ID to prevent duplicates
		field.String("user_a_id").
			MaxLen(36).
			NotEmpty(),

		field.String("user_b_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("server_type").
			Values("partner", "friend", "growth"),

		field.Enum("connection_status").
			Values("active", "terminated").
			Default("active"),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("terminated_at").
			Optional().
			Nillable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Connection.
func (Connection) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_a_id", "connection_status"),
		index.Fields("user_b_id", "connection_status"),
		index.Fields("user_a_id", "user_b_id", "server_type").Unique(),
		index.Fields("deleted_at"),
	}
}

// Edges of the Connection.
func (Connection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_a", User.Type).
			Ref("connections_as_a").
			Field("user_a_id").
			Unique().
			Required(),
		edge.From("user_b", User.Type).
			Ref("connections_as_b").
			Field("user_b_id").
			Unique().
			Required(),
		edge.To("streak", Streak.Type).
			Unique(),
		edge.To("reveals", Reveal.Type),
	}
}
