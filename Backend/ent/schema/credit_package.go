// Package schema contains the Ent schema definitions
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CreditPackage holds the schema definition for the CreditPackage entity (purchasable credit packages).
type CreditPackage struct {
	ent.Schema
}

// Fields of the CreditPackage.
func (CreditPackage) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("name").
			MaxLen(100).
			NotEmpty(),

		field.Text("description").
			Optional().
			Nillable(),

		field.Int("credit_amount").
			Min(1),

		field.Int("bonus_credits").
			Default(0),

		// Price in smallest currency unit (paise for INR)
		field.Int("price_amount").
			Min(100),

		field.String("currency").
			MaxLen(3).
			Default("INR"),

		field.Float("discount_percent").
			Default(0),

		field.String("badge_text").
			MaxLen(50).
			Optional().
			Nillable(),

		field.Bool("is_popular").
			Default(false),

		field.Bool("is_active").
			Default(true),

		field.Int("sort_order").
			Default(0),
	}
}

// Indexes of the CreditPackage.
func (CreditPackage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_active", "sort_order"),
	}
}

// Edges of the CreditPackage.
func (CreditPackage) Edges() []ent.Edge {
	return nil
}
