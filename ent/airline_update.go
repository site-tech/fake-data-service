// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/site-tech/fake-data-service/ent/airline"
	"github.com/site-tech/fake-data-service/ent/predicate"
)

// AirlineUpdate is the builder for updating Airline entities.
type AirlineUpdate struct {
	config
	hooks    []Hook
	mutation *AirlineMutation
}

// Where appends a list predicates to the AirlineUpdate builder.
func (au *AirlineUpdate) Where(ps ...predicate.Airline) *AirlineUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AirlineUpdate) SetName(s string) *AirlineUpdate {
	au.mutation.SetName(s)
	return au
}

// SetAlias sets the "alias" field.
func (au *AirlineUpdate) SetAlias(s string) *AirlineUpdate {
	au.mutation.SetAlias(s)
	return au
}

// SetCountry sets the "country" field.
func (au *AirlineUpdate) SetCountry(s string) *AirlineUpdate {
	au.mutation.SetCountry(s)
	return au
}

// SetActive sets the "active" field.
func (au *AirlineUpdate) SetActive(b bool) *AirlineUpdate {
	au.mutation.SetActive(b)
	return au
}

// Mutation returns the AirlineMutation object of the builder.
func (au *AirlineUpdate) Mutation() *AirlineMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AirlineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AirlineMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AirlineUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AirlineUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AirlineUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AirlineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(airline.Table, airline.Columns, sqlgraph.NewFieldSpec(airline.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(airline.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Alias(); ok {
		_spec.SetField(airline.FieldAlias, field.TypeString, value)
	}
	if value, ok := au.mutation.Country(); ok {
		_spec.SetField(airline.FieldCountry, field.TypeString, value)
	}
	if value, ok := au.mutation.Active(); ok {
		_spec.SetField(airline.FieldActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AirlineUpdateOne is the builder for updating a single Airline entity.
type AirlineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AirlineMutation
}

// SetName sets the "name" field.
func (auo *AirlineUpdateOne) SetName(s string) *AirlineUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetAlias sets the "alias" field.
func (auo *AirlineUpdateOne) SetAlias(s string) *AirlineUpdateOne {
	auo.mutation.SetAlias(s)
	return auo
}

// SetCountry sets the "country" field.
func (auo *AirlineUpdateOne) SetCountry(s string) *AirlineUpdateOne {
	auo.mutation.SetCountry(s)
	return auo
}

// SetActive sets the "active" field.
func (auo *AirlineUpdateOne) SetActive(b bool) *AirlineUpdateOne {
	auo.mutation.SetActive(b)
	return auo
}

// Mutation returns the AirlineMutation object of the builder.
func (auo *AirlineUpdateOne) Mutation() *AirlineMutation {
	return auo.mutation
}

// Where appends a list predicates to the AirlineUpdate builder.
func (auo *AirlineUpdateOne) Where(ps ...predicate.Airline) *AirlineUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AirlineUpdateOne) Select(field string, fields ...string) *AirlineUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Airline entity.
func (auo *AirlineUpdateOne) Save(ctx context.Context) (*Airline, error) {
	return withHooks[*Airline, AirlineMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AirlineUpdateOne) SaveX(ctx context.Context) *Airline {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AirlineUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AirlineUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AirlineUpdateOne) sqlSave(ctx context.Context) (_node *Airline, err error) {
	_spec := sqlgraph.NewUpdateSpec(airline.Table, airline.Columns, sqlgraph.NewFieldSpec(airline.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Airline.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airline.FieldID)
		for _, f := range fields {
			if !airline.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != airline.FieldID {
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
		_spec.SetField(airline.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Alias(); ok {
		_spec.SetField(airline.FieldAlias, field.TypeString, value)
	}
	if value, ok := auo.mutation.Country(); ok {
		_spec.SetField(airline.FieldCountry, field.TypeString, value)
	}
	if value, ok := auo.mutation.Active(); ok {
		_spec.SetField(airline.FieldActive, field.TypeBool, value)
	}
	_node = &Airline{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}