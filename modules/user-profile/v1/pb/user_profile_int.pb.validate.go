// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/user_profile_int.proto

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

// Validate checks the field values on CreateUserProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateUserProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUserProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserProfileRequestValidationError{
				field:  "UserProfile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateUserProfileRequestValidationError is the validation error returned by
// CreateUserProfileRequest.Validate if the designated constraints aren't met.
type CreateUserProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserProfileRequestValidationError) ErrorName() string {
	return "CreateUserProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserProfileRequestValidationError{}

// Validate checks the field values on DeleteUserProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteUserProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 3 {
		return DeleteUserProfileRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// DeleteUserProfileRequestValidationError is the validation error returned by
// DeleteUserProfileRequest.Validate if the designated constraints aren't met.
type DeleteUserProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserProfileRequestValidationError) ErrorName() string {
	return "DeleteUserProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserProfileRequestValidationError{}

// Validate checks the field values on UpdateUserProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateUserProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserProfile() == nil {
		return UpdateUserProfileRequestValidationError{
			field:  "UserProfile",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUserProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserProfileRequestValidationError{
				field:  "UserProfile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateUserProfileRequestValidationError is the validation error returned by
// UpdateUserProfileRequest.Validate if the designated constraints aren't met.
type UpdateUserProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserProfileRequestValidationError) ErrorName() string {
	return "UpdateUserProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserProfileRequestValidationError{}
