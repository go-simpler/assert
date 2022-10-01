package assert_test

import (
	"reflect"
	"testing"

	"github.com/junk1tm/assert"
	. "github.com/junk1tm/assert/dotimport"
)

// assertCall holds the information about calling an assertion.
type assertCall struct {
	helperCalls  int
	errorfCalled bool
	fatalfCalled bool
	format       string
	args         []any
}

// okCall returns an [assertCall] that should happen if the assertion was
// successful.
func okCall() assertCall {
	return assertCall{
		helperCalls: 1, // at least one t.Helper() call is always expected
	}
}

// errorfCall returns an [assertCall] that should happen if the assertion failed
// and [testing.T.Errrof] was called.
func errorfCall(format string, args ...any) assertCall {
	return assertCall{
		helperCalls:  2,
		errorfCalled: true,
		format:       format,
		args:         args,
	}
}

// fatalfCall returns an [assertCall] that should happen if the assertion failed
// and [testing.T.Fatalf] was called.
func fatalfCall(format string, args ...any) assertCall {
	return assertCall{
		helperCalls:  2,
		fatalfCalled: true,
		format:       format,
		args:         args,
	}
}

// testAssertCall ensures that the [assertCall] matches expectations.
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
		t.Errorf("format: got %q want %q", got.format, want.format)
	}
	if !reflect.DeepEqual(got.args, want.args) {
		t.Errorf("args: got %v want %v", got.args, want.args)
	}
}

// spyTB is an [assert.TB] implementation that records an [assertCall].
type spyTB struct{ assertCall }

func (tb *spyTB) Helper() { tb.helperCalls++ }

func (tb *spyTB) Errorf(format string, args ...any) {
	tb.errorfCalled = true
	tb.format, tb.args = format, args
}

func (tb *spyTB) Fatalf(format string, args ...any) {
	tb.fatalfCalled = true
	tb.format, tb.args = format, args
}

func TestEqual(t *testing.T) {
	equalE := func(tb *spyTB, got, want int, formatAndArgs []any) {
		assert.Equal[E](tb, got, want, formatAndArgs...)
	}
	equalF := func(tb *spyTB, got, want int, formatAndArgs []any) {
		assert.Equal[F](tb, got, want, formatAndArgs...)
	}

	tests := map[string]struct {
		fn            func(tb *spyTB, got, want int, formatAndArgs []any)
		a, b          int
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn: equalE,
			a:  1, b: 1,
			want: okCall(),
		},
		"fail [E]": {
			fn: equalE,
			a:  1, b: 2,
			want: errorfCall("got %v; want %v", 1, 2),
		},
		"fail [F]": {
			fn: equalF,
			a:  1, b: 2,
			want: fatalfCall("got %v; want %v", 1, 2),
		},
		"fail [F] (custom message)": {
			fn: equalF,
			a:  1, b: 2,
			formatAndArgs: []any{"actual %d, expected %d", 1, 2},
			want:          fatalfCall("actual %d, expected %d", 1, 2),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var spy spyTB
			tt.fn(&spy, tt.a, tt.b, tt.formatAndArgs)
			testAssertCall(t, spy.assertCall, tt.want)
		})
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

func TestNoErr(t *testing.T) {
	noerrE := func(tb *spyTB, err error, formatAndArgs []any) {
		assert.NoErr[E](tb, err, formatAndArgs...)
	}
	noerrF := func(tb *spyTB, err error, formatAndArgs []any) {
		assert.NoErr[F](tb, err, formatAndArgs...)
	}

	tests := map[string]struct {
		fn            func(tb *spyTB, err error, formatAndArgs []any)
		err           error
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn:   noerrE,
			err:  nil,
			want: okCall(),
		},
		"fail [E]": {
			fn:   noerrE,
			err:  errFoo,
			want: errorfCall("got %v; want no error", errFoo),
		},
		"fail [F]": {
			fn:   noerrF,
			err:  errFoo,
			want: fatalfCall("got %v; want no error", errFoo),
		},
		"fail [F] (custom message)": {
			fn:            noerrF,
			err:           errFoo,
			formatAndArgs: []any{"actual %v, expected no error", errFoo},
			want:          fatalfCall("actual %v, expected no error", errFoo),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var spy spyTB
			tt.fn(&spy, tt.err, tt.formatAndArgs)
			testAssertCall(t, spy.assertCall, tt.want)
		})
	}
}

func TestIsErr(t *testing.T) {
	iserrE := func(tb *spyTB, err, target error, formatAndArgs []any) {
		assert.IsErr[E](tb, err, target, formatAndArgs...)
	}
	iserrF := func(tb *spyTB, err, target error, formatAndArgs []any) {
		assert.IsErr[F](tb, err, target, formatAndArgs...)
	}

	tests := map[string]struct {
		fn            func(tb *spyTB, err, target error, formatAndArgs []any)
		err           error
		target        error
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn:     iserrE,
			err:    errFoo,
			target: errFoo,
			want:   okCall(),
		},
		"fail [E]": {
			fn:     iserrE,
			err:    errFoo,
			target: errBar,
			want:   errorfCall("got %v; want %v", errFoo, errBar),
		},
		"fail [F]": {
			fn:     iserrF,
			err:    errFoo,
			target: errBar,
			want:   fatalfCall("got %v; want %v", errFoo, errBar),
		},
		"fail [F] (custom message)": {
			fn:            iserrF,
			err:           errFoo,
			target:        errBar,
			formatAndArgs: []any{"actual %v, expected %v", errFoo, errBar},
			want:          fatalfCall("actual %v, expected %v", errFoo, errBar),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var spy spyTB
			tt.fn(&spy, tt.err, tt.target, tt.formatAndArgs)
			testAssertCall(t, spy.assertCall, tt.want)
		})
	}
}

func TestAsErr(t *testing.T) {
	aserrE := func(tb *spyTB, err error, target any, formatAndArgs []any) {
		assert.AsErr[E](tb, err, target, formatAndArgs...)
	}
	aserrF := func(tb *spyTB, err error, target any, formatAndArgs []any) {
		assert.AsErr[F](tb, err, target, formatAndArgs...)
	}

	tests := map[string]struct {
		fn            func(tb *spyTB, err error, target any, formatAndArgs []any)
		err           error
		target        any
		formatAndArgs []any
		want          assertCall
	}{
		"ok [E]": {
			fn:     aserrE,
			err:    errFoo,
			target: new(fooError),
			want:   okCall(),
		},
		"fail [E]": {
			fn:     aserrE,
			err:    errFoo,
			target: new(barError),
			want:   errorfCall("got %T; want %T", errFoo, new(barError)),
		},
		"fail [F]": {
			fn:     aserrF,
			err:    errFoo,
			target: new(barError),
			want:   fatalfCall("got %T; want %T", errFoo, new(barError)),
		},
		"fail [F] (custom message)": {
			fn:            aserrF,
			err:           errFoo,
			target:        new(barError),
			formatAndArgs: []any{"actual %T, expected %T", errFoo, new(barError)},
			want:          fatalfCall("actual %T, expected %T", errFoo, new(barError)),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var spy spyTB
			tt.fn(&spy, tt.err, tt.target, tt.formatAndArgs)
			testAssertCall(t, spy.assertCall, tt.want)
		})
	}
}
