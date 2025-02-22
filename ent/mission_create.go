// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/mission"
	"github.com/orbit-ops/launchpad-core/ent/request"
	"github.com/orbit-ops/launchpad-core/ent/rocket"
)

// MissionCreate is the builder for creating a Mission entity.
type MissionCreate struct {
	config
	mutation *MissionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (mc *MissionCreate) SetName(s string) *MissionCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetDescription sets the "description" field.
func (mc *MissionCreate) SetDescription(s string) *MissionCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MissionCreate) SetNillableDescription(s *string) *MissionCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetDuration sets the "duration" field.
func (mc *MissionCreate) SetDuration(i int) *MissionCreate {
	mc.mutation.SetDuration(i)
	return mc
}

// SetMinApprovers sets the "min_approvers" field.
func (mc *MissionCreate) SetMinApprovers(i int) *MissionCreate {
	mc.mutation.SetMinApprovers(i)
	return mc
}

// SetPossibleApprovers sets the "possible_approvers" field.
func (mc *MissionCreate) SetPossibleApprovers(s []string) *MissionCreate {
	mc.mutation.SetPossibleApprovers(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MissionCreate) SetID(u uuid.UUID) *MissionCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MissionCreate) SetNillableID(u *uuid.UUID) *MissionCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// AddRocketIDs adds the "rockets" edge to the Rocket entity by IDs.
func (mc *MissionCreate) AddRocketIDs(ids ...uuid.UUID) *MissionCreate {
	mc.mutation.AddRocketIDs(ids...)
	return mc
}

// AddRockets adds the "rockets" edges to the Rocket entity.
func (mc *MissionCreate) AddRockets(r ...*Rocket) *MissionCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRocketIDs(ids...)
}

// AddRequestIDs adds the "requests" edge to the Request entity by IDs.
func (mc *MissionCreate) AddRequestIDs(ids ...uuid.UUID) *MissionCreate {
	mc.mutation.AddRequestIDs(ids...)
	return mc
}

// AddRequests adds the "requests" edges to the Request entity.
func (mc *MissionCreate) AddRequests(r ...*Request) *MissionCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRequestIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (mc *MissionCreate) Mutation() *MissionMutation {
	return mc.mutation
}

// Save creates the Mission in the database.
func (mc *MissionCreate) Save(ctx context.Context) (*Mission, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MissionCreate) SaveX(ctx context.Context) *Mission {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MissionCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MissionCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MissionCreate) defaults() {
	if _, ok := mc.mutation.ID(); !ok {
		v := mission.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MissionCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Mission.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := mission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Mission.name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Mission.duration"`)}
	}
	if v, ok := mc.mutation.Duration(); ok {
		if err := mission.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Mission.duration": %w`, err)}
		}
	}
	if _, ok := mc.mutation.MinApprovers(); !ok {
		return &ValidationError{Name: "min_approvers", err: errors.New(`ent: missing required field "Mission.min_approvers"`)}
	}
	if v, ok := mc.mutation.MinApprovers(); ok {
		if err := mission.MinApproversValidator(v); err != nil {
			return &ValidationError{Name: "min_approvers", err: fmt.Errorf(`ent: validator failed for field "Mission.min_approvers": %w`, err)}
		}
	}
	if _, ok := mc.mutation.PossibleApprovers(); !ok {
		return &ValidationError{Name: "possible_approvers", err: errors.New(`ent: missing required field "Mission.possible_approvers"`)}
	}
	if len(mc.mutation.RocketsIDs()) == 0 {
		return &ValidationError{Name: "rockets", err: errors.New(`ent: missing required edge "Mission.rockets"`)}
	}
	return nil
}

func (mc *MissionCreate) sqlSave(ctx context.Context) (*Mission, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MissionCreate) createSpec() (*Mission, *sqlgraph.CreateSpec) {
	var (
		_node = &Mission{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(mission.Table, sqlgraph.NewFieldSpec(mission.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(mission.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(mission.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mc.mutation.Duration(); ok {
		_spec.SetField(mission.FieldDuration, field.TypeInt, value)
		_node.Duration = value
	}
	if value, ok := mc.mutation.MinApprovers(); ok {
		_spec.SetField(mission.FieldMinApprovers, field.TypeInt, value)
		_node.MinApprovers = value
	}
	if value, ok := mc.mutation.PossibleApprovers(); ok {
		_spec.SetField(mission.FieldPossibleApprovers, field.TypeJSON, value)
		_node.PossibleApprovers = value
	}
	if nodes := mc.mutation.RocketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.RocketsTable,
			Columns: []string{mission.RocketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rocket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mission.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MissionUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (mc *MissionCreate) OnConflict(opts ...sql.ConflictOption) *MissionUpsertOne {
	mc.conflict = opts
	return &MissionUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MissionCreate) OnConflictColumns(columns ...string) *MissionUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MissionUpsertOne{
		create: mc,
	}
}

type (
	// MissionUpsertOne is the builder for "upsert"-ing
	//  one Mission node.
	MissionUpsertOne struct {
		create *MissionCreate
	}

	// MissionUpsert is the "OnConflict" setter.
	MissionUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *MissionUpsert) SetName(v string) *MissionUpsert {
	u.Set(mission.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MissionUpsert) UpdateName() *MissionUpsert {
	u.SetExcluded(mission.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *MissionUpsert) SetDescription(v string) *MissionUpsert {
	u.Set(mission.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *MissionUpsert) UpdateDescription() *MissionUpsert {
	u.SetExcluded(mission.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *MissionUpsert) ClearDescription() *MissionUpsert {
	u.SetNull(mission.FieldDescription)
	return u
}

// SetDuration sets the "duration" field.
func (u *MissionUpsert) SetDuration(v int) *MissionUpsert {
	u.Set(mission.FieldDuration, v)
	return u
}

// UpdateDuration sets the "duration" field to the value that was provided on create.
func (u *MissionUpsert) UpdateDuration() *MissionUpsert {
	u.SetExcluded(mission.FieldDuration)
	return u
}

// AddDuration adds v to the "duration" field.
func (u *MissionUpsert) AddDuration(v int) *MissionUpsert {
	u.Add(mission.FieldDuration, v)
	return u
}

// SetMinApprovers sets the "min_approvers" field.
func (u *MissionUpsert) SetMinApprovers(v int) *MissionUpsert {
	u.Set(mission.FieldMinApprovers, v)
	return u
}

// UpdateMinApprovers sets the "min_approvers" field to the value that was provided on create.
func (u *MissionUpsert) UpdateMinApprovers() *MissionUpsert {
	u.SetExcluded(mission.FieldMinApprovers)
	return u
}

// AddMinApprovers adds v to the "min_approvers" field.
func (u *MissionUpsert) AddMinApprovers(v int) *MissionUpsert {
	u.Add(mission.FieldMinApprovers, v)
	return u
}

// SetPossibleApprovers sets the "possible_approvers" field.
func (u *MissionUpsert) SetPossibleApprovers(v []string) *MissionUpsert {
	u.Set(mission.FieldPossibleApprovers, v)
	return u
}

// UpdatePossibleApprovers sets the "possible_approvers" field to the value that was provided on create.
func (u *MissionUpsert) UpdatePossibleApprovers() *MissionUpsert {
	u.SetExcluded(mission.FieldPossibleApprovers)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Mission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MissionUpsertOne) UpdateNewValues() *MissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mission.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Mission.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MissionUpsertOne) Ignore() *MissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MissionUpsertOne) DoNothing() *MissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MissionCreate.OnConflict
// documentation for more info.
func (u *MissionUpsertOne) Update(set func(*MissionUpsert)) *MissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MissionUpsertOne) SetName(v string) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MissionUpsertOne) UpdateName() *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *MissionUpsertOne) SetDescription(v string) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *MissionUpsertOne) UpdateDescription() *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *MissionUpsertOne) ClearDescription() *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.ClearDescription()
	})
}

// SetDuration sets the "duration" field.
func (u *MissionUpsertOne) SetDuration(v int) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.SetDuration(v)
	})
}

// AddDuration adds v to the "duration" field.
func (u *MissionUpsertOne) AddDuration(v int) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.AddDuration(v)
	})
}

// UpdateDuration sets the "duration" field to the value that was provided on create.
func (u *MissionUpsertOne) UpdateDuration() *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateDuration()
	})
}

// SetMinApprovers sets the "min_approvers" field.
func (u *MissionUpsertOne) SetMinApprovers(v int) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.SetMinApprovers(v)
	})
}

// AddMinApprovers adds v to the "min_approvers" field.
func (u *MissionUpsertOne) AddMinApprovers(v int) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.AddMinApprovers(v)
	})
}

// UpdateMinApprovers sets the "min_approvers" field to the value that was provided on create.
func (u *MissionUpsertOne) UpdateMinApprovers() *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateMinApprovers()
	})
}

// SetPossibleApprovers sets the "possible_approvers" field.
func (u *MissionUpsertOne) SetPossibleApprovers(v []string) *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.SetPossibleApprovers(v)
	})
}

// UpdatePossibleApprovers sets the "possible_approvers" field to the value that was provided on create.
func (u *MissionUpsertOne) UpdatePossibleApprovers() *MissionUpsertOne {
	return u.Update(func(s *MissionUpsert) {
		s.UpdatePossibleApprovers()
	})
}

// Exec executes the query.
func (u *MissionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MissionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MissionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MissionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MissionUpsertOne.ID is not supported by MySQL driver. Use MissionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MissionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MissionCreateBulk is the builder for creating many Mission entities in bulk.
type MissionCreateBulk struct {
	config
	err      error
	builders []*MissionCreate
	conflict []sql.ConflictOption
}

// Save creates the Mission entities in the database.
func (mcb *MissionCreateBulk) Save(ctx context.Context) ([]*Mission, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mission, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MissionCreateBulk) SaveX(ctx context.Context) []*Mission {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MissionCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MissionCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mission.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MissionUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (mcb *MissionCreateBulk) OnConflict(opts ...sql.ConflictOption) *MissionUpsertBulk {
	mcb.conflict = opts
	return &MissionUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MissionCreateBulk) OnConflictColumns(columns ...string) *MissionUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MissionUpsertBulk{
		create: mcb,
	}
}

// MissionUpsertBulk is the builder for "upsert"-ing
// a bulk of Mission nodes.
type MissionUpsertBulk struct {
	create *MissionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Mission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MissionUpsertBulk) UpdateNewValues() *MissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mission.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Mission.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MissionUpsertBulk) Ignore() *MissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MissionUpsertBulk) DoNothing() *MissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MissionCreateBulk.OnConflict
// documentation for more info.
func (u *MissionUpsertBulk) Update(set func(*MissionUpsert)) *MissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MissionUpsertBulk) SetName(v string) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MissionUpsertBulk) UpdateName() *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *MissionUpsertBulk) SetDescription(v string) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *MissionUpsertBulk) UpdateDescription() *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *MissionUpsertBulk) ClearDescription() *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.ClearDescription()
	})
}

// SetDuration sets the "duration" field.
func (u *MissionUpsertBulk) SetDuration(v int) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.SetDuration(v)
	})
}

// AddDuration adds v to the "duration" field.
func (u *MissionUpsertBulk) AddDuration(v int) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.AddDuration(v)
	})
}

// UpdateDuration sets the "duration" field to the value that was provided on create.
func (u *MissionUpsertBulk) UpdateDuration() *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateDuration()
	})
}

// SetMinApprovers sets the "min_approvers" field.
func (u *MissionUpsertBulk) SetMinApprovers(v int) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.SetMinApprovers(v)
	})
}

// AddMinApprovers adds v to the "min_approvers" field.
func (u *MissionUpsertBulk) AddMinApprovers(v int) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.AddMinApprovers(v)
	})
}

// UpdateMinApprovers sets the "min_approvers" field to the value that was provided on create.
func (u *MissionUpsertBulk) UpdateMinApprovers() *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.UpdateMinApprovers()
	})
}

// SetPossibleApprovers sets the "possible_approvers" field.
func (u *MissionUpsertBulk) SetPossibleApprovers(v []string) *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.SetPossibleApprovers(v)
	})
}

// UpdatePossibleApprovers sets the "possible_approvers" field to the value that was provided on create.
func (u *MissionUpsertBulk) UpdatePossibleApprovers() *MissionUpsertBulk {
	return u.Update(func(s *MissionUpsert) {
		s.UpdatePossibleApprovers()
	})
}

// Exec executes the query.
func (u *MissionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MissionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MissionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MissionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
