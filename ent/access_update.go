// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/mission-control/ent/access"
	"github.com/orbit-ops/mission-control/ent/predicate"
)

// AccessUpdate is the builder for updating Access entities.
type AccessUpdate struct {
	config
	hooks    []Hook
	mutation *AccessMutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (au *AccessUpdate) Where(ps ...predicate.Access) *AccessUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetRolledBack sets the "rolled_back" field.
func (au *AccessUpdate) SetRolledBack(b bool) *AccessUpdate {
	au.mutation.SetRolledBack(b)
	return au
}

// SetNillableRolledBack sets the "rolled_back" field if the given value is not nil.
func (au *AccessUpdate) SetNillableRolledBack(b *bool) *AccessUpdate {
	if b != nil {
		au.SetRolledBack(*b)
	}
	return au
}

// SetRollbackTime sets the "rollback_time" field.
func (au *AccessUpdate) SetRollbackTime(t time.Time) *AccessUpdate {
	au.mutation.SetRollbackTime(t)
	return au
}

// SetNillableRollbackTime sets the "rollback_time" field if the given value is not nil.
func (au *AccessUpdate) SetNillableRollbackTime(t *time.Time) *AccessUpdate {
	if t != nil {
		au.SetRollbackTime(*t)
	}
	return au
}

// ClearRollbackTime clears the value of the "rollback_time" field.
func (au *AccessUpdate) ClearRollbackTime() *AccessUpdate {
	au.mutation.ClearRollbackTime()
	return au
}

// SetRequestID sets the "request_id" field.
func (au *AccessUpdate) SetRequestID(u uuid.UUID) *AccessUpdate {
	au.mutation.SetRequestID(u)
	return au
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (au *AccessUpdate) SetNillableRequestID(u *uuid.UUID) *AccessUpdate {
	if u != nil {
		au.SetRequestID(*u)
	}
	return au
}

// SetApprovalsID sets the "approvals" edge to the Access entity by ID.
func (au *AccessUpdate) SetApprovalsID(id uuid.UUID) *AccessUpdate {
	au.mutation.SetApprovalsID(id)
	return au
}

// SetNillableApprovalsID sets the "approvals" edge to the Access entity by ID if the given value is not nil.
func (au *AccessUpdate) SetNillableApprovalsID(id *uuid.UUID) *AccessUpdate {
	if id != nil {
		au = au.SetApprovalsID(*id)
	}
	return au
}

// SetApprovals sets the "approvals" edge to the Access entity.
func (au *AccessUpdate) SetApprovals(a *Access) *AccessUpdate {
	return au.SetApprovalsID(a.ID)
}

// Mutation returns the AccessMutation object of the builder.
func (au *AccessUpdate) Mutation() *AccessMutation {
	return au.mutation
}

// ClearApprovals clears the "approvals" edge to the Access entity.
func (au *AccessUpdate) ClearApprovals() *AccessUpdate {
	au.mutation.ClearApprovals()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccessUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccessUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccessUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccessUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.RolledBack(); ok {
		_spec.SetField(access.FieldRolledBack, field.TypeBool, value)
	}
	if value, ok := au.mutation.RollbackTime(); ok {
		_spec.SetField(access.FieldRollbackTime, field.TypeTime, value)
	}
	if au.mutation.RollbackTimeCleared() {
		_spec.ClearField(access.FieldRollbackTime, field.TypeTime)
	}
	if value, ok := au.mutation.RequestID(); ok {
		_spec.SetField(access.FieldRequestID, field.TypeUUID, value)
	}
	if au.mutation.ApprovalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   access.ApprovalsTable,
			Columns: []string{access.ApprovalsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   access.ApprovalsTable,
			Columns: []string{access.ApprovalsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccessUpdateOne is the builder for updating a single Access entity.
type AccessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessMutation
}

// SetRolledBack sets the "rolled_back" field.
func (auo *AccessUpdateOne) SetRolledBack(b bool) *AccessUpdateOne {
	auo.mutation.SetRolledBack(b)
	return auo
}

// SetNillableRolledBack sets the "rolled_back" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableRolledBack(b *bool) *AccessUpdateOne {
	if b != nil {
		auo.SetRolledBack(*b)
	}
	return auo
}

// SetRollbackTime sets the "rollback_time" field.
func (auo *AccessUpdateOne) SetRollbackTime(t time.Time) *AccessUpdateOne {
	auo.mutation.SetRollbackTime(t)
	return auo
}

// SetNillableRollbackTime sets the "rollback_time" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableRollbackTime(t *time.Time) *AccessUpdateOne {
	if t != nil {
		auo.SetRollbackTime(*t)
	}
	return auo
}

// ClearRollbackTime clears the value of the "rollback_time" field.
func (auo *AccessUpdateOne) ClearRollbackTime() *AccessUpdateOne {
	auo.mutation.ClearRollbackTime()
	return auo
}

// SetRequestID sets the "request_id" field.
func (auo *AccessUpdateOne) SetRequestID(u uuid.UUID) *AccessUpdateOne {
	auo.mutation.SetRequestID(u)
	return auo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableRequestID(u *uuid.UUID) *AccessUpdateOne {
	if u != nil {
		auo.SetRequestID(*u)
	}
	return auo
}

// SetApprovalsID sets the "approvals" edge to the Access entity by ID.
func (auo *AccessUpdateOne) SetApprovalsID(id uuid.UUID) *AccessUpdateOne {
	auo.mutation.SetApprovalsID(id)
	return auo
}

// SetNillableApprovalsID sets the "approvals" edge to the Access entity by ID if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableApprovalsID(id *uuid.UUID) *AccessUpdateOne {
	if id != nil {
		auo = auo.SetApprovalsID(*id)
	}
	return auo
}

// SetApprovals sets the "approvals" edge to the Access entity.
func (auo *AccessUpdateOne) SetApprovals(a *Access) *AccessUpdateOne {
	return auo.SetApprovalsID(a.ID)
}

// Mutation returns the AccessMutation object of the builder.
func (auo *AccessUpdateOne) Mutation() *AccessMutation {
	return auo.mutation
}

// ClearApprovals clears the "approvals" edge to the Access entity.
func (auo *AccessUpdateOne) ClearApprovals() *AccessUpdateOne {
	auo.mutation.ClearApprovals()
	return auo
}

// Where appends a list predicates to the AccessUpdate builder.
func (auo *AccessUpdateOne) Where(ps ...predicate.Access) *AccessUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccessUpdateOne) Select(field string, fields ...string) *AccessUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Access entity.
func (auo *AccessUpdateOne) Save(ctx context.Context) (*Access, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccessUpdateOne) SaveX(ctx context.Context) *Access {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccessUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccessUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccessUpdateOne) sqlSave(ctx context.Context) (_node *Access, err error) {
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Access.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, access.FieldID)
		for _, f := range fields {
			if !access.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != access.FieldID {
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
	if value, ok := auo.mutation.RolledBack(); ok {
		_spec.SetField(access.FieldRolledBack, field.TypeBool, value)
	}
	if value, ok := auo.mutation.RollbackTime(); ok {
		_spec.SetField(access.FieldRollbackTime, field.TypeTime, value)
	}
	if auo.mutation.RollbackTimeCleared() {
		_spec.ClearField(access.FieldRollbackTime, field.TypeTime)
	}
	if value, ok := auo.mutation.RequestID(); ok {
		_spec.SetField(access.FieldRequestID, field.TypeUUID, value)
	}
	if auo.mutation.ApprovalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   access.ApprovalsTable,
			Columns: []string{access.ApprovalsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   access.ApprovalsTable,
			Columns: []string{access.ApprovalsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(access.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Access{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
