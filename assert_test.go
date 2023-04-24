package assert_test

import (
	"reflect"
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
			want: errorfCall("\ngot\t%v\nwant\t%v", 1, 2),
		},
		"fail [F]": {
			fn: assert.Equal[F, int],
			a:  1, b: 2,
			want: fatalfCall("\ngot\t%v\nwant\t%v", 1, 2),
		},
		"fail [F] (custom message)": {
			fn: assert.Equal[F, int],
			a:  1, b: 2,
			formatAndArgs: []any{"\nactual\t%v\nexpected\t%v", 1, 2},
			want:          fatalfCall("\nactual\t%v\nexpected\t%v", 1, 2),
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
			want: errorfCall("\ngot\t%v\nwant\tno error", errFoo),
		},
		"fail [F]": {
			fn:   assert.NoErr[F],
			err:  errFoo,
			want: fatalfCall("\ngot\t%v\nwant\tno error", errFoo),
		},
		"fail [F] (custom message)": {
			fn:            assert.NoErr[F],
			err:           errFoo,
			formatAndArgs: []any{"\nactual\t%v\nexpected\tno error", errFoo},
			want:          fatalfCall("\nactual\t%v\nexpected\tno error", errFoo),
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
			want:   errorfCall("\ngot\t%v\nwant\t%v", errFoo, errBar),
		},
		"fail [F]": {
			fn:     assert.IsErr[F],
			err:    errFoo,
			target: errBar,
			want:   fatalfCall("\ngot\t%v\nwant\t%v", errFoo, errBar),
		},
		"fail [F] (custom message)": {
			fn:            assert.IsErr[F],
			err:           errFoo,
			target:        errBar,
			formatAndArgs: []any{"\nactual\t%v\nexpected\t%v", errFoo, errBar},
			want:          fatalfCall("\nactual\t%v\nexpected\t%v", errFoo, errBar),
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
			want:   errorfCall("\ngot\t%T\nwant\t%T", errFoo, new(barError)),
		},
		"fail [F]": {
			fn:     assert.AsErr[F],
			err:    errFoo,
			target: new(barError),
			want:   fatalfCall("\ngot\t%T\nwant\t%T", errFoo, new(barError)),
		},
		"fail [F] (custom message)": {
			fn:            assert.AsErr[F],
			err:           errFoo,
			target:        new(barError),
			formatAndArgs: []any{"\nactual\t%T\nexpected\t%T", errFoo, new(barError)},
			want:          fatalfCall("\nactual\t%T\nexpected\t%T", errFoo, new(barError)),
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
	format       string
	args         []any
}

func (ac *assertCall) Helper() { ac.helperCalls++ }

func (ac *assertCall) Errorf(format string, args ...any) {
	ac.errorfCalled = true
	ac.format, ac.args = format, args
}

func (ac *assertCall) Fatalf(format string, args ...any) {
	ac.fatalfCalled = true
	ac.format, ac.args = format, args
}

func okCall() assertCall {
	return assertCall{
		helperCalls: 1, // at least one t.Helper() call is always expected.
	}
}

func errorfCall(format string, args ...any) assertCall {
	return assertCall{
		helperCalls:  2,
		errorfCalled: true,
		format:       format,
		args:         args,
	}
}

func fatalfCall(format string, args ...any) assertCall {
	return assertCall{
		helperCalls:  2,
		fatalfCalled: true,
		format:       format,
		args:         args,
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
	if got.format != want.format {
		t.Errorf("invalid format: got %q want %q", got.format, want.format)
	}
	if !reflect.DeepEqual(got.args, want.args) {
		t.Errorf("invalid args: got %v want %v", got.args, want.args)
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
