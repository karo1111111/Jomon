// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requesttag"
	"github.com/traPtitech/Jomon/ent/tag"
)

// RequestTag is the model entity for the RequestTag schema.
type RequestTag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RequestTagQuery when eager-loading is set.
	Edges           RequestTagEdges `json:"edges"`
	request_tag     *int
	tag_request_tag *int
}

// RequestTagEdges holds the relations/edges for other nodes in the graph.
type RequestTagEdges struct {
	// Request holds the value of the request edge.
	Request *Request `json:"request,omitempty"`
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RequestOrErr returns the Request value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestTagEdges) RequestOrErr() (*Request, error) {
	if e.loadedTypes[0] {
		if e.Request == nil {
			// The edge request was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: request.Label}
		}
		return e.Request, nil
	}
	return nil, &NotLoadedError{edge: "request"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestTagEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[1] {
		if e.Tag == nil {
			// The edge tag was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RequestTag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case requesttag.FieldID:
			values[i] = new(sql.NullInt64)
		case requesttag.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case requesttag.ForeignKeys[0]: // request_tag
			values[i] = new(sql.NullInt64)
		case requesttag.ForeignKeys[1]: // tag_request_tag
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RequestTag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RequestTag fields.
func (rt *RequestTag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case requesttag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rt.ID = int(value.Int64)
		case requesttag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rt.CreatedAt = value.Time
			}
		case requesttag.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field request_tag", value)
			} else if value.Valid {
				rt.request_tag = new(int)
				*rt.request_tag = int(value.Int64)
			}
		case requesttag.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tag_request_tag", value)
			} else if value.Valid {
				rt.tag_request_tag = new(int)
				*rt.tag_request_tag = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRequest queries the "request" edge of the RequestTag entity.
func (rt *RequestTag) QueryRequest() *RequestQuery {
	return (&RequestTagClient{config: rt.config}).QueryRequest(rt)
}

// QueryTag queries the "tag" edge of the RequestTag entity.
func (rt *RequestTag) QueryTag() *TagQuery {
	return (&RequestTagClient{config: rt.config}).QueryTag(rt)
}

// Update returns a builder for updating this RequestTag.
// Note that you need to call RequestTag.Unwrap() before calling this method if this RequestTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RequestTag) Update() *RequestTagUpdateOne {
	return (&RequestTagClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the RequestTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *RequestTag) Unwrap() *RequestTag {
	tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RequestTag is not a transactional entity")
	}
	rt.config.driver = tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RequestTag) String() string {
	var builder strings.Builder
	builder.WriteString("RequestTag(")
	builder.WriteString(fmt.Sprintf("id=%v", rt.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(rt.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RequestTags is a parsable slice of RequestTag.
type RequestTags []*RequestTag

func (rt RequestTags) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
