package assert_test

import (
	"fmt"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/dotimport"
)

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		fn            func(t assert.TB, got, want int, formatAndArgs ...any)
		a, b          int
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn: assert.Equal[E, int],
			a:  1, b: 1,
			want: okCall(),
		},
		"fail [E]": {
			fn: assert.Equal[E, int],
			a:  1, b: 2,
			want: errorfCall("values are not equal\ngot:  1\nwant: 2"),
		},
		"fail [F]": {
			fn: assert.Equal[F, int],
			a:  1, b: 2,
			want: fatalfCall("values are not equal\ngot:  1\nwant: 2"),
		},
		"fail [F] (custom message)": {
			fn: assert.Equal[F, int],
			a:  1, b: 2,
			formatAndArgs: []any{"%d != %d", 1, 2},
			want:          fatalfCall("1 != 2"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var got assertCall
			tt.fn(&got, tt.a, tt.b, tt.formatAndArgs...)
			testAssertCall(t, got, tt.want)
		})
	}
}

func TestNoErr(t *testing.T) {
	tests := map[string]struct {
		fn            func(t assert.TB, err error, formatAndArgs ...any)
		err           error
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn:   assert.NoErr[E],
			err:  nil,
			want: okCall(),
		},
		"fail [E]": {
			fn:   assert.NoErr[E],
			err:  errFoo,
			want: errorfCall("unexpected error: foo"),
		},
		"fail [F]": {
			fn:   assert.NoErr[F],
			err:  errFoo,
			want: fatalfCall("unexpected error: foo"),
		},
		"fail [F] (custom message)": {
			fn:            assert.NoErr[F],
			err:           errFoo,
			formatAndArgs: []any{"%v != nil", errFoo},
			want:          fatalfCall("foo != nil"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var got assertCall
			tt.fn(&got, tt.err, tt.formatAndArgs...)
			testAssertCall(t, got, tt.want)
		})
	}
}

func TestIsErr(t *testing.T) {
	tests := map[string]struct {
		fn            func(t assert.TB, err, target error, formatAndArgs ...any)
		err           error
		target        error
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn:     assert.IsErr[E],
			err:    errFoo,
			target: errFoo,
			want:   okCall(),
		},
		"fail [E]": {
			fn:     assert.IsErr[E],
			err:    errFoo,
			target: errBar,
			want:   errorfCall("errors.Is == false\ngot:  foo\nwant: bar"),
		},
		"fail [F]": {
			fn:     assert.IsErr[F],
			err:    errFoo,
			target: errBar,
			want:   fatalfCall("errors.Is == false\ngot:  foo\nwant: bar"),
		},
		"fail [F] (custom message)": {
			fn:            assert.IsErr[F],
			err:           errFoo,
			target:        errBar,
			formatAndArgs: []any{"%v != %v", errFoo, errBar},
			want:          fatalfCall("foo != bar"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var got assertCall
			tt.fn(&got, tt.err, tt.target, tt.formatAndArgs...)
			testAssertCall(t, got, tt.want)
		})
	}
}

func TestAsErr(t *testing.T) {
	tests := map[string]struct {
		fn            func(t assert.TB, err error, target any, formatAndArgs ...any)
		err           error
		target        any
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn:     assert.AsErr[E],
			err:    errFoo,
			target: new(fooError),
			want:   okCall(),
		},
		"fail [E]": {
			fn:     assert.AsErr[E],
			err:    errFoo,
			target: new(barError),
			want:   errorfCall("errors.As == false\ngot:  assert_test.fooError\nwant: assert_test.barError"),
		},
		"fail [F]": {
			fn:     assert.AsErr[F],
			err:    errFoo,
			target: new(barError),
			want:   fatalfCall("errors.As == false\ngot:  assert_test.fooError\nwant: assert_test.barError"),
		},
		"fail [F] (custom message)": {
			fn:            assert.AsErr[F],
			err:           errFoo,
			target:        new(barError),
			formatAndArgs: []any{"%T != %T", errFoo, errBar},
			want:          fatalfCall("assert_test.fooError != assert_test.barError"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var got assertCall
			tt.fn(&got, tt.err, tt.target, tt.formatAndArgs...)
			testAssertCall(t, got, tt.want)
		})
	}
}

type assertCall struct {
	helperCalls  int
	errorfCalled bool
	fatalfCalled bool
	message      string
}

func (ac *assertCall) Helper() { ac.helperCalls++ }

func (ac *assertCall) Errorf(format string, args ...any) {
	ac.errorfCalled = true
	ac.message = fmt.Sprintf(format, args...)
}

func (ac *assertCall) Fatalf(format string, args ...any) {
	ac.fatalfCalled = true
	ac.message = fmt.Sprintf(format, args...)
}

func okCall() assertCall {
	return assertCall{
		helperCalls: 1, // at least one t.Helper() call is always expected.
	}
}

func errorfCall(message string) assertCall {
	return assertCall{
		helperCalls:  2,
		errorfCalled: true,
		message:      message,
	}
}

func fatalfCall(message string) assertCall {
	return assertCall{
		helperCalls:  2,
		fatalfCalled: true,
		message:      message,
	}
}

func testAssertCall(t *testing.T, got, want assertCall) {
	t.Helper()
	if got.helperCalls != want.helperCalls {
		t.Errorf("t.Helper() calls: got %d want %d", got.helperCalls, want.helperCalls)
	}
	if got.errorfCalled != want.errorfCalled {
		t.Errorf("t.Errorf() called: got %t want %t", got.errorfCalled, want.errorfCalled)
	}
	if got.fatalfCalled != want.fatalfCalled {
		t.Errorf("t.Fatalf() called: got %t want %t", got.fatalfCalled, want.fatalfCalled)
	}
	if got.message != want.message {
		t.Errorf("unexpected message\ngot:  %q\nwant: %q", got.message, want.message)
	}
}

type fooError struct{}

func (fooError) Error() string { return "foo" }

type barError struct{}

func (barError) Error() string { return "bar" }

var (
	errFoo fooError
	errBar barError
)
