// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/mission-control/ent/mission"
	"github.com/orbit-ops/mission-control/ent/predicate"
	"github.com/orbit-ops/mission-control/ent/request"
)

// MissionUpdate is the builder for updating Mission entities.
type MissionUpdate struct {
	config
	hooks    []Hook
	mutation *MissionMutation
}

// Where appends a list predicates to the MissionUpdate builder.
func (mu *MissionUpdate) Where(ps ...predicate.Mission) *MissionUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetDescription sets the "description" field.
func (mu *MissionUpdate) SetDescription(s string) *MissionUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableDescription(s *string) *MissionUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// ClearDescription clears the value of the "description" field.
func (mu *MissionUpdate) ClearDescription() *MissionUpdate {
	mu.mutation.ClearDescription()
	return mu
}

// SetImage sets the "image" field.
func (mu *MissionUpdate) SetImage(s string) *MissionUpdate {
	mu.mutation.SetImage(s)
	return mu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableImage(s *string) *MissionUpdate {
	if s != nil {
		mu.SetImage(*s)
	}
	return mu
}

// SetMinApprovers sets the "min_approvers" field.
func (mu *MissionUpdate) SetMinApprovers(i int) *MissionUpdate {
	mu.mutation.ResetMinApprovers()
	mu.mutation.SetMinApprovers(i)
	return mu
}

// SetNillableMinApprovers sets the "min_approvers" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableMinApprovers(i *int) *MissionUpdate {
	if i != nil {
		mu.SetMinApprovers(*i)
	}
	return mu
}

// AddMinApprovers adds i to the "min_approvers" field.
func (mu *MissionUpdate) AddMinApprovers(i int) *MissionUpdate {
	mu.mutation.AddMinApprovers(i)
	return mu
}

// AddRequestIDs adds the "requests" edge to the Request entity by IDs.
func (mu *MissionUpdate) AddRequestIDs(ids ...uuid.UUID) *MissionUpdate {
	mu.mutation.AddRequestIDs(ids...)
	return mu
}

// AddRequests adds the "requests" edges to the Request entity.
func (mu *MissionUpdate) AddRequests(r ...*Request) *MissionUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.AddRequestIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (mu *MissionUpdate) Mutation() *MissionMutation {
	return mu.mutation
}

// ClearRequests clears all "requests" edges to the Request entity.
func (mu *MissionUpdate) ClearRequests() *MissionUpdate {
	mu.mutation.ClearRequests()
	return mu
}

// RemoveRequestIDs removes the "requests" edge to Request entities by IDs.
func (mu *MissionUpdate) RemoveRequestIDs(ids ...uuid.UUID) *MissionUpdate {
	mu.mutation.RemoveRequestIDs(ids...)
	return mu
}

// RemoveRequests removes "requests" edges to Request entities.
func (mu *MissionUpdate) RemoveRequests(r ...*Request) *MissionUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.RemoveRequestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MissionUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MissionUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MissionUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MissionUpdate) check() error {
	if _, ok := mu.mutation.RocketID(); mu.mutation.RocketCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mission.rocket"`)
	}
	return nil
}

func (mu *MissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(mission.Table, mission.Columns, sqlgraph.NewFieldSpec(mission.FieldID, field.TypeString))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(mission.FieldDescription, field.TypeString, value)
	}
	if mu.mutation.DescriptionCleared() {
		_spec.ClearField(mission.FieldDescription, field.TypeString)
	}
	if value, ok := mu.mutation.Image(); ok {
		_spec.SetField(mission.FieldImage, field.TypeString, value)
	}
	if value, ok := mu.mutation.MinApprovers(); ok {
		_spec.SetField(mission.FieldMinApprovers, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedMinApprovers(); ok {
		_spec.AddField(mission.FieldMinApprovers, field.TypeInt, value)
	}
	if mu.mutation.RequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RequestsTable,
			Columns: []string{mission.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedRequestsIDs(); len(nodes) > 0 && !mu.mutation.RequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RequestsTable,
			Columns: []string{mission.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RequestsTable,
			Columns: []string{mission.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MissionUpdateOne is the builder for updating a single Mission entity.
type MissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MissionMutation
}

// SetDescription sets the "description" field.
func (muo *MissionUpdateOne) SetDescription(s string) *MissionUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableDescription(s *string) *MissionUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// ClearDescription clears the value of the "description" field.
func (muo *MissionUpdateOne) ClearDescription() *MissionUpdateOne {
	muo.mutation.ClearDescription()
	return muo
}

// SetImage sets the "image" field.
func (muo *MissionUpdateOne) SetImage(s string) *MissionUpdateOne {
	muo.mutation.SetImage(s)
	return muo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableImage(s *string) *MissionUpdateOne {
	if s != nil {
		muo.SetImage(*s)
	}
	return muo
}

// SetMinApprovers sets the "min_approvers" field.
func (muo *MissionUpdateOne) SetMinApprovers(i int) *MissionUpdateOne {
	muo.mutation.ResetMinApprovers()
	muo.mutation.SetMinApprovers(i)
	return muo
}

// SetNillableMinApprovers sets the "min_approvers" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableMinApprovers(i *int) *MissionUpdateOne {
	if i != nil {
		muo.SetMinApprovers(*i)
	}
	return muo
}

// AddMinApprovers adds i to the "min_approvers" field.
func (muo *MissionUpdateOne) AddMinApprovers(i int) *MissionUpdateOne {
	muo.mutation.AddMinApprovers(i)
	return muo
}

// AddRequestIDs adds the "requests" edge to the Request entity by IDs.
func (muo *MissionUpdateOne) AddRequestIDs(ids ...uuid.UUID) *MissionUpdateOne {
	muo.mutation.AddRequestIDs(ids...)
	return muo
}

// AddRequests adds the "requests" edges to the Request entity.
func (muo *MissionUpdateOne) AddRequests(r ...*Request) *MissionUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.AddRequestIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (muo *MissionUpdateOne) Mutation() *MissionMutation {
	return muo.mutation
}

// ClearRequests clears all "requests" edges to the Request entity.
func (muo *MissionUpdateOne) ClearRequests() *MissionUpdateOne {
	muo.mutation.ClearRequests()
	return muo
}

// RemoveRequestIDs removes the "requests" edge to Request entities by IDs.
func (muo *MissionUpdateOne) RemoveRequestIDs(ids ...uuid.UUID) *MissionUpdateOne {
	muo.mutation.RemoveRequestIDs(ids...)
	return muo
}

// RemoveRequests removes "requests" edges to Request entities.
func (muo *MissionUpdateOne) RemoveRequests(r ...*Request) *MissionUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.RemoveRequestIDs(ids...)
}

// Where appends a list predicates to the MissionUpdate builder.
func (muo *MissionUpdateOne) Where(ps ...predicate.Mission) *MissionUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MissionUpdateOne) Select(field string, fields ...string) *MissionUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mission entity.
func (muo *MissionUpdateOne) Save(ctx context.Context) (*Mission, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MissionUpdateOne) SaveX(ctx context.Context) *Mission {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MissionUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MissionUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MissionUpdateOne) check() error {
	if _, ok := muo.mutation.RocketID(); muo.mutation.RocketCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mission.rocket"`)
	}
	return nil
}

func (muo *MissionUpdateOne) sqlSave(ctx context.Context) (_node *Mission, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(mission.Table, mission.Columns, sqlgraph.NewFieldSpec(mission.FieldID, field.TypeString))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mission.FieldID)
		for _, f := range fields {
			if !mission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(mission.FieldDescription, field.TypeString, value)
	}
	if muo.mutation.DescriptionCleared() {
		_spec.ClearField(mission.FieldDescription, field.TypeString)
	}
	if value, ok := muo.mutation.Image(); ok {
		_spec.SetField(mission.FieldImage, field.TypeString, value)
	}
	if value, ok := muo.mutation.MinApprovers(); ok {
		_spec.SetField(mission.FieldMinApprovers, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedMinApprovers(); ok {
		_spec.AddField(mission.FieldMinApprovers, field.TypeInt, value)
	}
	if muo.mutation.RequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RequestsTable,
			Columns: []string{mission.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedRequestsIDs(); len(nodes) > 0 && !muo.mutation.RequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RequestsTable,
			Columns: []string{mission.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RequestsTable,
			Columns: []string{mission.RequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Mission{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
