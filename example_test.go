package assert_test

import (
	"fmt"
	"os"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/dotimport"
)

func ExampleEqual() {
	assert.Equal[E](t, 1, 2)
	// Output: values are not equal
	// got:  1
	// want: 2
}

func ExampleNoErr() {
	assert.NoErr[E](t, err)
	// Output: unexpected error: file already exists
}

func ExampleIsErr() {
	assert.IsErr[E](t, err, os.ErrNotExist)
	// Output: errors.Is() mismatch
	// got:  file already exists
	// want: file does not exist
}

func ExampleAsErr() {
	assert.AsErr[E](t, err, new(*os.PathError))
	// Output: errors.As() mismatch
	// got:  *errors.errorString
	// want: *fs.PathError
}

func ExamplePanics() {
	assert.Panics[E](t, func() { /* panic? */ }, 42)
	// Output: the function didn't panic
}

var err = os.ErrExist

var t printer

type printer struct{}

func (printer) Helper()                           {}
func (printer) Errorf(format string, args ...any) { fmt.Printf(format, args...) }
func (printer) Fatalf(format string, args ...any) { fmt.Printf(format, args...) }
