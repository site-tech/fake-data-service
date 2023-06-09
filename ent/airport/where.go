// Code generated by ent, DO NOT EDIT.

package airport

import (
	"entgo.io/ent/dialect/sql"
	"github.com/site-tech/fake-data-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldName, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldCity, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldCountry, v))
}

// Iata applies equality check predicate on the "iata" field. It's identical to IataEQ.
func Iata(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIata, v))
}

// Icao applies equality check predicate on the "icao" field. It's identical to IcaoEQ.
func Icao(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIcao, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLongitude, v))
}

// Altitude applies equality check predicate on the "altitude" field. It's identical to AltitudeEQ.
func Altitude(v int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldAltitude, v))
}

// Timezone applies equality check predicate on the "timezone" field. It's identical to TimezoneEQ.
func Timezone(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldTimezone, v))
}

// Dst applies equality check predicate on the "dst" field. It's identical to DstEQ.
func Dst(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldDst, v))
}

// TimezoneName applies equality check predicate on the "timezoneName" field. It's identical to TimezoneNameEQ.
func TimezoneName(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldTimezoneName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldType, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldSource, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldName, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldCity, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldCountry))
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldCountry))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldCountry, v))
}

// IataEQ applies the EQ predicate on the "iata" field.
func IataEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIata, v))
}

// IataNEQ applies the NEQ predicate on the "iata" field.
func IataNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldIata, v))
}

// IataIn applies the In predicate on the "iata" field.
func IataIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldIata, vs...))
}

// IataNotIn applies the NotIn predicate on the "iata" field.
func IataNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldIata, vs...))
}

// IataGT applies the GT predicate on the "iata" field.
func IataGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldIata, v))
}

// IataGTE applies the GTE predicate on the "iata" field.
func IataGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldIata, v))
}

// IataLT applies the LT predicate on the "iata" field.
func IataLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldIata, v))
}

// IataLTE applies the LTE predicate on the "iata" field.
func IataLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldIata, v))
}

// IataContains applies the Contains predicate on the "iata" field.
func IataContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldIata, v))
}

// IataHasPrefix applies the HasPrefix predicate on the "iata" field.
func IataHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldIata, v))
}

// IataHasSuffix applies the HasSuffix predicate on the "iata" field.
func IataHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldIata, v))
}

// IataIsNil applies the IsNil predicate on the "iata" field.
func IataIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldIata))
}

// IataNotNil applies the NotNil predicate on the "iata" field.
func IataNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldIata))
}

// IataEqualFold applies the EqualFold predicate on the "iata" field.
func IataEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldIata, v))
}

// IataContainsFold applies the ContainsFold predicate on the "iata" field.
func IataContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldIata, v))
}

// IcaoEQ applies the EQ predicate on the "icao" field.
func IcaoEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIcao, v))
}

// IcaoNEQ applies the NEQ predicate on the "icao" field.
func IcaoNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldIcao, v))
}

// IcaoIn applies the In predicate on the "icao" field.
func IcaoIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldIcao, vs...))
}

// IcaoNotIn applies the NotIn predicate on the "icao" field.
func IcaoNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldIcao, vs...))
}

// IcaoGT applies the GT predicate on the "icao" field.
func IcaoGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldIcao, v))
}

// IcaoGTE applies the GTE predicate on the "icao" field.
func IcaoGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldIcao, v))
}

// IcaoLT applies the LT predicate on the "icao" field.
func IcaoLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldIcao, v))
}

// IcaoLTE applies the LTE predicate on the "icao" field.
func IcaoLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldIcao, v))
}

// IcaoContains applies the Contains predicate on the "icao" field.
func IcaoContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldIcao, v))
}

// IcaoHasPrefix applies the HasPrefix predicate on the "icao" field.
func IcaoHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldIcao, v))
}

// IcaoHasSuffix applies the HasSuffix predicate on the "icao" field.
func IcaoHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldIcao, v))
}

// IcaoIsNil applies the IsNil predicate on the "icao" field.
func IcaoIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldIcao))
}

// IcaoNotNil applies the NotNil predicate on the "icao" field.
func IcaoNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldIcao))
}

// IcaoEqualFold applies the EqualFold predicate on the "icao" field.
func IcaoEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldIcao, v))
}

// IcaoContainsFold applies the ContainsFold predicate on the "icao" field.
func IcaoContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldIcao, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldLatitude, v))
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldLatitude))
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldLatitude))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldLongitude, v))
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldLongitude))
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldLongitude))
}

// AltitudeEQ applies the EQ predicate on the "altitude" field.
func AltitudeEQ(v int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldAltitude, v))
}

// AltitudeNEQ applies the NEQ predicate on the "altitude" field.
func AltitudeNEQ(v int) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldAltitude, v))
}

// AltitudeIn applies the In predicate on the "altitude" field.
func AltitudeIn(vs ...int) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldAltitude, vs...))
}

// AltitudeNotIn applies the NotIn predicate on the "altitude" field.
func AltitudeNotIn(vs ...int) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldAltitude, vs...))
}

// AltitudeGT applies the GT predicate on the "altitude" field.
func AltitudeGT(v int) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldAltitude, v))
}

// AltitudeGTE applies the GTE predicate on the "altitude" field.
func AltitudeGTE(v int) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldAltitude, v))
}

// AltitudeLT applies the LT predicate on the "altitude" field.
func AltitudeLT(v int) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldAltitude, v))
}

// AltitudeLTE applies the LTE predicate on the "altitude" field.
func AltitudeLTE(v int) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldAltitude, v))
}

// AltitudeIsNil applies the IsNil predicate on the "altitude" field.
func AltitudeIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldAltitude))
}

// AltitudeNotNil applies the NotNil predicate on the "altitude" field.
func AltitudeNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldAltitude))
}

// TimezoneEQ applies the EQ predicate on the "timezone" field.
func TimezoneEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldTimezone, v))
}

// TimezoneNEQ applies the NEQ predicate on the "timezone" field.
func TimezoneNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldTimezone, v))
}

// TimezoneIn applies the In predicate on the "timezone" field.
func TimezoneIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldTimezone, vs...))
}

// TimezoneNotIn applies the NotIn predicate on the "timezone" field.
func TimezoneNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldTimezone, vs...))
}

// TimezoneGT applies the GT predicate on the "timezone" field.
func TimezoneGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldTimezone, v))
}

// TimezoneGTE applies the GTE predicate on the "timezone" field.
func TimezoneGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldTimezone, v))
}

// TimezoneLT applies the LT predicate on the "timezone" field.
func TimezoneLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldTimezone, v))
}

// TimezoneLTE applies the LTE predicate on the "timezone" field.
func TimezoneLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldTimezone, v))
}

// TimezoneContains applies the Contains predicate on the "timezone" field.
func TimezoneContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldTimezone, v))
}

// TimezoneHasPrefix applies the HasPrefix predicate on the "timezone" field.
func TimezoneHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldTimezone, v))
}

// TimezoneHasSuffix applies the HasSuffix predicate on the "timezone" field.
func TimezoneHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldTimezone, v))
}

// TimezoneIsNil applies the IsNil predicate on the "timezone" field.
func TimezoneIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldTimezone))
}

// TimezoneNotNil applies the NotNil predicate on the "timezone" field.
func TimezoneNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldTimezone))
}

// TimezoneEqualFold applies the EqualFold predicate on the "timezone" field.
func TimezoneEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldTimezone, v))
}

// TimezoneContainsFold applies the ContainsFold predicate on the "timezone" field.
func TimezoneContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldTimezone, v))
}

// DstEQ applies the EQ predicate on the "dst" field.
func DstEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldDst, v))
}

// DstNEQ applies the NEQ predicate on the "dst" field.
func DstNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldDst, v))
}

// DstIn applies the In predicate on the "dst" field.
func DstIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldDst, vs...))
}

// DstNotIn applies the NotIn predicate on the "dst" field.
func DstNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldDst, vs...))
}

// DstGT applies the GT predicate on the "dst" field.
func DstGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldDst, v))
}

// DstGTE applies the GTE predicate on the "dst" field.
func DstGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldDst, v))
}

// DstLT applies the LT predicate on the "dst" field.
func DstLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldDst, v))
}

// DstLTE applies the LTE predicate on the "dst" field.
func DstLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldDst, v))
}

// DstContains applies the Contains predicate on the "dst" field.
func DstContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldDst, v))
}

// DstHasPrefix applies the HasPrefix predicate on the "dst" field.
func DstHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldDst, v))
}

// DstHasSuffix applies the HasSuffix predicate on the "dst" field.
func DstHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldDst, v))
}

// DstIsNil applies the IsNil predicate on the "dst" field.
func DstIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldDst))
}

// DstNotNil applies the NotNil predicate on the "dst" field.
func DstNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldDst))
}

// DstEqualFold applies the EqualFold predicate on the "dst" field.
func DstEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldDst, v))
}

// DstContainsFold applies the ContainsFold predicate on the "dst" field.
func DstContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldDst, v))
}

// TimezoneNameEQ applies the EQ predicate on the "timezoneName" field.
func TimezoneNameEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldTimezoneName, v))
}

// TimezoneNameNEQ applies the NEQ predicate on the "timezoneName" field.
func TimezoneNameNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldTimezoneName, v))
}

// TimezoneNameIn applies the In predicate on the "timezoneName" field.
func TimezoneNameIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldTimezoneName, vs...))
}

// TimezoneNameNotIn applies the NotIn predicate on the "timezoneName" field.
func TimezoneNameNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldTimezoneName, vs...))
}

// TimezoneNameGT applies the GT predicate on the "timezoneName" field.
func TimezoneNameGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldTimezoneName, v))
}

// TimezoneNameGTE applies the GTE predicate on the "timezoneName" field.
func TimezoneNameGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldTimezoneName, v))
}

// TimezoneNameLT applies the LT predicate on the "timezoneName" field.
func TimezoneNameLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldTimezoneName, v))
}

// TimezoneNameLTE applies the LTE predicate on the "timezoneName" field.
func TimezoneNameLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldTimezoneName, v))
}

// TimezoneNameContains applies the Contains predicate on the "timezoneName" field.
func TimezoneNameContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldTimezoneName, v))
}

// TimezoneNameHasPrefix applies the HasPrefix predicate on the "timezoneName" field.
func TimezoneNameHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldTimezoneName, v))
}

// TimezoneNameHasSuffix applies the HasSuffix predicate on the "timezoneName" field.
func TimezoneNameHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldTimezoneName, v))
}

// TimezoneNameIsNil applies the IsNil predicate on the "timezoneName" field.
func TimezoneNameIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldTimezoneName))
}

// TimezoneNameNotNil applies the NotNil predicate on the "timezoneName" field.
func TimezoneNameNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldTimezoneName))
}

// TimezoneNameEqualFold applies the EqualFold predicate on the "timezoneName" field.
func TimezoneNameEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldTimezoneName, v))
}

// TimezoneNameContainsFold applies the ContainsFold predicate on the "timezoneName" field.
func TimezoneNameContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldTimezoneName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldType, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldSource, v))
}

// SourceIsNil applies the IsNil predicate on the "source" field.
func SourceIsNil() predicate.Airport {
	return predicate.Airport(sql.FieldIsNull(FieldSource))
}

// SourceNotNil applies the NotNil predicate on the "source" field.
func SourceNotNil() predicate.Airport {
	return predicate.Airport(sql.FieldNotNull(FieldSource))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldSource, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Airport) predicate.Airport {
	return predicate.Airport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Airport) predicate.Airport {
	return predicate.Airport(func(s *sql.Selector) {
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
func Not(p predicate.Airport) predicate.Airport {
	return predicate.Airport(func(s *sql.Selector) {
		p(s.Not())
	})
}
