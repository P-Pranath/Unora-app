// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CreditTransaction holds the schema definition for the CreditTransaction entity.
type CreditTransaction struct {
	ent.Schema
}

// Fields of the CreditTransaction.
func (CreditTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("transaction_type").
			Values("purchase", "streak_recovery", "early_reveal", "referral_bonus", "welcome_bonus", "refund", "admin_adjustment"),

		field.Int("credit_amount").
			Comment("Positive for credits added, negative for credits spent"),

		field.Int("balance_after").
			Comment("User's credit balance after this transaction"),

		// Reference to what triggered this transaction
		field.String("reference_type").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("reference_id").
			MaxLen(36).
			Optional().
			Nillable(),

		field.Text("description").
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Indexes of the CreditTransaction.
func (CreditTransaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "created_at"),
		index.Fields("transaction_type"),
	}
}

// Edges of the CreditTransaction.
func (CreditTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("credit_transactions").
			Field("user_id").
			Unique().
			Required(),
	}
}
