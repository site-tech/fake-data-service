// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/site-tech/fake-data-service/ent/airport"
	"github.com/site-tech/fake-data-service/ent/predicate"
)

// AirportUpdate is the builder for updating Airport entities.
type AirportUpdate struct {
	config
	hooks    []Hook
	mutation *AirportMutation
}

// Where appends a list predicates to the AirportUpdate builder.
func (au *AirportUpdate) Where(ps ...predicate.Airport) *AirportUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AirportUpdate) SetName(s string) *AirportUpdate {
	au.mutation.SetName(s)
	return au
}

// SetCity sets the "city" field.
func (au *AirportUpdate) SetCity(s string) *AirportUpdate {
	au.mutation.SetCity(s)
	return au
}

// SetCountry sets the "country" field.
func (au *AirportUpdate) SetCountry(s string) *AirportUpdate {
	au.mutation.SetCountry(s)
	return au
}

// SetIata sets the "iata" field.
func (au *AirportUpdate) SetIata(s string) *AirportUpdate {
	au.mutation.SetIata(s)
	return au
}

// SetIcao sets the "icao" field.
func (au *AirportUpdate) SetIcao(s string) *AirportUpdate {
	au.mutation.SetIcao(s)
	return au
}

// SetLatitude sets the "latitude" field.
func (au *AirportUpdate) SetLatitude(f float64) *AirportUpdate {
	au.mutation.ResetLatitude()
	au.mutation.SetLatitude(f)
	return au
}

// AddLatitude adds f to the "latitude" field.
func (au *AirportUpdate) AddLatitude(f float64) *AirportUpdate {
	au.mutation.AddLatitude(f)
	return au
}

// SetLongitude sets the "longitude" field.
func (au *AirportUpdate) SetLongitude(f float64) *AirportUpdate {
	au.mutation.ResetLongitude()
	au.mutation.SetLongitude(f)
	return au
}

// AddLongitude adds f to the "longitude" field.
func (au *AirportUpdate) AddLongitude(f float64) *AirportUpdate {
	au.mutation.AddLongitude(f)
	return au
}

// SetAltitude sets the "altitude" field.
func (au *AirportUpdate) SetAltitude(f float64) *AirportUpdate {
	au.mutation.ResetAltitude()
	au.mutation.SetAltitude(f)
	return au
}

// AddAltitude adds f to the "altitude" field.
func (au *AirportUpdate) AddAltitude(f float64) *AirportUpdate {
	au.mutation.AddAltitude(f)
	return au
}

// SetTimezone sets the "timezone" field.
func (au *AirportUpdate) SetTimezone(f float64) *AirportUpdate {
	au.mutation.ResetTimezone()
	au.mutation.SetTimezone(f)
	return au
}

// AddTimezone adds f to the "timezone" field.
func (au *AirportUpdate) AddTimezone(f float64) *AirportUpdate {
	au.mutation.AddTimezone(f)
	return au
}

// SetDst sets the "dst" field.
func (au *AirportUpdate) SetDst(s string) *AirportUpdate {
	au.mutation.SetDst(s)
	return au
}

// SetTimezoneName sets the "timezoneName" field.
func (au *AirportUpdate) SetTimezoneName(s string) *AirportUpdate {
	au.mutation.SetTimezoneName(s)
	return au
}

// SetType sets the "type" field.
func (au *AirportUpdate) SetType(s string) *AirportUpdate {
	au.mutation.SetType(s)
	return au
}

// SetSource sets the "source" field.
func (au *AirportUpdate) SetSource(s string) *AirportUpdate {
	au.mutation.SetSource(s)
	return au
}

// Mutation returns the AirportMutation object of the builder.
func (au *AirportUpdate) Mutation() *AirportMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AirportUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AirportMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AirportUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AirportUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AirportUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AirportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.City(); ok {
		_spec.SetField(airport.FieldCity, field.TypeString, value)
	}
	if value, ok := au.mutation.Country(); ok {
		_spec.SetField(airport.FieldCountry, field.TypeString, value)
	}
	if value, ok := au.mutation.Iata(); ok {
		_spec.SetField(airport.FieldIata, field.TypeString, value)
	}
	if value, ok := au.mutation.Icao(); ok {
		_spec.SetField(airport.FieldIcao, field.TypeString, value)
	}
	if value, ok := au.mutation.Latitude(); ok {
		_spec.SetField(airport.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedLatitude(); ok {
		_spec.AddField(airport.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.Longitude(); ok {
		_spec.SetField(airport.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedLongitude(); ok {
		_spec.AddField(airport.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.Altitude(); ok {
		_spec.SetField(airport.FieldAltitude, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedAltitude(); ok {
		_spec.AddField(airport.FieldAltitude, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.Timezone(); ok {
		_spec.SetField(airport.FieldTimezone, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedTimezone(); ok {
		_spec.AddField(airport.FieldTimezone, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.Dst(); ok {
		_spec.SetField(airport.FieldDst, field.TypeString, value)
	}
	if value, ok := au.mutation.TimezoneName(); ok {
		_spec.SetField(airport.FieldTimezoneName, field.TypeString, value)
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.SetField(airport.FieldType, field.TypeString, value)
	}
	if value, ok := au.mutation.Source(); ok {
		_spec.SetField(airport.FieldSource, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AirportUpdateOne is the builder for updating a single Airport entity.
type AirportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AirportMutation
}

// SetName sets the "name" field.
func (auo *AirportUpdateOne) SetName(s string) *AirportUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetCity sets the "city" field.
func (auo *AirportUpdateOne) SetCity(s string) *AirportUpdateOne {
	auo.mutation.SetCity(s)
	return auo
}

// SetCountry sets the "country" field.
func (auo *AirportUpdateOne) SetCountry(s string) *AirportUpdateOne {
	auo.mutation.SetCountry(s)
	return auo
}

// SetIata sets the "iata" field.
func (auo *AirportUpdateOne) SetIata(s string) *AirportUpdateOne {
	auo.mutation.SetIata(s)
	return auo
}

// SetIcao sets the "icao" field.
func (auo *AirportUpdateOne) SetIcao(s string) *AirportUpdateOne {
	auo.mutation.SetIcao(s)
	return auo
}

// SetLatitude sets the "latitude" field.
func (auo *AirportUpdateOne) SetLatitude(f float64) *AirportUpdateOne {
	auo.mutation.ResetLatitude()
	auo.mutation.SetLatitude(f)
	return auo
}

// AddLatitude adds f to the "latitude" field.
func (auo *AirportUpdateOne) AddLatitude(f float64) *AirportUpdateOne {
	auo.mutation.AddLatitude(f)
	return auo
}

// SetLongitude sets the "longitude" field.
func (auo *AirportUpdateOne) SetLongitude(f float64) *AirportUpdateOne {
	auo.mutation.ResetLongitude()
	auo.mutation.SetLongitude(f)
	return auo
}

// AddLongitude adds f to the "longitude" field.
func (auo *AirportUpdateOne) AddLongitude(f float64) *AirportUpdateOne {
	auo.mutation.AddLongitude(f)
	return auo
}

// SetAltitude sets the "altitude" field.
func (auo *AirportUpdateOne) SetAltitude(f float64) *AirportUpdateOne {
	auo.mutation.ResetAltitude()
	auo.mutation.SetAltitude(f)
	return auo
}

// AddAltitude adds f to the "altitude" field.
func (auo *AirportUpdateOne) AddAltitude(f float64) *AirportUpdateOne {
	auo.mutation.AddAltitude(f)
	return auo
}

// SetTimezone sets the "timezone" field.
func (auo *AirportUpdateOne) SetTimezone(f float64) *AirportUpdateOne {
	auo.mutation.ResetTimezone()
	auo.mutation.SetTimezone(f)
	return auo
}

// AddTimezone adds f to the "timezone" field.
func (auo *AirportUpdateOne) AddTimezone(f float64) *AirportUpdateOne {
	auo.mutation.AddTimezone(f)
	return auo
}

// SetDst sets the "dst" field.
func (auo *AirportUpdateOne) SetDst(s string) *AirportUpdateOne {
	auo.mutation.SetDst(s)
	return auo
}

// SetTimezoneName sets the "timezoneName" field.
func (auo *AirportUpdateOne) SetTimezoneName(s string) *AirportUpdateOne {
	auo.mutation.SetTimezoneName(s)
	return auo
}

// SetType sets the "type" field.
func (auo *AirportUpdateOne) SetType(s string) *AirportUpdateOne {
	auo.mutation.SetType(s)
	return auo
}

// SetSource sets the "source" field.
func (auo *AirportUpdateOne) SetSource(s string) *AirportUpdateOne {
	auo.mutation.SetSource(s)
	return auo
}

// Mutation returns the AirportMutation object of the builder.
func (auo *AirportUpdateOne) Mutation() *AirportMutation {
	return auo.mutation
}

// Where appends a list predicates to the AirportUpdate builder.
func (auo *AirportUpdateOne) Where(ps ...predicate.Airport) *AirportUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AirportUpdateOne) Select(field string, fields ...string) *AirportUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Airport entity.
func (auo *AirportUpdateOne) Save(ctx context.Context) (*Airport, error) {
	return withHooks[*Airport, AirportMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AirportUpdateOne) SaveX(ctx context.Context) *Airport {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AirportUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AirportUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AirportUpdateOne) sqlSave(ctx context.Context) (_node *Airport, err error) {
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Airport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airport.FieldID)
		for _, f := range fields {
			if !airport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != airport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.City(); ok {
		_spec.SetField(airport.FieldCity, field.TypeString, value)
	}
	if value, ok := auo.mutation.Country(); ok {
		_spec.SetField(airport.FieldCountry, field.TypeString, value)
	}
	if value, ok := auo.mutation.Iata(); ok {
		_spec.SetField(airport.FieldIata, field.TypeString, value)
	}
	if value, ok := auo.mutation.Icao(); ok {
		_spec.SetField(airport.FieldIcao, field.TypeString, value)
	}
	if value, ok := auo.mutation.Latitude(); ok {
		_spec.SetField(airport.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedLatitude(); ok {
		_spec.AddField(airport.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.Longitude(); ok {
		_spec.SetField(airport.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedLongitude(); ok {
		_spec.AddField(airport.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.Altitude(); ok {
		_spec.SetField(airport.FieldAltitude, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedAltitude(); ok {
		_spec.AddField(airport.FieldAltitude, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.Timezone(); ok {
		_spec.SetField(airport.FieldTimezone, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedTimezone(); ok {
		_spec.AddField(airport.FieldTimezone, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.Dst(); ok {
		_spec.SetField(airport.FieldDst, field.TypeString, value)
	}
	if value, ok := auo.mutation.TimezoneName(); ok {
		_spec.SetField(airport.FieldTimezoneName, field.TypeString, value)
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.SetField(airport.FieldType, field.TypeString, value)
	}
	if value, ok := auo.mutation.Source(); ok {
		_spec.SetField(airport.FieldSource, field.TypeString, value)
	}
	_node = &Airport{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
