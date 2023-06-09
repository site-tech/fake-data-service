// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/site-tech/fake-data-service/ent/plane"
)

// Plane is the model entity for the Plane schema.
type Plane struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// TailNumber holds the value of the "tailNumber" field.
	TailNumber string `json:"tailNumber,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Plane) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case plane.FieldID:
			values[i] = new(sql.NullInt64)
		case plane.FieldName, plane.FieldTailNumber:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Plane", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Plane fields.
func (pl *Plane) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plane.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case plane.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case plane.FieldTailNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tailNumber", values[i])
			} else if value.Valid {
				pl.TailNumber = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Plane.
// Note that you need to call Plane.Unwrap() before calling this method if this Plane
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Plane) Update() *PlaneUpdateOne {
	return NewPlaneClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Plane entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Plane) Unwrap() *Plane {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Plane is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Plane) String() string {
	var builder strings.Builder
	builder.WriteString("Plane(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("tailNumber=")
	builder.WriteString(pl.TailNumber)
	builder.WriteByte(')')
	return builder.String()
}

// Planes is a parsable slice of Plane.
type Planes []*Plane
