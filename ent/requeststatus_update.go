// Code generated by entc, DO NOT EDIT.

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
	"github.com/traPtitech/Jomon/ent/predicate"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requeststatus"
	"github.com/traPtitech/Jomon/ent/user"
)

// RequestStatusUpdate is the builder for updating RequestStatus entities.
type RequestStatusUpdate struct {
	config
	hooks    []Hook
	mutation *RequestStatusMutation
}

// Where appends a list predicates to the RequestStatusUpdate builder.
func (rsu *RequestStatusUpdate) Where(ps ...predicate.RequestStatus) *RequestStatusUpdate {
	rsu.mutation.Where(ps...)
	return rsu
}

// SetStatus sets the "status" field.
func (rsu *RequestStatusUpdate) SetStatus(r requeststatus.Status) *RequestStatusUpdate {
	rsu.mutation.SetStatus(r)
	return rsu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rsu *RequestStatusUpdate) SetNillableStatus(r *requeststatus.Status) *RequestStatusUpdate {
	if r != nil {
		rsu.SetStatus(*r)
	}
	return rsu
}

// SetCreatedAt sets the "created_at" field.
func (rsu *RequestStatusUpdate) SetCreatedAt(t time.Time) *RequestStatusUpdate {
	rsu.mutation.SetCreatedAt(t)
	return rsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rsu *RequestStatusUpdate) SetNillableCreatedAt(t *time.Time) *RequestStatusUpdate {
	if t != nil {
		rsu.SetCreatedAt(*t)
	}
	return rsu
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (rsu *RequestStatusUpdate) SetRequestID(id uuid.UUID) *RequestStatusUpdate {
	rsu.mutation.SetRequestID(id)
	return rsu
}

// SetRequest sets the "request" edge to the Request entity.
func (rsu *RequestStatusUpdate) SetRequest(r *Request) *RequestStatusUpdate {
	return rsu.SetRequestID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rsu *RequestStatusUpdate) SetUserID(id uuid.UUID) *RequestStatusUpdate {
	rsu.mutation.SetUserID(id)
	return rsu
}

// SetUser sets the "user" edge to the User entity.
func (rsu *RequestStatusUpdate) SetUser(u *User) *RequestStatusUpdate {
	return rsu.SetUserID(u.ID)
}

// Mutation returns the RequestStatusMutation object of the builder.
func (rsu *RequestStatusUpdate) Mutation() *RequestStatusMutation {
	return rsu.mutation
}

// ClearRequest clears the "request" edge to the Request entity.
func (rsu *RequestStatusUpdate) ClearRequest() *RequestStatusUpdate {
	rsu.mutation.ClearRequest()
	return rsu
}

// ClearUser clears the "user" edge to the User entity.
func (rsu *RequestStatusUpdate) ClearUser() *RequestStatusUpdate {
	rsu.mutation.ClearUser()
	return rsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rsu *RequestStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rsu.hooks) == 0 {
		if err = rsu.check(); err != nil {
			return 0, err
		}
		affected, err = rsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rsu.check(); err != nil {
				return 0, err
			}
			rsu.mutation = mutation
			affected, err = rsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rsu.hooks) - 1; i >= 0; i-- {
			if rsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rsu *RequestStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := rsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rsu *RequestStatusUpdate) Exec(ctx context.Context) error {
	_, err := rsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsu *RequestStatusUpdate) ExecX(ctx context.Context) {
	if err := rsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsu *RequestStatusUpdate) check() error {
	if v, ok := rsu.mutation.Status(); ok {
		if err := requeststatus.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RequestStatus.status": %w`, err)}
		}
	}
	if _, ok := rsu.mutation.RequestID(); rsu.mutation.RequestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RequestStatus.request"`)
	}
	if _, ok := rsu.mutation.UserID(); rsu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RequestStatus.user"`)
	}
	return nil
}

func (rsu *RequestStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requeststatus.Table,
			Columns: requeststatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requeststatus.FieldID,
			},
		},
	}
	if ps := rsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: requeststatus.FieldStatus,
		})
	}
	if value, ok := rsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requeststatus.FieldCreatedAt,
		})
	}
	if rsu.mutation.RequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requeststatus.RequestTable,
			Columns: []string{requeststatus.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requeststatus.RequestTable,
			Columns: []string{requeststatus.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   requeststatus.UserTable,
			Columns: []string{requeststatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   requeststatus.UserTable,
			Columns: []string{requeststatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{requeststatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RequestStatusUpdateOne is the builder for updating a single RequestStatus entity.
type RequestStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RequestStatusMutation
}

// SetStatus sets the "status" field.
func (rsuo *RequestStatusUpdateOne) SetStatus(r requeststatus.Status) *RequestStatusUpdateOne {
	rsuo.mutation.SetStatus(r)
	return rsuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rsuo *RequestStatusUpdateOne) SetNillableStatus(r *requeststatus.Status) *RequestStatusUpdateOne {
	if r != nil {
		rsuo.SetStatus(*r)
	}
	return rsuo
}

// SetCreatedAt sets the "created_at" field.
func (rsuo *RequestStatusUpdateOne) SetCreatedAt(t time.Time) *RequestStatusUpdateOne {
	rsuo.mutation.SetCreatedAt(t)
	return rsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rsuo *RequestStatusUpdateOne) SetNillableCreatedAt(t *time.Time) *RequestStatusUpdateOne {
	if t != nil {
		rsuo.SetCreatedAt(*t)
	}
	return rsuo
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (rsuo *RequestStatusUpdateOne) SetRequestID(id uuid.UUID) *RequestStatusUpdateOne {
	rsuo.mutation.SetRequestID(id)
	return rsuo
}

// SetRequest sets the "request" edge to the Request entity.
func (rsuo *RequestStatusUpdateOne) SetRequest(r *Request) *RequestStatusUpdateOne {
	return rsuo.SetRequestID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rsuo *RequestStatusUpdateOne) SetUserID(id uuid.UUID) *RequestStatusUpdateOne {
	rsuo.mutation.SetUserID(id)
	return rsuo
}

// SetUser sets the "user" edge to the User entity.
func (rsuo *RequestStatusUpdateOne) SetUser(u *User) *RequestStatusUpdateOne {
	return rsuo.SetUserID(u.ID)
}

// Mutation returns the RequestStatusMutation object of the builder.
func (rsuo *RequestStatusUpdateOne) Mutation() *RequestStatusMutation {
	return rsuo.mutation
}

// ClearRequest clears the "request" edge to the Request entity.
func (rsuo *RequestStatusUpdateOne) ClearRequest() *RequestStatusUpdateOne {
	rsuo.mutation.ClearRequest()
	return rsuo
}

// ClearUser clears the "user" edge to the User entity.
func (rsuo *RequestStatusUpdateOne) ClearUser() *RequestStatusUpdateOne {
	rsuo.mutation.ClearUser()
	return rsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rsuo *RequestStatusUpdateOne) Select(field string, fields ...string) *RequestStatusUpdateOne {
	rsuo.fields = append([]string{field}, fields...)
	return rsuo
}

// Save executes the query and returns the updated RequestStatus entity.
func (rsuo *RequestStatusUpdateOne) Save(ctx context.Context) (*RequestStatus, error) {
	var (
		err  error
		node *RequestStatus
	)
	if len(rsuo.hooks) == 0 {
		if err = rsuo.check(); err != nil {
			return nil, err
		}
		node, err = rsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rsuo.check(); err != nil {
				return nil, err
			}
			rsuo.mutation = mutation
			node, err = rsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rsuo.hooks) - 1; i >= 0; i-- {
			if rsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rsuo *RequestStatusUpdateOne) SaveX(ctx context.Context) *RequestStatus {
	node, err := rsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rsuo *RequestStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := rsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsuo *RequestStatusUpdateOne) ExecX(ctx context.Context) {
	if err := rsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsuo *RequestStatusUpdateOne) check() error {
	if v, ok := rsuo.mutation.Status(); ok {
		if err := requeststatus.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RequestStatus.status": %w`, err)}
		}
	}
	if _, ok := rsuo.mutation.RequestID(); rsuo.mutation.RequestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RequestStatus.request"`)
	}
	if _, ok := rsuo.mutation.UserID(); rsuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RequestStatus.user"`)
	}
	return nil
}

func (rsuo *RequestStatusUpdateOne) sqlSave(ctx context.Context) (_node *RequestStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requeststatus.Table,
			Columns: requeststatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requeststatus.FieldID,
			},
		},
	}
	id, ok := rsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RequestStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, requeststatus.FieldID)
		for _, f := range fields {
			if !requeststatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != requeststatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: requeststatus.FieldStatus,
		})
	}
	if value, ok := rsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requeststatus.FieldCreatedAt,
		})
	}
	if rsuo.mutation.RequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requeststatus.RequestTable,
			Columns: []string{requeststatus.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requeststatus.RequestTable,
			Columns: []string{requeststatus.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   requeststatus.UserTable,
			Columns: []string{requeststatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   requeststatus.UserTable,
			Columns: []string{requeststatus.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RequestStatus{config: rsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{requeststatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
