// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DiscoveryCard holds the schema definition for the DiscoveryCard entity.
type DiscoveryCard struct {
	ent.Schema
}

// Fields of the DiscoveryCard.
func (DiscoveryCard) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("batch_id").
			MaxLen(36).
			NotEmpty(),

		field.String("candidate_user_id").
			MaxLen(36).
			NotEmpty(),

		field.Int("display_order").
			Min(1).
			Max(5),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the DiscoveryCard.
func (DiscoveryCard) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("batch_id", "display_order"),
		index.Fields("candidate_user_id"),
		index.Fields("deleted_at"),
	}
}

// Edges of the DiscoveryCard.
func (DiscoveryCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("batch", DiscoveryBatch.Type).
			Ref("cards").
			Field("batch_id").
			Unique().
			Required(),
		edge.From("candidate", User.Type).
			Ref("discovery_appearances").
			Field("candidate_user_id").
			Unique().
			Required(),
		edge.To("interests", Interest.Type),
	}
}

