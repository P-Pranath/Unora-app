// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PaymentOrder holds the schema definition for the PaymentOrder entity (Razorpay orders).
type PaymentOrder struct {
	ent.Schema
}

// Fields of the PaymentOrder.
func (PaymentOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("package_id").
			MaxLen(36).
			Optional().
			Nillable(),

		// Razorpay order details
		field.String("razorpay_order_id").
			MaxLen(100).
			Unique(),

		field.String("razorpay_payment_id").
			MaxLen(100).
			Optional().
			Nillable(),

		field.String("razorpay_signature").
			MaxLen(256).
			Optional().
			Nillable(),

		// Order details
		field.Int("amount").
			Comment("Amount in smallest currency unit (paise)"),

		field.String("currency").
			MaxLen(3).
			Default("INR"),

		field.Int("credits_to_add").
			Comment("Credits to add on successful payment"),

		field.Enum("order_status").
			Values("created", "paid", "failed", "refunded", "expired").
			Default("created"),

		field.Text("failure_reason").
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("paid_at").
			Optional().
			Nillable(),
		field.Time("expires_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the PaymentOrder.
func (PaymentOrder) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "order_status"),
		index.Fields("razorpay_order_id"),
		index.Fields("order_status", "created_at"),
	}
}

// Edges of the PaymentOrder.
func (PaymentOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("payment_orders").
			Field("user_id").
			Unique().
			Required(),
	}
}
