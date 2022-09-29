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
		if len(formatAndArgs) > 0 {
			method(t)(formatAndArgs[0].(string), formatAndArgs[1:]...)
		} else {
			method(t)("got %v; want %v", got, want)
		}
	}
}

// NoErr asserts that err is nil. Optional formatAndArgs can be provided to
// customize the error message, the first element must be a string, otherwise
// NoErr panics.
func NoErr[T E | F](t T, err error, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if err != nil {
		if len(formatAndArgs) > 0 {
			method(t)(formatAndArgs[0].(string), formatAndArgs[1:]...)
		} else {
			method(t)("got %v; want no error", err)
		}
	}
}

// IsErr asserts that [errors.Is](err, target) is true. Optional formatAndArgs
// can be provided to customize the error message, the first element must be a
// string, otherwise IsErr panics.
func IsErr[T E | F](t T, err, target error, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if !errors.Is(err, target) {
		if len(formatAndArgs) > 0 {
			method(t)(formatAndArgs[0].(string), formatAndArgs[1:]...)
		} else {
			method(t)("got %v; want %v", err, target)
		}
	}
}

// AsErr asserts that [errors.As](err, target) is true. Optional formatAndArgs
// can be provided to customize the error message, the first element must be a
// string, otherwise AsErr panics.
func AsErr[T E | F](t T, err error, target any, formatAndArgs ...any) {
	(*testing.T)(t).Helper()
	if !errors.As(err, target) {
		if len(formatAndArgs) > 0 {
			method(t)(formatAndArgs[0].(string), formatAndArgs[1:]...)
		} else {
			method(t)("got %T; want %T", err, target)
		}
	}
}

// method returns the method to call based on t's type.
func method[T E | F](t T) func(format string, args ...any) {
	switch any(t).(type) {
	case E, *testing.T:
		return (*testing.T)(t).Errorf
	case F:
		return (*testing.T)(t).Fatalf
	default:
		panic("unreachable")
	}
}
