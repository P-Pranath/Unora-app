// Package schema contains the Ent schema definitions
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserReport holds the schema definition for the UserReport entity.
type UserReport struct {
	ent.Schema
}

// Fields of the UserReport.
func (UserReport) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),

		field.String("reporter_user_id").
			MaxLen(36).
			NotEmpty(),

		field.String("reported_user_id").
			MaxLen(36).
			NotEmpty(),

		field.Enum("report_reason").
			Values(
				"inappropriate_content",
				"harassment",
				"spam",
				"fake_profile",
				"underage",
				"offensive_behavior",
				"other",
			),

		field.Text("description").
			Optional().
			Nillable(),

		// Optional reference to specific content
		field.String("reference_type").
			MaxLen(50).
			Optional().
			Nillable(),
		field.String("reference_id").
			MaxLen(36).
			Optional().
			Nillable(),

		// Admin review
		field.Enum("report_status").
			Values("pending", "reviewed", "action_taken", "dismissed").
			Default("pending"),

		field.String("reviewed_by").
			MaxLen(36).
			Optional().
			Nillable(),
		field.Text("admin_notes").
			Optional().
			Nillable(),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("reviewed_at").
			Optional().
			Nillable(),
	}
}

// Indexes of the UserReport.
func (UserReport) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("reporter_user_id", "reported_user_id", "report_reason"),
		index.Fields("reported_user_id"),
		index.Fields("report_status", "created_at"),
	}
}

// Edges of the UserReport.
func (UserReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reporter", User.Type).
			Ref("reports_given").
			Field("reporter_user_id").
			Unique().
			Required(),
		edge.From("reported", User.Type).
			Ref("reports_received").
			Field("reported_user_id").
			Unique().
			Required(),
	}
}
