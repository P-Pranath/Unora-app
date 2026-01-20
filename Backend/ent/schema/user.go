// Package schema contains the Ent schema definitions for the Unora application
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// Primary key
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		// Auth fields (supports both email/OAuth and phone)
		field.String("email").
			MaxLen(255).
			Optional().
			Nillable(),
		field.String("phone_number").
			MaxLen(20).
			Optional().
			Nillable(),
		field.String("phone_country_code").
			MaxLen(5).
			Default("+91").
			Optional(),

		// OAuth provider info
		field.String("provider").
			MaxLen(20).
			Optional().
			Nillable(),
		field.String("provider_user_id").
			MaxLen(255).
			Optional().
			Nillable(),

		// Basic profile (from OAuth or user input)
		field.String("name").
			MaxLen(100).
			Optional().
			Nillable(),
		field.String("first_name").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("last_name").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("picture").
			MaxLen(500).
			Optional().
			Nillable(),

		// Extended profile (PRD required)
		field.Time("date_of_birth").
			Optional().
			Nillable(),
		field.Enum("gender").
			Values("male", "female", "non_binary", "prefer_not_to_say").
			Optional().
			Nillable(),
		field.String("city").
			MaxLen(100).
			Optional().
			Nillable(),
		field.String("education").
			MaxLen(100).
			Optional().
			Nillable(),
		field.String("profession").
			MaxLen(100).
			Optional().
			Nillable(),
		field.String("religion").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("bio").
			MaxLen(500).
			Optional().
			Nillable(),

		// Verification and tier
		field.Enum("verification_status").
			Values("pending", "phone_verified", "photos_submitted", "photo_quality_verified", "gov_id_verified").
			Default("pending"),
		field.Enum("subscription_tier").
			Values("free", "plus", "pro").
			Default("free"),

		// Tier tracking
		field.Int("free_recoveries_used").
			Default(0),
		field.Int("nudges_sent_today").
			Default(0),
		field.Time("nudges_reset_at").
			Optional().
			Nillable(),

		// Global inventory (spans all servers)
		field.Int("active_connection_count").
			Default(0).
			NonNegative(),
		field.Time("last_global_refresh_at").
			Optional().
			Nillable(),
		field.Time("refresh_available_at").
			Optional().
			Nillable(),

		// Credit balance
		field.Int("credit_balance").
			Default(0),

		// Behavioral trust (computed by AI, stored for queries)
		field.Float("behavioral_trust_score").
			Optional().
			Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "DECIMAL(3,2)",
			}),

		// Account status and activity
		field.Enum("account_status").
			Values("active", "suspended", "deleted", "pending_verification").
			Default("active"),
		field.Enum("onboarding_status").
			Values("started", "profile_basic", "photos_added", "hobbies_added", "completed").
			Default("started"),
		field.Time("last_active_at").
			Optional().
			Nillable(),
		field.Time("suspended_at").
			Optional().
			Nillable(),
		field.String("suspension_reason").
			MaxLen(500).
			Optional().
			Nillable(),

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

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
		index.Fields("phone_number"),
		index.Fields("provider", "provider_user_id"),
		index.Fields("deleted_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Profile module
		edge.To("profile", Profile.Type).
			Unique(),
		edge.To("photos", Photo.Type),
		edge.To("hobbies", Hobby.Type),

		// Discovery module
		edge.To("filters", Filter.Type),
		edge.To("discovery_batches", DiscoveryBatch.Type),
		edge.To("discovery_appearances", DiscoveryCard.Type),

		// Matching module
		edge.To("sent_interests", Interest.Type),
		edge.To("received_interests", Interest.Type),
		edge.To("connections_as_a", Connection.Type),
		edge.To("connections_as_b", Connection.Type),

		// Streak module
		edge.To("broken_streaks", Streak.Type),
		edge.To("check_ins", CheckIn.Type),
		edge.To("sent_nudges", Nudge.Type),
		edge.To("received_nudges", Nudge.Type),

		// Monetization module
		edge.To("credit_transactions", CreditTransaction.Type),
		edge.To("payment_orders", PaymentOrder.Type),

		// Safety module
		edge.To("blocks_given", UserBlock.Type),
		edge.To("blocks_received", UserBlock.Type),
		edge.To("reports_given", UserReport.Type),
		edge.To("reports_received", UserReport.Type),
	}
}

