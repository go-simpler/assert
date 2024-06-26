// Package assert implements assertions for the standard [testing] package.
package assert

import (
	"errors"
	"reflect"
)

// TB is a tiny subset of [testing.TB].
type TB interface {
	Helper()
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

// Param controls the behavior of an assertion if it fails.
// Either [E] or [F] must be specified as the type parameter.
type Param interface {
	method(t TB) func(format string, args ...any)
}

// E is a [Param] that marks the test as failed but continues execution (similar to [testing.T.Errorf]).
type E struct{}

func (E) method(t TB) func(format string, args ...any) { return t.Errorf }

// F is a [Param] that marks the test as failed and stops execution (similar to [testing.T.Fatalf]).
type F struct{}

func (F) method(t TB) func(format string, args ...any) { return t.Fatalf }

// Equal asserts that two values are equal.
func Equal[T Param, V any](t TB, got, want V, formatAndArgs ...any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		fail[T](t, formatAndArgs, "values are not equal\ngot:  %v\nwant: %v", got, want)
	}
}

// NoErr asserts that the error is nil.
func NoErr[T Param](t TB, err error, formatAndArgs ...any) {
	t.Helper()
	if err != nil {
		fail[T](t, formatAndArgs, "unexpected error: %v", err)
	}
}

// IsErr asserts that [errors.Is](err, target) is true.
func IsErr[T Param](t TB, err, target error, formatAndArgs ...any) {
	t.Helper()
	if !errors.Is(err, target) {
		fail[T](t, formatAndArgs, "errors.Is() mismatch\ngot:  %v\nwant: %v", err, target)
	}
}

// AsErr asserts that [errors.As](err, target) is true.
func AsErr[T Param](t TB, err error, target any, formatAndArgs ...any) {
	t.Helper()
	if !errors.As(err, target) {
		typ := reflect.TypeOf(target).Elem() // dereference the pointer to get the real type.
		fail[T](t, formatAndArgs, "errors.As() mismatch\ngot:  %T\nwant: %s", err, typ)
	}
}

// Panics asserts that the given function panics with the argument v.
// If v is nil, the panic argument is ignored.
func Panics[T Param](t TB, fn func(), v any, formatAndArgs ...any) {
	t.Helper()
	defer func() {
		t.Helper()
		switch r := recover(); {
		case r == nil:
			fail[T](t, formatAndArgs, "the function didn't panic")
		case v != nil && !reflect.DeepEqual(r, v):
			fail[T](t, nil, "panic argument mismatch\ngot:  %v\nwant: %v", r, v)
		}
	}()
	fn()
}

func fail[T Param](t TB, customFormatAndArgs []any, format string, args ...any) {
	t.Helper()
	if len(customFormatAndArgs) > 0 {
		format = customFormatAndArgs[0].(string)
		args = customFormatAndArgs[1:]
	}
	(*new(T)).method(t)(format, args...)
}
