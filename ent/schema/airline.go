package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Airline holds the schema definition for the Airline entity.
type Airline struct {
	ent.Schema
}

// Fields of the Airline.
func (Airline) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name"),
		field.String("alias"),
		field.String("country"),
		field.Bool("active"),
	}
}

// Edges of the Airline.
func (Airline) Edges() []ent.Edge {
	return nil
}
