// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DiscoveryBatch holds the schema definition for the DiscoveryBatch entity.
type DiscoveryBatch struct {
	ent.Schema
}

// Fields of the DiscoveryBatch.
func (DiscoveryBatch) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("server_type").
			Values("partner", "friend", "growth"),

		field.Enum("batch_status").
			Values("active", "consumed", "expired").
			Default("active"),

		// Snapshot of filter at time of batch creation
		field.JSON("filter_snapshot", map[string]interface{}{}).
			Optional(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("expires_at").
			Optional().
			Nillable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the DiscoveryBatch.
func (DiscoveryBatch) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "server_type", "batch_status"),
		index.Fields("deleted_at"),
	}
}

// Edges of the DiscoveryBatch.
func (DiscoveryBatch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("discovery_batches").
			Field("user_id").
			Unique().
			Required(),
		edge.To("cards", DiscoveryCard.Type),
	}
}
