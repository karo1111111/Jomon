// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupbudget"
)

// GroupBudgetCreate is the builder for creating a GroupBudget entity.
type GroupBudgetCreate struct {
	config
	mutation *GroupBudgetMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (gbc *GroupBudgetCreate) SetAmount(i int) *GroupBudgetCreate {
	gbc.mutation.SetAmount(i)
	return gbc
}

// SetComment sets the "comment" field.
func (gbc *GroupBudgetCreate) SetComment(s string) *GroupBudgetCreate {
	gbc.mutation.SetComment(s)
	return gbc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (gbc *GroupBudgetCreate) SetNillableComment(s *string) *GroupBudgetCreate {
	if s != nil {
		gbc.SetComment(*s)
	}
	return gbc
}

// SetCreatedAt sets the "created_at" field.
func (gbc *GroupBudgetCreate) SetCreatedAt(t time.Time) *GroupBudgetCreate {
	gbc.mutation.SetCreatedAt(t)
	return gbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gbc *GroupBudgetCreate) SetNillableCreatedAt(t *time.Time) *GroupBudgetCreate {
	if t != nil {
		gbc.SetCreatedAt(*t)
	}
	return gbc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gbc *GroupBudgetCreate) SetGroupID(id int) *GroupBudgetCreate {
	gbc.mutation.SetGroupID(id)
	return gbc
}

// SetGroup sets the "group" edge to the Group entity.
func (gbc *GroupBudgetCreate) SetGroup(g *Group) *GroupBudgetCreate {
	return gbc.SetGroupID(g.ID)
}

// Mutation returns the GroupBudgetMutation object of the builder.
func (gbc *GroupBudgetCreate) Mutation() *GroupBudgetMutation {
	return gbc.mutation
}

// Save creates the GroupBudget in the database.
func (gbc *GroupBudgetCreate) Save(ctx context.Context) (*GroupBudget, error) {
	var (
		err  error
		node *GroupBudget
	)
	gbc.defaults()
	if len(gbc.hooks) == 0 {
		if err = gbc.check(); err != nil {
			return nil, err
		}
		node, err = gbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupBudgetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gbc.check(); err != nil {
				return nil, err
			}
			gbc.mutation = mutation
			node, err = gbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gbc.hooks) - 1; i >= 0; i-- {
			mut = gbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gbc *GroupBudgetCreate) SaveX(ctx context.Context) *GroupBudget {
	v, err := gbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gbc *GroupBudgetCreate) defaults() {
	if _, ok := gbc.mutation.CreatedAt(); !ok {
		v := groupbudget.DefaultCreatedAt()
		gbc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gbc *GroupBudgetCreate) check() error {
	if _, ok := gbc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if _, ok := gbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := gbc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New("ent: missing required edge \"group\"")}
	}
	return nil
}

func (gbc *GroupBudgetCreate) sqlSave(ctx context.Context) (*GroupBudget, error) {
	_node, _spec := gbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gbc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gbc *GroupBudgetCreate) createSpec() (*GroupBudget, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupBudget{config: gbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupbudget.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupbudget.FieldID,
			},
		}
	)
	if value, ok := gbc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbudget.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := gbc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupbudget.FieldComment,
		})
		_node.Comment = &value
	}
	if value, ok := gbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupbudget.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := gbc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupbudget.GroupTable,
			Columns: []string{groupbudget.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_group_budget = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupBudgetCreateBulk is the builder for creating many GroupBudget entities in bulk.
type GroupBudgetCreateBulk struct {
	config
	builders []*GroupBudgetCreate
}

// Save creates the GroupBudget entities in the database.
func (gbcb *GroupBudgetCreateBulk) Save(ctx context.Context) ([]*GroupBudget, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gbcb.builders))
	nodes := make([]*GroupBudget, len(gbcb.builders))
	mutators := make([]Mutator, len(gbcb.builders))
	for i := range gbcb.builders {
		func(i int, root context.Context) {
			builder := gbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupBudgetMutation)
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
					_, err = mutators[i+1].Mutate(root, gbcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gbcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gbcb *GroupBudgetCreateBulk) SaveX(ctx context.Context) []*GroupBudget {
	v, err := gbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
