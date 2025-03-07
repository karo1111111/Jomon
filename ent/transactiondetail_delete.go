// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/predicate"
	"github.com/traPtitech/Jomon/ent/transactiondetail"
)

// TransactionDetailDelete is the builder for deleting a TransactionDetail entity.
type TransactionDetailDelete struct {
	config
	hooks    []Hook
	mutation *TransactionDetailMutation
}

// Where appends a list predicates to the TransactionDetailDelete builder.
func (tdd *TransactionDetailDelete) Where(ps ...predicate.TransactionDetail) *TransactionDetailDelete {
	tdd.mutation.Where(ps...)
	return tdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tdd *TransactionDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tdd.hooks) == 0 {
		affected, err = tdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tdd.mutation = mutation
			affected, err = tdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tdd.hooks) - 1; i >= 0; i-- {
			if tdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdd *TransactionDetailDelete) ExecX(ctx context.Context) int {
	n, err := tdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tdd *TransactionDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: transactiondetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transactiondetail.FieldID,
			},
		},
	}
	if ps := tdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TransactionDetailDeleteOne is the builder for deleting a single TransactionDetail entity.
type TransactionDetailDeleteOne struct {
	tdd *TransactionDetailDelete
}

// Exec executes the deletion query.
func (tddo *TransactionDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := tddo.tdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{transactiondetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tddo *TransactionDetailDeleteOne) ExecX(ctx context.Context) {
	tddo.tdd.ExecX(ctx)
}
