// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Interest holds the schema definition for the Interest entity.
type Interest struct {
	ent.Schema
}

// Fields of the Interest.
func (Interest) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("sender_user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("receiver_user_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("server_type").
			Values("partner", "friend", "growth"),

		field.String("discovery_card_id").
			MaxLen(36).
			Optional().
			Nillable(),

		field.Enum("interest_status").
			Values("pending", "matched", "expired", "wiped").
			Default("pending"),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("matched_at").
			Optional().
			Nillable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Interest.
func (Interest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sender_user_id", "interest_status"),
		index.Fields("receiver_user_id"),
		index.Fields("deleted_at"),
	}
}

// Edges of the Interest.
func (Interest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).
			Ref("sent_interests").
			Field("sender_user_id").
			Unique().
			Required(),
		edge.From("receiver", User.Type).
			Ref("received_interests").
			Field("receiver_user_id").
			Unique().
			Required(),
		edge.From("discovery_card", DiscoveryCard.Type).
			Ref("interests").
			Field("discovery_card_id").
			Unique(),
	}
}
