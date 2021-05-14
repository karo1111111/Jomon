// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/tag"
	"github.com/traPtitech/Jomon/ent/transaction"
	"github.com/traPtitech/Jomon/ent/transactiontag"
)

// TransactionTagCreate is the builder for creating a TransactionTag entity.
type TransactionTagCreate struct {
	config
	mutation *TransactionTagMutation
	hooks    []Hook
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (ttc *TransactionTagCreate) SetTransactionID(id int) *TransactionTagCreate {
	ttc.mutation.SetTransactionID(id)
	return ttc
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (ttc *TransactionTagCreate) SetTransaction(t *Transaction) *TransactionTagCreate {
	return ttc.SetTransactionID(t.ID)
}

// SetTagID sets the "tag" edge to the Tag entity by ID.
func (ttc *TransactionTagCreate) SetTagID(id int) *TransactionTagCreate {
	ttc.mutation.SetTagID(id)
	return ttc
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttc *TransactionTagCreate) SetTag(t *Tag) *TransactionTagCreate {
	return ttc.SetTagID(t.ID)
}

// Mutation returns the TransactionTagMutation object of the builder.
func (ttc *TransactionTagCreate) Mutation() *TransactionTagMutation {
	return ttc.mutation
}

// Save creates the TransactionTag in the database.
func (ttc *TransactionTagCreate) Save(ctx context.Context) (*TransactionTag, error) {
	var (
		err  error
		node *TransactionTag
	)
	if len(ttc.hooks) == 0 {
		if err = ttc.check(); err != nil {
			return nil, err
		}
		node, err = ttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttc.check(); err != nil {
				return nil, err
			}
			ttc.mutation = mutation
			node, err = ttc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttc.hooks) - 1; i >= 0; i-- {
			mut = ttc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttc *TransactionTagCreate) SaveX(ctx context.Context) *TransactionTag {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TransactionTagCreate) check() error {
	if _, ok := ttc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction", err: errors.New("ent: missing required edge \"transaction\"")}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag", err: errors.New("ent: missing required edge \"tag\"")}
	}
	return nil
}

func (ttc *TransactionTagCreate) sqlSave(ctx context.Context) (*TransactionTag, error) {
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ttc *TransactionTagCreate) createSpec() (*TransactionTag, *sqlgraph.CreateSpec) {
	var (
		_node = &TransactionTag{config: ttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transactiontag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactiontag.FieldID,
			},
		}
	)
	if nodes := ttc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactiontag.TransactionTable,
			Columns: []string{transactiontag.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.transaction_tag = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ttc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactiontag.TagTable,
			Columns: []string{transactiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.tag_transaction_tag = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionTagCreateBulk is the builder for creating many TransactionTag entities in bulk.
type TransactionTagCreateBulk struct {
	config
	builders []*TransactionTagCreate
}

// Save creates the TransactionTag entities in the database.
func (ttcb *TransactionTagCreateBulk) Save(ctx context.Context) ([]*TransactionTag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TransactionTag, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionTagMutation)
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
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttcb *TransactionTagCreateBulk) SaveX(ctx context.Context) []*TransactionTag {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
