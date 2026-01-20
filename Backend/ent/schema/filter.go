// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Filter holds the schema definition for the Filter entity (user's saved filters per server).
type Filter struct {
	ent.Schema
}

// Fields of the Filter.
func (Filter) Fields() []ent.Field {
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

		// Universal filters
		field.Int("min_age").
			Optional().
			Nillable(),
		field.Int("max_age").
			Optional().
			Nillable(),

		// Partner server filters
		field.Enum("gender_preference").
			Values("male", "female", "non_binary", "any").
			Optional().
			Nillable(),
		field.Enum("relationship_intent").
			Values("exploring", "open_to_commitment", "seeking_committed").
			Optional().
			Nillable(),
		field.Enum("family_planning").
			Values("wants_children", "no_children", "undecided", "prefer_not_to_say").
			Optional().
			Nillable(),
		field.Enum("living_situation").
			Values("alone", "with_family", "with_roommates", "flexible").
			Optional().
			Nillable(),
		field.Enum("dietary_preference").
			Values("vegetarian", "non_vegetarian", "vegan", "no_preference").
			Optional().
			Nillable(),

		// Friend server filters
		field.Enum("friendship_style").
			Values("activity_based", "conversation_based", "support_based", "all_rounder").
			Optional().
			Nillable(),
		field.Enum("social_energy").
			Values("introvert", "extrovert", "ambivert").
			Optional().
			Nillable(),
		field.Enum("hangout_preference").
			Values("in_person", "virtual", "mixed").
			Optional().
			Nillable(),
		field.Enum("conversation_depth").
			Values("light_and_fun", "deep_discussions", "both").
			Optional().
			Nillable(),

		// Growth server filters
		field.Enum("goal_category").
			Values("fitness", "career", "learning", "creative", "financial", "mental_health", "habit_building", "other").
			Optional().
			Nillable(),
		field.Enum("accountability_style").
			Values("gentle", "direct", "structured").
			Optional().
			Nillable(),
		field.Enum("check_in_frequency").
			Values("daily", "every_few_days", "weekly").
			Optional().
			Nillable(),
		field.Enum("commitment_level").
			Values("experimenting", "moderately_committed", "fully_dedicated").
			Optional().
			Nillable(),

		// Legacy JSON config (for any additional filters)
		field.JSON("filter_config", map[string]interface{}{}).
			Optional(),

		field.Bool("is_pending").
			Default(false),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Indexes of the Filter.
func (Filter) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "server_type").Unique(),
	}
}

// Edges of the Filter.
func (Filter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("filters").
			Field("user_id").
			Unique().
			Required(),
	}
}
