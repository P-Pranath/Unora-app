// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AuditLog holds the schema definition for the AuditLog entity.
type AuditLog struct {
	ent.Schema
}

// Fields of the AuditLog.
func (AuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		// Admin who performed action
		field.String("admin_identifier").
			MaxLen(100).
			NotEmpty(),

		// Action details
		field.String("action").
			MaxLen(50).
			NotEmpty(),
		field.String("resource_type").
			MaxLen(50).
			NotEmpty(),
		field.String("resource_id").
			MaxLen(36).
			Optional().
			Nillable(),

		// Request details
		field.Text("request_body").
			Optional().
			Nillable(),
		field.Text("response_summary").
			Optional().
			Nillable(),

		// Context
		field.String("ip_address").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("user_agent").
			MaxLen(255).
			Optional().
			Nillable(),

		// Status
		field.Bool("success").
			Default(true),
		field.Text("error_message").
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Indexes of the AuditLog.
func (AuditLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("admin_identifier"),
		index.Fields("action"),
		index.Fields("resource_type", "resource_id"),
		index.Fields("created_at"),
	}
}

// Edges of the AuditLog.
func (AuditLog) Edges() []ent.Edge {
	return nil
}
