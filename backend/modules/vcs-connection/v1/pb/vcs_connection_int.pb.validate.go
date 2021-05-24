// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/vcs_connection_int.proto

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

// Validate checks the field values on VCSConnection with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *VCSConnection) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Label

	// no validation rules for Provider

	// no validation rules for ConnectionId

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if v, ok := interface{}(m.GetAccessTokenExpiry()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VCSConnectionValidationError{
				field:  "AccessTokenExpiry",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRefreshTokenExpiry()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VCSConnectionValidationError{
				field:  "RefreshTokenExpiry",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Revoked

	// no validation rules for AccountId

	// no validation rules for UserName

	return nil
}

// VCSConnectionValidationError is the validation error returned by
// VCSConnection.Validate if the designated constraints aren't met.
type VCSConnectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VCSConnectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VCSConnectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VCSConnectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VCSConnectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VCSConnectionValidationError) ErrorName() string { return "VCSConnectionValidationError" }

// Error satisfies the builtin error interface
func (e VCSConnectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVCSConnection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VCSConnectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VCSConnectionValidationError{}

// Validate checks the field values on CreateVCSConnectionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateVCSConnectionRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetVcsConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateVCSConnectionRequestValidationError{
				field:  "VcsConnection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateVCSConnectionRequestValidationError is the validation error returned
// by CreateVCSConnectionRequest.Validate if the designated constraints aren't met.
type CreateVCSConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateVCSConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateVCSConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateVCSConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateVCSConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateVCSConnectionRequestValidationError) ErrorName() string {
	return "CreateVCSConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateVCSConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateVCSConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateVCSConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateVCSConnectionRequestValidationError{}