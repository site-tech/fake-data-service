package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Plane holds the schema definition for the Plane entity.
type Plane struct {
	ent.Schema
}

// Fields of the Plane.
func (Plane) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("tailNumber"),
	}
}

// Edges of the Plane.
func (Plane) Edges() []ent.Edge {
	return nil
}
