package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMixin{},
		TimeMixin{},
	}
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(3),
		field.Text("description").Optional().Nillable(),
		// field.UUID("user_id", uuid.UUID{}),
		field.Int("user_id"),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "user_id").Unique(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("projects").
			Field("user_id").
			Unique().
			Required(),
	}
}
