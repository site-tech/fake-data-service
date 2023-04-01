// Code generated by ent, DO NOT EDIT.

package route

import (
	"entgo.io/ent/dialect/sql"
	"github.com/site-tech/fake-data-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldID, id))
}

// Airline applies equality check predicate on the "airline" field. It's identical to AirlineEQ.
func Airline(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldAirline, v))
}

// SourceAirportId applies equality check predicate on the "sourceAirportId" field. It's identical to SourceAirportIdEQ.
func SourceAirportId(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldSourceAirportId, v))
}

// DestinationAirportId applies equality check predicate on the "destinationAirportId" field. It's identical to DestinationAirportIdEQ.
func DestinationAirportId(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldDestinationAirportId, v))
}

// NumberOfStops applies equality check predicate on the "numberOfStops" field. It's identical to NumberOfStopsEQ.
func NumberOfStops(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldNumberOfStops, v))
}

// AirlineEQ applies the EQ predicate on the "airline" field.
func AirlineEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldAirline, v))
}

// AirlineNEQ applies the NEQ predicate on the "airline" field.
func AirlineNEQ(v string) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldAirline, v))
}

// AirlineIn applies the In predicate on the "airline" field.
func AirlineIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldAirline, vs...))
}

// AirlineNotIn applies the NotIn predicate on the "airline" field.
func AirlineNotIn(vs ...string) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldAirline, vs...))
}

// AirlineGT applies the GT predicate on the "airline" field.
func AirlineGT(v string) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldAirline, v))
}

// AirlineGTE applies the GTE predicate on the "airline" field.
func AirlineGTE(v string) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldAirline, v))
}

// AirlineLT applies the LT predicate on the "airline" field.
func AirlineLT(v string) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldAirline, v))
}

// AirlineLTE applies the LTE predicate on the "airline" field.
func AirlineLTE(v string) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldAirline, v))
}

// AirlineContains applies the Contains predicate on the "airline" field.
func AirlineContains(v string) predicate.Route {
	return predicate.Route(sql.FieldContains(FieldAirline, v))
}

// AirlineHasPrefix applies the HasPrefix predicate on the "airline" field.
func AirlineHasPrefix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasPrefix(FieldAirline, v))
}

// AirlineHasSuffix applies the HasSuffix predicate on the "airline" field.
func AirlineHasSuffix(v string) predicate.Route {
	return predicate.Route(sql.FieldHasSuffix(FieldAirline, v))
}

// AirlineEqualFold applies the EqualFold predicate on the "airline" field.
func AirlineEqualFold(v string) predicate.Route {
	return predicate.Route(sql.FieldEqualFold(FieldAirline, v))
}

// AirlineContainsFold applies the ContainsFold predicate on the "airline" field.
func AirlineContainsFold(v string) predicate.Route {
	return predicate.Route(sql.FieldContainsFold(FieldAirline, v))
}

// SourceAirportIdEQ applies the EQ predicate on the "sourceAirportId" field.
func SourceAirportIdEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldSourceAirportId, v))
}

// SourceAirportIdNEQ applies the NEQ predicate on the "sourceAirportId" field.
func SourceAirportIdNEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldSourceAirportId, v))
}

// SourceAirportIdIn applies the In predicate on the "sourceAirportId" field.
func SourceAirportIdIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldSourceAirportId, vs...))
}

// SourceAirportIdNotIn applies the NotIn predicate on the "sourceAirportId" field.
func SourceAirportIdNotIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldSourceAirportId, vs...))
}

// SourceAirportIdGT applies the GT predicate on the "sourceAirportId" field.
func SourceAirportIdGT(v int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldSourceAirportId, v))
}

// SourceAirportIdGTE applies the GTE predicate on the "sourceAirportId" field.
func SourceAirportIdGTE(v int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldSourceAirportId, v))
}

// SourceAirportIdLT applies the LT predicate on the "sourceAirportId" field.
func SourceAirportIdLT(v int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldSourceAirportId, v))
}

// SourceAirportIdLTE applies the LTE predicate on the "sourceAirportId" field.
func SourceAirportIdLTE(v int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldSourceAirportId, v))
}

// DestinationAirportIdEQ applies the EQ predicate on the "destinationAirportId" field.
func DestinationAirportIdEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldDestinationAirportId, v))
}

// DestinationAirportIdNEQ applies the NEQ predicate on the "destinationAirportId" field.
func DestinationAirportIdNEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldDestinationAirportId, v))
}

// DestinationAirportIdIn applies the In predicate on the "destinationAirportId" field.
func DestinationAirportIdIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldDestinationAirportId, vs...))
}

// DestinationAirportIdNotIn applies the NotIn predicate on the "destinationAirportId" field.
func DestinationAirportIdNotIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldDestinationAirportId, vs...))
}

// DestinationAirportIdGT applies the GT predicate on the "destinationAirportId" field.
func DestinationAirportIdGT(v int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldDestinationAirportId, v))
}

// DestinationAirportIdGTE applies the GTE predicate on the "destinationAirportId" field.
func DestinationAirportIdGTE(v int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldDestinationAirportId, v))
}

// DestinationAirportIdLT applies the LT predicate on the "destinationAirportId" field.
func DestinationAirportIdLT(v int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldDestinationAirportId, v))
}

// DestinationAirportIdLTE applies the LTE predicate on the "destinationAirportId" field.
func DestinationAirportIdLTE(v int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldDestinationAirportId, v))
}

// NumberOfStopsEQ applies the EQ predicate on the "numberOfStops" field.
func NumberOfStopsEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldEQ(FieldNumberOfStops, v))
}

// NumberOfStopsNEQ applies the NEQ predicate on the "numberOfStops" field.
func NumberOfStopsNEQ(v int) predicate.Route {
	return predicate.Route(sql.FieldNEQ(FieldNumberOfStops, v))
}

// NumberOfStopsIn applies the In predicate on the "numberOfStops" field.
func NumberOfStopsIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldIn(FieldNumberOfStops, vs...))
}

// NumberOfStopsNotIn applies the NotIn predicate on the "numberOfStops" field.
func NumberOfStopsNotIn(vs ...int) predicate.Route {
	return predicate.Route(sql.FieldNotIn(FieldNumberOfStops, vs...))
}

// NumberOfStopsGT applies the GT predicate on the "numberOfStops" field.
func NumberOfStopsGT(v int) predicate.Route {
	return predicate.Route(sql.FieldGT(FieldNumberOfStops, v))
}

// NumberOfStopsGTE applies the GTE predicate on the "numberOfStops" field.
func NumberOfStopsGTE(v int) predicate.Route {
	return predicate.Route(sql.FieldGTE(FieldNumberOfStops, v))
}

// NumberOfStopsLT applies the LT predicate on the "numberOfStops" field.
func NumberOfStopsLT(v int) predicate.Route {
	return predicate.Route(sql.FieldLT(FieldNumberOfStops, v))
}

// NumberOfStopsLTE applies the LTE predicate on the "numberOfStops" field.
func NumberOfStopsLTE(v int) predicate.Route {
	return predicate.Route(sql.FieldLTE(FieldNumberOfStops, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Route) predicate.Route {
	return predicate.Route(func(s *sql.Selector) {
		p(s.Not())
	})
}
