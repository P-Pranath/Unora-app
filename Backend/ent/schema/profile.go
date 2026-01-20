// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		// Identity fields
		field.String("first_name").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("last_name").
			MaxLen(50).
			Optional().
			Nillable(),
		field.Time("date_of_birth").
			Optional().
			Nillable(),
		field.String("gender").
			MaxLen(20).
			Optional().
			Nillable(),

		// Location
		field.String("city").
			MaxLen(100).
			Optional().
			Nillable(),

		// Bio and intent
		field.Text("bio").
			Optional().
			Nillable(),
		field.Text("intent_statement").
			Optional().
			Nillable(),

		// Optional fields as JSON
		field.JSON("optional_fields", map[string]interface{}{}).
			Optional(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Profile.
func (Profile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").
			Unique(),
		index.Fields("city"),
		index.Fields("date_of_birth"),
		index.Fields("deleted_at"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("profile").
			Field("user_id").
			Unique().
			Required(),
	}
}
