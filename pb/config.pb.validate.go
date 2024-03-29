// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/config.proto

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _config_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetListener()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Listener",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLogger()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Logger",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("Routes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatch() == nil {
		return RouteValidationError{
			field:  "Match",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetResponseCode() <= 0 {
		return RouteValidationError{
			field:  "ResponseCode",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for ResponseHeaders

	// no validation rules for ResponseBody

	if d := m.GetDelay(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return RouteValidationError{
				field:  "Delay",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur < gte {
			return RouteValidationError{
				field:  "Delay",
				reason: "value must be greater than or equal to 0s",
			}
		}

	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on RouteMatch with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RouteMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteMatchValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.PathSpecifier.(type) {

	case *RouteMatch_Prefix:
		// no validation rules for Prefix

	case *RouteMatch_Path:
		// no validation rules for Path

	default:
		return RouteMatchValidationError{
			field:  "PathSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteMatchValidationError is the validation error returned by
// RouteMatch.Validate if the designated constraints aren't met.
type RouteMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteMatchValidationError) ErrorName() string { return "RouteMatchValidationError" }

// Error satisfies the builtin error interface
func (e RouteMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteMatchValidationError{}

// Validate checks the field values on HeaderMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HeaderMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return HeaderMatcherValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	if !_HeaderMatcher_Name_Pattern.MatchString(m.GetName()) {
		return HeaderMatcherValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
		}
	}

	switch m.HeaderMatchSpecifier.(type) {

	case *HeaderMatcher_ExactMatch:
		// no validation rules for ExactMatch

	}

	return nil
}

// HeaderMatcherValidationError is the validation error returned by
// HeaderMatcher.Validate if the designated constraints aren't met.
type HeaderMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderMatcherValidationError) ErrorName() string { return "HeaderMatcherValidationError" }

// Error satisfies the builtin error interface
func (e HeaderMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderMatcherValidationError{}

var _HeaderMatcher_Name_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

// Validate checks the field values on Logger with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Logger) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Level

	switch m.Format.(type) {

	case *Logger_Pretty:
		// no validation rules for Pretty

	}

	return nil
}

// LoggerValidationError is the validation error returned by Logger.Validate if
// the designated constraints aren't met.
type LoggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoggerValidationError) ErrorName() string { return "LoggerValidationError" }

// Error satisfies the builtin error interface
func (e LoggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoggerValidationError{}

// Validate checks the field values on TCPSocket with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TCPSocket) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetAddress()) < 1 {
		return TCPSocketValidationError{
			field:  "Address",
			reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetPort() > 65535 {
		return TCPSocketValidationError{
			field:  "Port",
			reason: "value must be less than or equal to 65535",
		}
	}

	// no validation rules for Secure

	return nil
}

// TCPSocketValidationError is the validation error returned by
// TCPSocket.Validate if the designated constraints aren't met.
type TCPSocketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TCPSocketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TCPSocketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TCPSocketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TCPSocketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TCPSocketValidationError) ErrorName() string { return "TCPSocketValidationError" }

// Error satisfies the builtin error interface
func (e TCPSocketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTCPSocket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TCPSocketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TCPSocketValidationError{}

// Validate checks the field values on UnixSocket with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UnixSocket) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPath()) < 1 {
		return UnixSocketValidationError{
			field:  "Path",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// UnixSocketValidationError is the validation error returned by
// UnixSocket.Validate if the designated constraints aren't met.
type UnixSocketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnixSocketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnixSocketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnixSocketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnixSocketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnixSocketValidationError) ErrorName() string { return "UnixSocketValidationError" }

// Error satisfies the builtin error interface
func (e UnixSocketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnixSocket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnixSocketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnixSocketValidationError{}

// Validate checks the field values on Listener with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Listener) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Socket.(type) {

	case *Listener_Tcp:

		if v, ok := interface{}(m.GetTcp()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "Tcp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Listener_Unix:

		if v, ok := interface{}(m.GetUnix()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  "Unix",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ListenerValidationError{
			field:  "Socket",
			reason: "value is required",
		}

	}

	return nil
}

// ListenerValidationError is the validation error returned by
// Listener.Validate if the designated constraints aren't met.
type ListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerValidationError) ErrorName() string { return "ListenerValidationError" }

// Error satisfies the builtin error interface
func (e ListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerValidationError{}
