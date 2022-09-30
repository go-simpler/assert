// Package assert provides common assertions to use with the standard
// [testing] package.
package assert

import (
	"errors"
	"reflect"
)

// tb is a tiny subset of [testing.TB] used by [assert].
type tb interface {
	Helper()
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

// parameter is type parameter that control the behaviour of an assertion in
// case it fails. Either [E] or [F] should be specified when calling the
// assertion.
type parameter interface {
	// method returns t's method to call, either [tb.Errorf] or [tb.Fatalf].
	method(t tb) func(format string, args ...any)
}

// E is a [parameter] that marks the test as having failed but continues its
// execution (similar to [testing.T.Errorf]).
type E struct{}

func (E) method(t tb) func(format string, args ...any) { return t.Errorf }

// F is a [parameter] that marks the test as having failed and stops its
// execution (similar to [testing.T.Fatalf]).
type F struct{}

func (F) method(t tb) func(format string, args ...any) { return t.Fatalf }

// Equal asserts that got and want are equal. Optional formatAndArgs can be
// provided to customize the error message, the first element must be a string,
// otherwise Equal panics.
func Equal[T parameter, V any](t tb, got, want V, formatAndArgs ...any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		fail[T](t, formatAndArgs, "got %v; want %v", got, want)
	}
}

// NoErr asserts that err is nil. Optional formatAndArgs can be provided to
// customize the error message, the first element must be a string, otherwise
// NoErr panics.
func NoErr[T parameter](t tb, err error, formatAndArgs ...any) {
	t.Helper()
	if err != nil {
		fail[T](t, formatAndArgs, "got %v; want no error", err)
	}
}

// IsErr asserts that [errors.Is](err, target) is true. Optional formatAndArgs
// can be provided to customize the error message, the first element must be a
// string, otherwise IsErr panics.
func IsErr[T parameter](t tb, err, target error, formatAndArgs ...any) {
	t.Helper()
	if !errors.Is(err, target) {
		fail[T](t, formatAndArgs, "got %v; want %v", err, target)
	}
}

// AsErr asserts that [errors.As](err, target) is true. Optional formatAndArgs
// can be provided to customize the error message, the first element must be a
// string, otherwise AsErr panics.
func AsErr[T parameter](t tb, err error, target any, formatAndArgs ...any) {
	t.Helper()
	if !errors.As(err, target) {
		fail[T](t, formatAndArgs, "got %T; want %T", err, target)
	}
}

// fail marks the test as having failed and continues/stops its execution based
// on T's type.
func fail[T parameter](t tb, customFormatAndArgs []any, format string, args ...any) {
	t.Helper()
	if len(customFormatAndArgs) > 0 {
		format = customFormatAndArgs[0].(string)
		args = customFormatAndArgs[1:]
	}
	(*new(T)).method(t)(format, args...)
}
