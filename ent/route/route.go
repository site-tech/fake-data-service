// Code generated by ent, DO NOT EDIT.

package route

const (
	// Label holds the string label denoting the route type in the database.
	Label = "route"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAirlineId holds the string denoting the airlineid field in the database.
	FieldAirlineId = "airline_id"
	// FieldSourceAirportId holds the string denoting the sourceairportid field in the database.
	FieldSourceAirportId = "source_airport_id"
	// FieldDestinationAirportId holds the string denoting the destinationairportid field in the database.
	FieldDestinationAirportId = "destination_airport_id"
	// FieldPlaneId holds the string denoting the planeid field in the database.
	FieldPlaneId = "plane_id"
	// FieldNumberOfStops holds the string denoting the numberofstops field in the database.
	FieldNumberOfStops = "number_of_stops"
	// Table holds the table name of the route in the database.
	Table = "routes"
)

// Columns holds all SQL columns for route fields.
var Columns = []string{
	FieldID,
	FieldAirlineId,
	FieldSourceAirportId,
	FieldDestinationAirportId,
	FieldPlaneId,
	FieldNumberOfStops,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
