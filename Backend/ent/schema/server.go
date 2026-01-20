// Package schema contains the Ent schema definitions
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Server holds the schema definition for the Server entity (partner/friend/growth).
type Server struct {
	ent.Schema
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.Enum("server_type").
			Values("partner", "friend", "growth"),

		field.String("display_name").
			MaxLen(50).
			NotEmpty(),

		field.String("icon_name").
			MaxLen(50).
			NotEmpty(),

		field.Int("sort_order").
			Default(0),
	}
}

// Indexes of the Server.
func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_type").Unique(),
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return nil
}
