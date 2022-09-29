// Package assert provides common assertions to use with the standard
// [testing] package.
package assert

import (
	"errors"
	"reflect"
	"testing"
)

// These types control the behaviour of an assertion in case it fails. Either
// [E] or [F] should be specified as a type parameter.
type (
	// E marks the test as having failed but continues its execution (similar to
	// [testing.T.Errorf]).
	E *testing.T
	// F marks the test as having failed and stops its execution (similar to
	// [testing.T.Fatalf]).
	F *testing.T
)

// Equal asserts that got and want are equal. Optional formatAndArgs can be
// provided to customize the error message, the first element must be a string,
// otherwise Equal panics.
func Equal[T E | F, V any](t T, got, want V, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if !reflect.DeepEqual(got, want) {
		fail(t, formatAndArgs, "got %v; want %v", got, want)
	}
}

// NoErr asserts that err is nil. Optional formatAndArgs can be provided to
// customize the error message, the first element must be a string, otherwise
// NoErr panics.
func NoErr[T E | F](t T, err error, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if err != nil {
		fail(t, formatAndArgs, "got %v; want no error", err)
	}
}

// IsErr asserts that [errors.Is](err, target) is true. Optional formatAndArgs
// can be provided to customize the error message, the first element must be a
// string, otherwise IsErr panics.
func IsErr[T E | F](t T, err, target error, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if !errors.Is(err, target) {
		fail(t, formatAndArgs, "got %v; want %v", err, target)
	}
}

// AsErr asserts that [errors.As](err, target) is true. Optional formatAndArgs
// can be provided to customize the error message, the first element must be a
// string, otherwise AsErr panics.
func AsErr[T E | F](t T, err error, target any, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if !errors.As(err, target) {
		fail(t, formatAndArgs, "got %T; want %T", err, target)
	}
}

// fail calls either [testing.T.Errorf] or [testing.T.Fatalf] based on t's type.
func fail[T E | F](t T, customFormatAndArgs []any, format string, args ...any) {
	(*testing.T)(t).Helper()
	if len(customFormatAndArgs) > 0 {
		format = customFormatAndArgs[0].(string)
		args = customFormatAndArgs[1:]
	}
	switch any(t).(type) {
	case E:
		(*testing.T)(t).Errorf(format, args...)
	case F:
		(*testing.T)(t).Fatalf(format, args...)
	default:
		panic("unreachable")
	}
}
