package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Airport holds the schema definition for the Airport entity.
type Airport struct {
	ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("city"),
		field.String("country"),
		field.String("iata"),
		field.String("icao"),
		field.Float("latitude"),
		field.Float("longitude"),
		field.Float("altitude"),
		field.Float("timezone"),
		field.String("dst"),
		field.String("timezoneName"),
		field.String("type"),
		field.String("source"),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return nil
}
