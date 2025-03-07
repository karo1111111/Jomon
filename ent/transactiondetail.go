// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/transaction"
	"github.com/traPtitech/Jomon/ent/transactiondetail"
)

// TransactionDetail is the model entity for the TransactionDetail schema.
type TransactionDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionDetailQuery when eager-loading is set.
	Edges              TransactionDetailEdges `json:"edges"`
	transaction_detail *uuid.UUID
}

// TransactionDetailEdges holds the relations/edges for other nodes in the graph.
type TransactionDetailEdges struct {
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionDetailEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[0] {
		if e.Transaction == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransactionDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transactiondetail.FieldAmount:
			values[i] = new(sql.NullInt64)
		case transactiondetail.FieldTarget:
			values[i] = new(sql.NullString)
		case transactiondetail.FieldCreatedAt, transactiondetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case transactiondetail.FieldID:
			values[i] = new(uuid.UUID)
		case transactiondetail.ForeignKeys[0]: // transaction_detail
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type TransactionDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransactionDetail fields.
func (td *TransactionDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transactiondetail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				td.ID = *value
			}
		case transactiondetail.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				td.Amount = int(value.Int64)
			}
		case transactiondetail.FieldTarget:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target", values[i])
			} else if value.Valid {
				td.Target = value.String
			}
		case transactiondetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				td.CreatedAt = value.Time
			}
		case transactiondetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				td.UpdatedAt = value.Time
			}
		case transactiondetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_detail", values[i])
			} else if value.Valid {
				td.transaction_detail = new(uuid.UUID)
				*td.transaction_detail = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryTransaction queries the "transaction" edge of the TransactionDetail entity.
func (td *TransactionDetail) QueryTransaction() *TransactionQuery {
	return (&TransactionDetailClient{config: td.config}).QueryTransaction(td)
}

// Update returns a builder for updating this TransactionDetail.
// Note that you need to call TransactionDetail.Unwrap() before calling this method if this TransactionDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (td *TransactionDetail) Update() *TransactionDetailUpdateOne {
	return (&TransactionDetailClient{config: td.config}).UpdateOne(td)
}

// Unwrap unwraps the TransactionDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (td *TransactionDetail) Unwrap() *TransactionDetail {
	_tx, ok := td.config.driver.(*txDriver)
	if !ok {
		panic("ent: TransactionDetail is not a transactional entity")
	}
	td.config.driver = _tx.drv
	return td
}

// String implements the fmt.Stringer.
func (td *TransactionDetail) String() string {
	var builder strings.Builder
	builder.WriteString("TransactionDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", td.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", td.Amount))
	builder.WriteString(", ")
	builder.WriteString("target=")
	builder.WriteString(td.Target)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(td.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(td.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TransactionDetails is a parsable slice of TransactionDetail.
type TransactionDetails []*TransactionDetail

func (td TransactionDetails) config(cfg config) {
	for _i := range td {
		td[_i].config = cfg
	}
}
