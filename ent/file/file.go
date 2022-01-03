// Code generated by entc, DO NOT EDIT.

package file

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMimeType holds the string denoting the mime_type field in the database.
	FieldMimeType = "mime_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeRequest holds the string denoting the request edge name in mutations.
	EdgeRequest = "request"
	// Table holds the table name of the file in the database.
	Table = "files"
	// RequestTable is the table that holds the request relation/edge.
	RequestTable = "files"
	// RequestInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestInverseTable = "requests"
	// RequestColumn is the table column denoting the request relation/edge.
	RequestColumn = "request_file"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMimeType,
	FieldCreatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"request_file",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
