// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/site-tech/fake-data-service/ent/plane"
	"github.com/site-tech/fake-data-service/ent/predicate"
)

// PlaneUpdate is the builder for updating Plane entities.
type PlaneUpdate struct {
	config
	hooks    []Hook
	mutation *PlaneMutation
}

// Where appends a list predicates to the PlaneUpdate builder.
func (pu *PlaneUpdate) Where(ps ...predicate.Plane) *PlaneUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlaneUpdate) SetName(s string) *PlaneUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetTailNumber sets the "tailNumber" field.
func (pu *PlaneUpdate) SetTailNumber(s string) *PlaneUpdate {
	pu.mutation.SetTailNumber(s)
	return pu
}

// Mutation returns the PlaneMutation object of the builder.
func (pu *PlaneUpdate) Mutation() *PlaneMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlaneUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PlaneMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlaneUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlaneUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlaneUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlaneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(plane.Table, plane.Columns, sqlgraph.NewFieldSpec(plane.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(plane.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.TailNumber(); ok {
		_spec.SetField(plane.FieldTailNumber, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plane.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlaneUpdateOne is the builder for updating a single Plane entity.
type PlaneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlaneMutation
}

// SetName sets the "name" field.
func (puo *PlaneUpdateOne) SetName(s string) *PlaneUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetTailNumber sets the "tailNumber" field.
func (puo *PlaneUpdateOne) SetTailNumber(s string) *PlaneUpdateOne {
	puo.mutation.SetTailNumber(s)
	return puo
}

// Mutation returns the PlaneMutation object of the builder.
func (puo *PlaneUpdateOne) Mutation() *PlaneMutation {
	return puo.mutation
}

// Where appends a list predicates to the PlaneUpdate builder.
func (puo *PlaneUpdateOne) Where(ps ...predicate.Plane) *PlaneUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlaneUpdateOne) Select(field string, fields ...string) *PlaneUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Plane entity.
func (puo *PlaneUpdateOne) Save(ctx context.Context) (*Plane, error) {
	return withHooks[*Plane, PlaneMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlaneUpdateOne) SaveX(ctx context.Context) *Plane {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlaneUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlaneUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlaneUpdateOne) sqlSave(ctx context.Context) (_node *Plane, err error) {
	_spec := sqlgraph.NewUpdateSpec(plane.Table, plane.Columns, sqlgraph.NewFieldSpec(plane.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Plane.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plane.FieldID)
		for _, f := range fields {
			if !plane.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != plane.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(plane.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.TailNumber(); ok {
		_spec.SetField(plane.FieldTailNumber, field.TypeString, value)
	}
	_node = &Plane{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plane.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
