// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Photo holds the schema definition for the Photo entity.
type Photo struct {
	ent.Schema
}

// Fields of the Photo.
func (Photo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("storage_key").
			MaxLen(255).
			NotEmpty(),

		field.Int("display_order").
			Min(1).
			Max(6),

		field.Bool("is_face_photo").
			Default(false),

		// AI validation
		field.Time("ai_validated_at").
			Optional().
			Nillable(),
		field.Bool("ai_validation_passed").
			Optional().
			Nillable(),

		// Moderation fields
		field.Bool("is_verified").
			Default(false),
		field.Bool("is_flagged").
			Default(false),
		field.String("flag_reason").
			MaxLen(500).
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Photo.
func (Photo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("deleted_at"),
	}
}

// Edges of the Photo.
func (Photo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("photos").
			Field("user_id").
			Unique().
			Required(),
	}
}
