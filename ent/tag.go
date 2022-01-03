// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/tag"
	"github.com/traPtitech/Jomon/ent/transaction"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges           TagEdges `json:"edges"`
	transaction_tag *uuid.UUID
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Request holds the value of the request edge.
	Request []*Request `json:"request,omitempty"`
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RequestOrErr returns the Request value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) RequestOrErr() ([]*Request, error) {
	if e.loadedTypes[0] {
		return e.Request, nil
	}
	return nil, &NotLoadedError{edge: "request"}
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TagEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[1] {
		if e.Transaction == nil {
			// The edge transaction was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldName, tag.FieldDescription:
			values[i] = new(sql.NullString)
		case tag.FieldCreatedAt, tag.FieldUpdatedAt, tag.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case tag.FieldID:
			values[i] = new(uuid.UUID)
		case tag.ForeignKeys[0]: // transaction_tag
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tag.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case tag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case tag.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = new(time.Time)
				*t.DeletedAt = value.Time
			}
		case tag.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_tag", values[i])
			} else if value.Valid {
				t.transaction_tag = new(uuid.UUID)
				*t.transaction_tag = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryRequest queries the "request" edge of the Tag entity.
func (t *Tag) QueryRequest() *RequestQuery {
	return (&TagClient{config: t.config}).QueryRequest(t)
}

// QueryTransaction queries the "transaction" edge of the Tag entity.
func (t *Tag) QueryTransaction() *TransactionQuery {
	return (&TagClient{config: t.config}).QueryTransaction(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return (&TagClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	if v := t.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag

func (t Tags) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
