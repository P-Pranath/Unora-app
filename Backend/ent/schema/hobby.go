// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Hobby holds the schema definition for the Hobby entity (user's selected hobbies).
type Hobby struct {
	ent.Schema
}

// Fields of the Hobby.
func (Hobby) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("hobby_option_id").
			MaxLen(36).
			NotEmpty(),

		field.String("micro_description").
			MaxLen(200).
			Optional().
			Nillable(),

		field.Int("display_order").
			Default(0),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the Hobby.
func (Hobby) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "hobby_option_id").
			Unique(),
		index.Fields("deleted_at"),
	}
}

// Edges of the Hobby.
func (Hobby) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("hobbies").
			Field("user_id").
			Unique().
			Required(),
		edge.From("hobby_option", HobbyOption.Type).
			Ref("hobbies").
			Field("hobby_option_id").
			Unique().
			Required(),
	}
}
