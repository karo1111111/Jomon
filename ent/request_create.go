// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/comment"
	"github.com/traPtitech/Jomon/ent/file"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requeststatus"
	"github.com/traPtitech/Jomon/ent/requesttarget"
	"github.com/traPtitech/Jomon/ent/tag"
	"github.com/traPtitech/Jomon/ent/transaction"
	"github.com/traPtitech/Jomon/ent/user"
)

// RequestCreate is the builder for creating a Request entity.
type RequestCreate struct {
	config
	mutation *RequestMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (rc *RequestCreate) SetAmount(i int) *RequestCreate {
	rc.mutation.SetAmount(i)
	return rc
}

// SetTitle sets the "title" field.
func (rc *RequestCreate) SetTitle(s string) *RequestCreate {
	rc.mutation.SetTitle(s)
	return rc
}

// SetContent sets the "content" field.
func (rc *RequestCreate) SetContent(s string) *RequestCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RequestCreate) SetCreatedAt(t time.Time) *RequestCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RequestCreate) SetNillableCreatedAt(t *time.Time) *RequestCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RequestCreate) SetUpdatedAt(t time.Time) *RequestCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RequestCreate) SetNillableUpdatedAt(t *time.Time) *RequestCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RequestCreate) SetID(u uuid.UUID) *RequestCreate {
	rc.mutation.SetID(u)
	return rc
}

// AddStatuIDs adds the "status" edge to the RequestStatus entity by IDs.
func (rc *RequestCreate) AddStatuIDs(ids ...uuid.UUID) *RequestCreate {
	rc.mutation.AddStatuIDs(ids...)
	return rc
}

// AddStatus adds the "status" edges to the RequestStatus entity.
func (rc *RequestCreate) AddStatus(r ...*RequestStatus) *RequestCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddStatuIDs(ids...)
}

// AddTargetIDs adds the "target" edge to the RequestTarget entity by IDs.
func (rc *RequestCreate) AddTargetIDs(ids ...uuid.UUID) *RequestCreate {
	rc.mutation.AddTargetIDs(ids...)
	return rc
}

// AddTarget adds the "target" edges to the RequestTarget entity.
func (rc *RequestCreate) AddTarget(r ...*RequestTarget) *RequestCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddTargetIDs(ids...)
}

// AddFileIDs adds the "file" edge to the File entity by IDs.
func (rc *RequestCreate) AddFileIDs(ids ...uuid.UUID) *RequestCreate {
	rc.mutation.AddFileIDs(ids...)
	return rc
}

// AddFile adds the "file" edges to the File entity.
func (rc *RequestCreate) AddFile(f ...*File) *RequestCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return rc.AddFileIDs(ids...)
}

// AddTagIDs adds the "tag" edge to the Tag entity by IDs.
func (rc *RequestCreate) AddTagIDs(ids ...uuid.UUID) *RequestCreate {
	rc.mutation.AddTagIDs(ids...)
	return rc
}

// AddTag adds the "tag" edges to the Tag entity.
func (rc *RequestCreate) AddTag(t ...*Tag) *RequestCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddTagIDs(ids...)
}

// AddTransactionIDs adds the "transaction" edge to the Transaction entity by IDs.
func (rc *RequestCreate) AddTransactionIDs(ids ...uuid.UUID) *RequestCreate {
	rc.mutation.AddTransactionIDs(ids...)
	return rc
}

// AddTransaction adds the "transaction" edges to the Transaction entity.
func (rc *RequestCreate) AddTransaction(t ...*Transaction) *RequestCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddTransactionIDs(ids...)
}

// AddCommentIDs adds the "comment" edge to the Comment entity by IDs.
func (rc *RequestCreate) AddCommentIDs(ids ...uuid.UUID) *RequestCreate {
	rc.mutation.AddCommentIDs(ids...)
	return rc
}

// AddComment adds the "comment" edges to the Comment entity.
func (rc *RequestCreate) AddComment(c ...*Comment) *RequestCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddCommentIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rc *RequestCreate) SetUserID(id uuid.UUID) *RequestCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (rc *RequestCreate) SetNillableUserID(id *uuid.UUID) *RequestCreate {
	if id != nil {
		rc = rc.SetUserID(*id)
	}
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RequestCreate) SetUser(u *User) *RequestCreate {
	return rc.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (rc *RequestCreate) SetGroupID(id uuid.UUID) *RequestCreate {
	rc.mutation.SetGroupID(id)
	return rc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (rc *RequestCreate) SetNillableGroupID(id *uuid.UUID) *RequestCreate {
	if id != nil {
		rc = rc.SetGroupID(*id)
	}
	return rc
}

// SetGroup sets the "group" edge to the Group entity.
func (rc *RequestCreate) SetGroup(g *Group) *RequestCreate {
	return rc.SetGroupID(g.ID)
}

// Mutation returns the RequestMutation object of the builder.
func (rc *RequestCreate) Mutation() *RequestMutation {
	return rc.mutation
}

// Save creates the Request in the database.
func (rc *RequestCreate) Save(ctx context.Context) (*Request, error) {
	var (
		err  error
		node *Request
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RequestCreate) SaveX(ctx context.Context) *Request {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (rc *RequestCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := request.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := request.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := request.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RequestCreate) check() error {
	if _, ok := rc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if _, ok := rc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New("ent: missing required field \"content\"")}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (rc *RequestCreate) sqlSave(ctx context.Context) (*Request, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (rc *RequestCreate) createSpec() (*Request, *sqlgraph.CreateSpec) {
	var (
		_node = &Request{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: request.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: request.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: request.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := rc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: request.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: request.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: request.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: request.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := rc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   request.StatusTable,
			Columns: []string{request.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: requeststatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   request.TargetTable,
			Columns: []string{request.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: requesttarget.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   request.FileTable,
			Columns: []string{request.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   request.TagTable,
			Columns: []string{request.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   request.TransactionTable,
			Columns: []string{request.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   request.CommentTable,
			Columns: []string{request.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   request.UserTable,
			Columns: []string{request.UserColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   request.GroupTable,
			Columns: []string{request.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_request = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RequestCreateBulk is the builder for creating many Request entities in bulk.
type RequestCreateBulk struct {
	config
	builders []*RequestCreate
}

// Save creates the Request entities in the database.
func (rcb *RequestCreateBulk) Save(ctx context.Context) ([]*Request, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Request, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RequestCreateBulk) SaveX(ctx context.Context) []*Request {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
