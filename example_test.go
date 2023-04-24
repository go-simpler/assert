package assert_test

import (
	"fmt"
	"os"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/dotimport"
)

func ExampleEqual() {
	assert.Equal[E](t, 1, 2)
	// Output:
	// got	1
	// want	2
}

func ExampleNoErr() {
	assert.NoErr[E](t, os.ErrExist)
	// Output:
	// got	file already exists
	// want	no error
}

func ExampleIsErr() {
	assert.IsErr[E](t, os.ErrExist, os.ErrNotExist)
	// Output:
	// got	file already exists
	// want	file does not exist
}

func ExampleAsErr() {
	assert.AsErr[E](t, os.ErrExist, new(*os.PathError))
	// Output:
	// got	*errors.errorString
	// want	**fs.PathError
}

var t printer

type printer struct{}

func (printer) Helper()                           {}
func (printer) Errorf(format string, args ...any) { fmt.Printf(format, args...) }
func (printer) Fatalf(format string, args ...any) { fmt.Printf(format, args...) }
