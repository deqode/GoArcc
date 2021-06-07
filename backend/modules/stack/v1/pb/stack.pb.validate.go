// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/stack.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// define the regex for a UUID once up-front
var _stack_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Stack with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Stack) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 100 {
		return StackValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetSlug()); l < 1 || l > 100 {
		return StackValidationError{
			field:  "Slug",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
	}

	if err := m._validateUuid(m.GetAccountId()); err != nil {
		return StackValidationError{
			field:  "AccountId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCloudConnectionId()); err != nil {
		return StackValidationError{
			field:  "CloudConnectionId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetVcsConnectionId()); err != nil {
		return StackValidationError{
			field:  "VcsConnectionId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if uri, err := url.Parse(m.GetGitUrl()); err != nil {
		return StackValidationError{
			field:  "GitUrl",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return StackValidationError{
			field:  "GitUrl",
			reason: "value must be absolute",
		}
	}

	if utf8.RuneCountInString(m.GetGitBranch()) < 1 {
		return StackValidationError{
			field:  "GitBranch",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Environment

	// no validation rules for StackType

	if v, ok := interface{}(m.GetArchivedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StackValidationError{
				field:  "ArchivedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReleasedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StackValidationError{
				field:  "ReleasedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for WebUrl

	return nil
}

func (m *Stack) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// StackValidationError is the validation error returned by Stack.Validate if
// the designated constraints aren't met.
type StackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StackValidationError) ErrorName() string { return "StackValidationError" }

// Error satisfies the builtin error interface
func (e StackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStack.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StackValidationError{}

// Validate checks the field values on CreateStackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateStackRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetStack()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateStackRequestValidationError{
				field:  "Stack",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateStackRequestValidationError is the validation error returned by
// CreateStackRequest.Validate if the designated constraints aren't met.
type CreateStackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStackRequestValidationError) ErrorName() string {
	return "CreateStackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateStackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStackRequestValidationError{}

// Validate checks the field values on DeleteStackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteStackRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return DeleteStackRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *DeleteStackRequest) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteStackRequestValidationError is the validation error returned by
// DeleteStackRequest.Validate if the designated constraints aren't met.
type DeleteStackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteStackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteStackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteStackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteStackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteStackRequestValidationError) ErrorName() string {
	return "DeleteStackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteStackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteStackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteStackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteStackRequestValidationError{}

// Validate checks the field values on UpdateStackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateStackRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStack() == nil {
		return UpdateStackRequestValidationError{
			field:  "Stack",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetStack()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateStackRequestValidationError{
				field:  "Stack",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateStackRequestValidationError is the validation error returned by
// UpdateStackRequest.Validate if the designated constraints aren't met.
type UpdateStackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStackRequestValidationError) ErrorName() string {
	return "UpdateStackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStackRequestValidationError{}

// Validate checks the field values on GetStackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetStackRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return GetStackRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *GetStackRequest) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetStackRequestValidationError is the validation error returned by
// GetStackRequest.Validate if the designated constraints aren't met.
type GetStackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStackRequestValidationError) ErrorName() string { return "GetStackRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetStackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStackRequestValidationError{}

// Validate checks the field values on ListStackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListStackRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetAccountId()); err != nil {
		return ListStackRequestValidationError{
			field:  "AccountId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *ListStackRequest) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ListStackRequestValidationError is the validation error returned by
// ListStackRequest.Validate if the designated constraints aren't met.
type ListStackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStackRequestValidationError) ErrorName() string { return "ListStackRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListStackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStackRequestValidationError{}

// Validate checks the field values on ListStackResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListStackResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStacks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListStackResponseValidationError{
					field:  fmt.Sprintf("Stacks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListStackResponseValidationError is the validation error returned by
// ListStackResponse.Validate if the designated constraints aren't met.
type ListStackResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStackResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStackResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStackResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStackResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStackResponseValidationError) ErrorName() string {
	return "ListStackResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListStackResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStackResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStackResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStackResponseValidationError{}

// Validate checks the field values on StackBuild with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StackBuild) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if err := m._validateUuid(m.GetStackId()); err != nil {
		return StackBuildValidationError{
			field:  "StackId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for Status

	if l := utf8.RuneCountInString(m.GetSlug()); l < 1 || l > 100 {
		return StackBuildValidationError{
			field:  "Slug",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
	}

	if err := m._validateUuid(m.GetAccountId()); err != nil {
		return StackBuildValidationError{
			field:  "AccountId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for LogStreamChannel

	if v, ok := interface{}(m.GetCompletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StackBuildValidationError{
				field:  "CompletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

func (m *StackBuild) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// StackBuildValidationError is the validation error returned by
// StackBuild.Validate if the designated constraints aren't met.
type StackBuildValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StackBuildValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StackBuildValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StackBuildValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StackBuildValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StackBuildValidationError) ErrorName() string { return "StackBuildValidationError" }

// Error satisfies the builtin error interface
func (e StackBuildValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStackBuild.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StackBuildValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StackBuildValidationError{}

// Validate checks the field values on CreateStackBuildRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateStackBuildRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetStackBuild()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateStackBuildRequestValidationError{
				field:  "StackBuild",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateStackBuildRequestValidationError is the validation error returned by
// CreateStackBuildRequest.Validate if the designated constraints aren't met.
type CreateStackBuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStackBuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStackBuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStackBuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStackBuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStackBuildRequestValidationError) ErrorName() string {
	return "CreateStackBuildRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateStackBuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStackBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStackBuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStackBuildRequestValidationError{}

// Validate checks the field values on DeleteStackBuildRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteStackBuildRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return DeleteStackBuildRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for DeletePhysical

	return nil
}

func (m *DeleteStackBuildRequest) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteStackBuildRequestValidationError is the validation error returned by
// DeleteStackBuildRequest.Validate if the designated constraints aren't met.
type DeleteStackBuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteStackBuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteStackBuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteStackBuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteStackBuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteStackBuildRequestValidationError) ErrorName() string {
	return "DeleteStackBuildRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteStackBuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteStackBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteStackBuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteStackBuildRequestValidationError{}

// Validate checks the field values on GetStackBuildRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetStackBuildRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return GetStackBuildRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *GetStackBuildRequest) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetStackBuildRequestValidationError is the validation error returned by
// GetStackBuildRequest.Validate if the designated constraints aren't met.
type GetStackBuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStackBuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStackBuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStackBuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStackBuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStackBuildRequestValidationError) ErrorName() string {
	return "GetStackBuildRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetStackBuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStackBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStackBuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStackBuildRequestValidationError{}

// Validate checks the field values on ListStackBuildRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListStackBuildRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetStackId()); err != nil {
		return ListStackBuildRequestValidationError{
			field:  "StackId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *ListStackBuildRequest) _validateUuid(uuid string) error {
	if matched := _stack_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ListStackBuildRequestValidationError is the validation error returned by
// ListStackBuildRequest.Validate if the designated constraints aren't met.
type ListStackBuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStackBuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStackBuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStackBuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStackBuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStackBuildRequestValidationError) ErrorName() string {
	return "ListStackBuildRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListStackBuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStackBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStackBuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStackBuildRequestValidationError{}

// Validate checks the field values on ListStackBuildResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListStackBuildResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStackBuilds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListStackBuildResponseValidationError{
					field:  fmt.Sprintf("StackBuilds[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListStackBuildResponseValidationError is the validation error returned by
// ListStackBuildResponse.Validate if the designated constraints aren't met.
type ListStackBuildResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStackBuildResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStackBuildResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStackBuildResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStackBuildResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStackBuildResponseValidationError) ErrorName() string {
	return "ListStackBuildResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListStackBuildResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStackBuildResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStackBuildResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStackBuildResponseValidationError{}
