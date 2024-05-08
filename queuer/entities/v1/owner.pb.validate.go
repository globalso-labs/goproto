// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/entities/v1/owner.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _owner_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Owner with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Owner) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Owner with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OwnerMultiError, or nil if none found.
func (m *Owner) ValidateAll() error {
	return m.validate(true)
}

func (m *Owner) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = OwnerValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 100 {
		err := OwnerValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Consumers != nil {
		// no validation rules for Consumers
	}

	if m.Streams != nil {
		// no validation rules for Streams
	}

	if len(errors) > 0 {
		return OwnerMultiError(errors)
	}

	return nil
}

func (m *Owner) _validateUuid(uuid string) error {
	if matched := _owner_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OwnerMultiError is an error wrapping multiple validation errors returned by
// Owner.ValidateAll() if the designated constraints aren't met.
type OwnerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OwnerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OwnerMultiError) AllErrors() []error { return m }

// OwnerValidationError is the validation error returned by Owner.Validate if
// the designated constraints aren't met.
type OwnerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OwnerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OwnerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OwnerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OwnerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OwnerValidationError) ErrorName() string { return "OwnerValidationError" }

// Error satisfies the builtin error interface
func (e OwnerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOwner.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OwnerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OwnerValidationError{}
