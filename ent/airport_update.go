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