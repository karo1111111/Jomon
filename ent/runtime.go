// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/comment"
	"github.com/traPtitech/Jomon/ent/file"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupbudget"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requeststatus"
	"github.com/traPtitech/Jomon/ent/requesttarget"
	"github.com/traPtitech/Jomon/ent/schema"
	"github.com/traPtitech/Jomon/ent/tag"
	"github.com/traPtitech/Jomon/ent/transaction"
	"github.com/traPtitech/Jomon/ent/transactiondetail"
	"github.com/traPtitech/Jomon/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[2].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentFields[3].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.DefaultID holds the default value on creation for the id field.
	comment.DefaultID = commentDescID.Default.(func() uuid.UUID)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescName is the schema descriptor for name field.
	fileDescName := fileFields[1].Descriptor()
	// file.NameValidator is a validator for the "name" field. It is called by the builders before save.
	file.NameValidator = fileDescName.Validators[0].(func(string) error)
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileFields[3].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescID is the schema descriptor for id field.
	fileDescID := fileFields[0].Descriptor()
	// file.DefaultID holds the default value on creation for the id field.
	file.DefaultID = fileDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[4].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[5].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	groupbudgetFields := schema.GroupBudget{}.Fields()
	_ = groupbudgetFields
	// groupbudgetDescCreatedAt is the schema descriptor for created_at field.
	groupbudgetDescCreatedAt := groupbudgetFields[3].Descriptor()
	// groupbudget.DefaultCreatedAt holds the default value on creation for the created_at field.
	groupbudget.DefaultCreatedAt = groupbudgetDescCreatedAt.Default.(func() time.Time)
	// groupbudgetDescID is the schema descriptor for id field.
	groupbudgetDescID := groupbudgetFields[0].Descriptor()
	// groupbudget.DefaultID holds the default value on creation for the id field.
	groupbudget.DefaultID = groupbudgetDescID.Default.(func() uuid.UUID)
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescCreatedAt is the schema descriptor for created_at field.
	requestDescCreatedAt := requestFields[4].Descriptor()
	// request.DefaultCreatedAt holds the default value on creation for the created_at field.
	request.DefaultCreatedAt = requestDescCreatedAt.Default.(func() time.Time)
	// requestDescUpdatedAt is the schema descriptor for updated_at field.
	requestDescUpdatedAt := requestFields[5].Descriptor()
	// request.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	request.DefaultUpdatedAt = requestDescUpdatedAt.Default.(func() time.Time)
	// request.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	request.UpdateDefaultUpdatedAt = requestDescUpdatedAt.UpdateDefault.(func() time.Time)
	// requestDescID is the schema descriptor for id field.
	requestDescID := requestFields[0].Descriptor()
	// request.DefaultID holds the default value on creation for the id field.
	request.DefaultID = requestDescID.Default.(func() uuid.UUID)
	requeststatusFields := schema.RequestStatus{}.Fields()
	_ = requeststatusFields
	// requeststatusDescCreatedAt is the schema descriptor for created_at field.
	requeststatusDescCreatedAt := requeststatusFields[2].Descriptor()
	// requeststatus.DefaultCreatedAt holds the default value on creation for the created_at field.
	requeststatus.DefaultCreatedAt = requeststatusDescCreatedAt.Default.(func() time.Time)
	// requeststatusDescID is the schema descriptor for id field.
	requeststatusDescID := requeststatusFields[0].Descriptor()
	// requeststatus.DefaultID holds the default value on creation for the id field.
	requeststatus.DefaultID = requeststatusDescID.Default.(func() uuid.UUID)
	requesttargetFields := schema.RequestTarget{}.Fields()
	_ = requesttargetFields
	// requesttargetDescCreatedAt is the schema descriptor for created_at field.
	requesttargetDescCreatedAt := requesttargetFields[4].Descriptor()
	// requesttarget.DefaultCreatedAt holds the default value on creation for the created_at field.
	requesttarget.DefaultCreatedAt = requesttargetDescCreatedAt.Default.(func() time.Time)
	// requesttargetDescID is the schema descriptor for id field.
	requesttargetDescID := requesttargetFields[0].Descriptor()
	// requesttarget.DefaultID holds the default value on creation for the id field.
	requesttarget.DefaultID = requesttargetDescID.Default.(func() uuid.UUID)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[1].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[3].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagFields[4].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagFields[0].Descriptor()
	// tag.DefaultID holds the default value on creation for the id field.
	tag.DefaultID = tagDescID.Default.(func() uuid.UUID)
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescCreatedAt is the schema descriptor for created_at field.
	transactionDescCreatedAt := transactionFields[1].Descriptor()
	// transaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	transaction.DefaultCreatedAt = transactionDescCreatedAt.Default.(func() time.Time)
	// transactionDescID is the schema descriptor for id field.
	transactionDescID := transactionFields[0].Descriptor()
	// transaction.DefaultID holds the default value on creation for the id field.
	transaction.DefaultID = transactionDescID.Default.(func() uuid.UUID)
	transactiondetailFields := schema.TransactionDetail{}.Fields()
	_ = transactiondetailFields
	// transactiondetailDescAmount is the schema descriptor for amount field.
	transactiondetailDescAmount := transactiondetailFields[1].Descriptor()
	// transactiondetail.DefaultAmount holds the default value on creation for the amount field.
	transactiondetail.DefaultAmount = transactiondetailDescAmount.Default.(int)
	// transactiondetailDescTarget is the schema descriptor for target field.
	transactiondetailDescTarget := transactiondetailFields[2].Descriptor()
	// transactiondetail.DefaultTarget holds the default value on creation for the target field.
	transactiondetail.DefaultTarget = transactiondetailDescTarget.Default.(string)
	// transactiondetailDescCreatedAt is the schema descriptor for created_at field.
	transactiondetailDescCreatedAt := transactiondetailFields[3].Descriptor()
	// transactiondetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	transactiondetail.DefaultCreatedAt = transactiondetailDescCreatedAt.Default.(func() time.Time)
	// transactiondetailDescUpdatedAt is the schema descriptor for updated_at field.
	transactiondetailDescUpdatedAt := transactiondetailFields[4].Descriptor()
	// transactiondetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	transactiondetail.DefaultUpdatedAt = transactiondetailDescUpdatedAt.Default.(func() time.Time)
	// transactiondetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	transactiondetail.UpdateDefaultUpdatedAt = transactiondetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// transactiondetailDescID is the schema descriptor for id field.
	transactiondetailDescID := transactiondetailFields[0].Descriptor()
	// transactiondetail.DefaultID holds the default value on creation for the id field.
	transactiondetail.DefaultID = transactiondetailDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[3].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
