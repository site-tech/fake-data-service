// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AirportsColumns holds the columns for the "airports" table.
	AirportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "iata", Type: field.TypeString},
		{Name: "icao", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "altitude", Type: field.TypeFloat64},
		{Name: "timezone", Type: field.TypeFloat64},
		{Name: "dst", Type: field.TypeString},
		{Name: "timezone_name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
	}
	// AirportsTable holds the schema information for the "airports" table.
	AirportsTable = &schema.Table{
		Name:       "airports",
		Columns:    AirportsColumns,
		PrimaryKey: []*schema.Column{AirportsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AirportsTable,
	}
)

func init() {
}
