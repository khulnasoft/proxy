// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/thrift_to_metadata/v3/thrift_to_metadata.proto

package thrift_to_metadatav3

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

	v3 "github.com/khulnasoft/proxy/go/envoy/extensions/filters/network/thrift_proxy/v3"
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

	_ = v3.ProtocolType(0)
)

// Validate checks the field values on KeyValuePair with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KeyValuePair) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KeyValuePair with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KeyValuePairMultiError, or
// nil if none found.
func (m *KeyValuePair) ValidateAll() error {
	return m.validate(true)
}

func (m *KeyValuePair) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MetadataNamespace

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := KeyValuePairValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, KeyValuePairValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, KeyValuePairValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeyValuePairValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return KeyValuePairMultiError(errors)
	}

	return nil
}

// KeyValuePairMultiError is an error wrapping multiple validation errors
// returned by KeyValuePair.ValidateAll() if the designated constraints aren't met.
type KeyValuePairMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KeyValuePairMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KeyValuePairMultiError) AllErrors() []error { return m }

// KeyValuePairValidationError is the validation error returned by
// KeyValuePair.Validate if the designated constraints aren't met.
type KeyValuePairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyValuePairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyValuePairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyValuePairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyValuePairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyValuePairValidationError) ErrorName() string { return "KeyValuePairValidationError" }

// Error satisfies the builtin error interface
func (e KeyValuePairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyValuePair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyValuePairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyValuePairValidationError{}

// Validate checks the field values on FieldSelector with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FieldSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FieldSelector with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FieldSelectorMultiError, or
// nil if none found.
func (m *FieldSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *FieldSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := FieldSelectorValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetId(); val < -32768 || val > 32767 {
		err := FieldSelectorValidationError{
			field:  "Id",
			reason: "value must be inside range [-32768, 32767]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetChild()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FieldSelectorValidationError{
					field:  "Child",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FieldSelectorValidationError{
					field:  "Child",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChild()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FieldSelectorValidationError{
				field:  "Child",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FieldSelectorMultiError(errors)
	}

	return nil
}

// FieldSelectorMultiError is an error wrapping multiple validation errors
// returned by FieldSelector.ValidateAll() if the designated constraints
// aren't met.
type FieldSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldSelectorMultiError) AllErrors() []error { return m }

// FieldSelectorValidationError is the validation error returned by
// FieldSelector.Validate if the designated constraints aren't met.
type FieldSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldSelectorValidationError) ErrorName() string { return "FieldSelectorValidationError" }

// Error satisfies the builtin error interface
func (e FieldSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldSelectorValidationError{}

// Validate checks the field values on Rule with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Rule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Rule with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RuleMultiError, or nil if none found.
func (m *Rule) ValidateAll() error {
	return m.validate(true)
}

func (m *Rule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	if all {
		switch v := interface{}(m.GetFieldSelector()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RuleValidationError{
					field:  "FieldSelector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RuleValidationError{
					field:  "FieldSelector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFieldSelector()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuleValidationError{
				field:  "FieldSelector",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MethodName

	if all {
		switch v := interface{}(m.GetOnPresent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RuleValidationError{
					field:  "OnPresent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RuleValidationError{
					field:  "OnPresent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnPresent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuleValidationError{
				field:  "OnPresent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOnMissing()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RuleValidationError{
					field:  "OnMissing",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RuleValidationError{
					field:  "OnMissing",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnMissing()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuleValidationError{
				field:  "OnMissing",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RuleMultiError(errors)
	}

	return nil
}

// RuleMultiError is an error wrapping multiple validation errors returned by
// Rule.ValidateAll() if the designated constraints aren't met.
type RuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RuleMultiError) AllErrors() []error { return m }

// RuleValidationError is the validation error returned by Rule.Validate if the
// designated constraints aren't met.
type RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuleValidationError) ErrorName() string { return "RuleValidationError" }

// Error satisfies the builtin error interface
func (e RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuleValidationError{}

// Validate checks the field values on ThriftToMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ThriftToMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ThriftToMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ThriftToMetadataMultiError, or nil if none found.
func (m *ThriftToMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *ThriftToMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRequestRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftToMetadataValidationError{
						field:  fmt.Sprintf("RequestRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftToMetadataValidationError{
						field:  fmt.Sprintf("RequestRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftToMetadataValidationError{
					field:  fmt.Sprintf("RequestRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResponseRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftToMetadataValidationError{
						field:  fmt.Sprintf("ResponseRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftToMetadataValidationError{
						field:  fmt.Sprintf("ResponseRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftToMetadataValidationError{
					field:  fmt.Sprintf("ResponseRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if _, ok := v3.TransportType_name[int32(m.GetTransport())]; !ok {
		err := ThriftToMetadataValidationError{
			field:  "Transport",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v3.ProtocolType_name[int32(m.GetProtocol())]; !ok {
		err := ThriftToMetadataValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetAllowContentTypes() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 1 {
			err := ThriftToMetadataValidationError{
				field:  fmt.Sprintf("AllowContentTypes[%v]", idx),
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for AllowEmptyContentType

	if len(errors) > 0 {
		return ThriftToMetadataMultiError(errors)
	}

	return nil
}

// ThriftToMetadataMultiError is an error wrapping multiple validation errors
// returned by ThriftToMetadata.ValidateAll() if the designated constraints
// aren't met.
type ThriftToMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThriftToMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThriftToMetadataMultiError) AllErrors() []error { return m }

// ThriftToMetadataValidationError is the validation error returned by
// ThriftToMetadata.Validate if the designated constraints aren't met.
type ThriftToMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftToMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftToMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftToMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftToMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftToMetadataValidationError) ErrorName() string { return "ThriftToMetadataValidationError" }

// Error satisfies the builtin error interface
func (e ThriftToMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftToMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftToMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftToMetadataValidationError{}

// Validate checks the field values on ThriftToMetadataPerRoute with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ThriftToMetadataPerRoute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ThriftToMetadataPerRoute with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ThriftToMetadataPerRouteMultiError, or nil if none found.
func (m *ThriftToMetadataPerRoute) ValidateAll() error {
	return m.validate(true)
}

func (m *ThriftToMetadataPerRoute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRequestRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftToMetadataPerRouteValidationError{
						field:  fmt.Sprintf("RequestRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftToMetadataPerRouteValidationError{
						field:  fmt.Sprintf("RequestRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftToMetadataPerRouteValidationError{
					field:  fmt.Sprintf("RequestRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResponseRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftToMetadataPerRouteValidationError{
						field:  fmt.Sprintf("ResponseRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftToMetadataPerRouteValidationError{
						field:  fmt.Sprintf("ResponseRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftToMetadataPerRouteValidationError{
					field:  fmt.Sprintf("ResponseRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ThriftToMetadataPerRouteMultiError(errors)
	}

	return nil
}

// ThriftToMetadataPerRouteMultiError is an error wrapping multiple validation
// errors returned by ThriftToMetadataPerRoute.ValidateAll() if the designated
// constraints aren't met.
type ThriftToMetadataPerRouteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThriftToMetadataPerRouteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThriftToMetadataPerRouteMultiError) AllErrors() []error { return m }

// ThriftToMetadataPerRouteValidationError is the validation error returned by
// ThriftToMetadataPerRoute.Validate if the designated constraints aren't met.
type ThriftToMetadataPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftToMetadataPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftToMetadataPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftToMetadataPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftToMetadataPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftToMetadataPerRouteValidationError) ErrorName() string {
	return "ThriftToMetadataPerRouteValidationError"
}

// Error satisfies the builtin error interface
func (e ThriftToMetadataPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftToMetadataPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftToMetadataPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftToMetadataPerRouteValidationError{}
