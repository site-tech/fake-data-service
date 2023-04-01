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
		field.String("name").Optional(),
		field.String("city").Optional(),
		field.String("country").Optional(),
		field.String("iata").Optional(),
		field.String("icao").Optional(),
		field.Float("latitude").Optional(),
		field.Float("longitude").Optional(),
		field.Int("altitude").Optional(),
		field.String("timezone").Optional(),
		field.String("dst").Optional(),
		field.String("timezoneName").Optional(),
		field.String("type").Optional(),
		field.String("source").Optional(),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return nil
}
