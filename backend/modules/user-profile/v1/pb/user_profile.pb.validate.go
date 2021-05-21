// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/user_profile.proto

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

	types "alfred/protos/types"
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

	_ = types.VCSProviders(0)
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

// Validate checks the field values on GetUserProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUserProfileRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 3 {
		return GetUserProfileRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// GetUserProfileRequestValidationError is the validation error returned by
// GetUserProfileRequest.Validate if the designated constraints aren't met.
type GetUserProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserProfileRequestValidationError) ErrorName() string {
	return "GetUserProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserProfileRequestValidationError{}

// Validate checks the field values on GetUserProfileBySubRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUserProfileBySubRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetSub()) < 3 {
		return GetUserProfileBySubRequestValidationError{
			field:  "Sub",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// GetUserProfileBySubRequestValidationError is the validation error returned
// by GetUserProfileBySubRequest.Validate if the designated constraints aren't met.
type GetUserProfileBySubRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserProfileBySubRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserProfileBySubRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserProfileBySubRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserProfileBySubRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserProfileBySubRequestValidationError) ErrorName() string {
	return "GetUserProfileBySubRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserProfileBySubRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserProfileBySubRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserProfileBySubRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserProfileBySubRequestValidationError{}

// Validate checks the field values on UserProfile with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserProfile) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetSub()); l < 3 || l > 100 {
		return UserProfileValidationError{
			field:  "Sub",
			reason: "value length must be between 3 and 100 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 100 {
		return UserProfileValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
	}

	// no validation rules for UserName

	// no validation rules for Email

	// no validation rules for PhoneNumber

	// no validation rules for ExternalSource

	// no validation rules for ProfilePicUrl

	if v, ok := interface{}(m.GetTokenValidTill()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserProfileValidationError{
				field:  "TokenValidTill",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserProfileValidationError is the validation error returned by
// UserProfile.Validate if the designated constraints aren't met.
type UserProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserProfileValidationError) ErrorName() string { return "UserProfileValidationError" }

// Error satisfies the builtin error interface
func (e UserProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserProfileValidationError{}