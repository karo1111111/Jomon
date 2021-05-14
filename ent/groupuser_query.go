// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupuser"
	"github.com/traPtitech/Jomon/ent/predicate"
)

// GroupUserQuery is the builder for querying GroupUser entities.
type GroupUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GroupUser
	// eager-loading edges.
	withGroup *GroupQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupUserQuery builder.
func (guq *GroupUserQuery) Where(ps ...predicate.GroupUser) *GroupUserQuery {
	guq.predicates = append(guq.predicates, ps...)
	return guq
}

// Limit adds a limit step to the query.
func (guq *GroupUserQuery) Limit(limit int) *GroupUserQuery {
	guq.limit = &limit
	return guq
}

// Offset adds an offset step to the query.
func (guq *GroupUserQuery) Offset(offset int) *GroupUserQuery {
	guq.offset = &offset
	return guq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (guq *GroupUserQuery) Unique(unique bool) *GroupUserQuery {
	guq.unique = &unique
	return guq
}

// Order adds an order step to the query.
func (guq *GroupUserQuery) Order(o ...OrderFunc) *GroupUserQuery {
	guq.order = append(guq.order, o...)
	return guq
}

// QueryGroup chains the current query on the "group" edge.
func (guq *GroupUserQuery) QueryGroup() *GroupQuery {
	query := &GroupQuery{config: guq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := guq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := guq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupuser.Table, groupuser.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, groupuser.GroupTable, groupuser.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(guq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupUser entity from the query.
// Returns a *NotFoundError when no GroupUser was found.
func (guq *GroupUserQuery) First(ctx context.Context) (*GroupUser, error) {
	nodes, err := guq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (guq *GroupUserQuery) FirstX(ctx context.Context) *GroupUser {
	node, err := guq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupUser ID from the query.
// Returns a *NotFoundError when no GroupUser ID was found.
func (guq *GroupUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = guq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (guq *GroupUserQuery) FirstIDX(ctx context.Context) int {
	id, err := guq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GroupUser entity is not found.
// Returns a *NotFoundError when no GroupUser entities are found.
func (guq *GroupUserQuery) Only(ctx context.Context) (*GroupUser, error) {
	nodes, err := guq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupuser.Label}
	default:
		return nil, &NotSingularError{groupuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (guq *GroupUserQuery) OnlyX(ctx context.Context) *GroupUser {
	node, err := guq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupUser ID in the query.
// Returns a *NotSingularError when exactly one GroupUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (guq *GroupUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = guq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = &NotSingularError{groupuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (guq *GroupUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := guq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupUsers.
func (guq *GroupUserQuery) All(ctx context.Context) ([]*GroupUser, error) {
	if err := guq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return guq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (guq *GroupUserQuery) AllX(ctx context.Context) []*GroupUser {
	nodes, err := guq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupUser IDs.
func (guq *GroupUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := guq.Select(groupuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (guq *GroupUserQuery) IDsX(ctx context.Context) []int {
	ids, err := guq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (guq *GroupUserQuery) Count(ctx context.Context) (int, error) {
	if err := guq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return guq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (guq *GroupUserQuery) CountX(ctx context.Context) int {
	count, err := guq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (guq *GroupUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := guq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return guq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (guq *GroupUserQuery) ExistX(ctx context.Context) bool {
	exist, err := guq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (guq *GroupUserQuery) Clone() *GroupUserQuery {
	if guq == nil {
		return nil
	}
	return &GroupUserQuery{
		config:     guq.config,
		limit:      guq.limit,
		offset:     guq.offset,
		order:      append([]OrderFunc{}, guq.order...),
		predicates: append([]predicate.GroupUser{}, guq.predicates...),
		withGroup:  guq.withGroup.Clone(),
		// clone intermediate query.
		sql:  guq.sql.Clone(),
		path: guq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (guq *GroupUserQuery) WithGroup(opts ...func(*GroupQuery)) *GroupUserQuery {
	query := &GroupQuery{config: guq.config}
	for _, opt := range opts {
		opt(query)
	}
	guq.withGroup = query
	return guq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupUser.Query().
//		GroupBy(groupuser.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (guq *GroupUserQuery) GroupBy(field string, fields ...string) *GroupUserGroupBy {
	group := &GroupUserGroupBy{config: guq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := guq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return guq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.GroupUser.Query().
//		Select(groupuser.FieldUserID).
//		Scan(ctx, &v)
//
func (guq *GroupUserQuery) Select(field string, fields ...string) *GroupUserSelect {
	guq.fields = append([]string{field}, fields...)
	return &GroupUserSelect{GroupUserQuery: guq}
}

func (guq *GroupUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range guq.fields {
		if !groupuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if guq.path != nil {
		prev, err := guq.path(ctx)
		if err != nil {
			return err
		}
		guq.sql = prev
	}
	return nil
}

func (guq *GroupUserQuery) sqlAll(ctx context.Context) ([]*GroupUser, error) {
	var (
		nodes       = []*GroupUser{}
		withFKs     = guq.withFKs
		_spec       = guq.querySpec()
		loadedTypes = [1]bool{
			guq.withGroup != nil,
		}
	)
	if guq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, groupuser.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GroupUser{config: guq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, guq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := guq.withGroup; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*GroupUser)
		for i := range nodes {
			if nodes[i].group_user == nil {
				continue
			}
			fk := *nodes[i].group_user
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(group.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Group = n
			}
		}
	}

	return nodes, nil
}

func (guq *GroupUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := guq.querySpec()
	return sqlgraph.CountNodes(ctx, guq.driver, _spec)
}

func (guq *GroupUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := guq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (guq *GroupUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupuser.Table,
			Columns: groupuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupuser.FieldID,
			},
		},
		From:   guq.sql,
		Unique: true,
	}
	if unique := guq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := guq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupuser.FieldID)
		for i := range fields {
			if fields[i] != groupuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := guq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := guq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := guq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := guq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (guq *GroupUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(guq.driver.Dialect())
	t1 := builder.Table(groupuser.Table)
	selector := builder.Select(t1.Columns(groupuser.Columns...)...).From(t1)
	if guq.sql != nil {
		selector = guq.sql
		selector.Select(selector.Columns(groupuser.Columns...)...)
	}
	for _, p := range guq.predicates {
		p(selector)
	}
	for _, p := range guq.order {
		p(selector)
	}
	if offset := guq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := guq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupUserGroupBy is the group-by builder for GroupUser entities.
type GroupUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gugb *GroupUserGroupBy) Aggregate(fns ...AggregateFunc) *GroupUserGroupBy {
	gugb.fns = append(gugb.fns, fns...)
	return gugb
}

// Scan applies the group-by query and scans the result into the given value.
func (gugb *GroupUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gugb.path(ctx)
	if err != nil {
		return err
	}
	gugb.sql = query
	return gugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gugb *GroupUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GroupUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gugb *GroupUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := gugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gugb *GroupUserGroupBy) StringX(ctx context.Context) string {
	v, err := gugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GroupUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gugb *GroupUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := gugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gugb *GroupUserGroupBy) IntX(ctx context.Context) int {
	v, err := gugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GroupUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gugb *GroupUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gugb *GroupUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GroupUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gugb *GroupUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GroupUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gugb *GroupUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := gugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gugb *GroupUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gugb.fields {
		if !groupuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gugb *GroupUserGroupBy) sqlQuery() *sql.Selector {
	selector := gugb.sql
	columns := make([]string, 0, len(gugb.fields)+len(gugb.fns))
	columns = append(columns, gugb.fields...)
	for _, fn := range gugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(gugb.fields...)
}

// GroupUserSelect is the builder for selecting fields of GroupUser entities.
type GroupUserSelect struct {
	*GroupUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gus *GroupUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gus.prepareQuery(ctx); err != nil {
		return err
	}
	gus.sql = gus.GroupUserQuery.sqlQuery(ctx)
	return gus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gus *GroupUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GroupUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gus *GroupUserSelect) StringsX(ctx context.Context) []string {
	v, err := gus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gus *GroupUserSelect) StringX(ctx context.Context) string {
	v, err := gus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GroupUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gus *GroupUserSelect) IntsX(ctx context.Context) []int {
	v, err := gus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gus *GroupUserSelect) IntX(ctx context.Context) int {
	v, err := gus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GroupUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gus *GroupUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gus *GroupUserSelect) Float64X(ctx context.Context) float64 {
	v, err := gus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GroupUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gus *GroupUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := gus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gus *GroupUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{groupuser.Label}
	default:
		err = fmt.Errorf("ent: GroupUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gus *GroupUserSelect) BoolX(ctx context.Context) bool {
	v, err := gus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gus *GroupUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gus.sqlQuery().Query()
	if err := gus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gus *GroupUserSelect) sqlQuery() sql.Querier {
	selector := gus.sql
	selector.Select(selector.Columns(gus.fields...)...)
	return selector
}
