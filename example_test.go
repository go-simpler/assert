package assert_test

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/go-simpler/assert"
	. "github.com/go-simpler/assert/dotimport"
)

var t *testing.T

func ExampleEqual() {
	assert.Equal[E](t, 1, 2) // prints "got 1; want 2"
	assert.Equal[F](t, 1, 2) // prints "got 1; want 2" and stops the test
}

func ExampleNoErr() {
	err := errors.New("test")
	assert.NoErr[E](t, err) // prints "got test; want no error"
	assert.NoErr[F](t, err) // prints "got test; want no error" and stops the test
}

func ExampleIsErr() {
	err := errors.New("test")
	assert.IsErr[E](t, err, fs.ErrNotExist) // prints "got test; want file does not exist"
	assert.IsErr[F](t, err, fs.ErrNotExist) // prints "got test; want file does not exist" and stops the test
}

func ExampleAsErr() {
	err := errors.New("test")
	assert.AsErr[E](t, err, new(*fs.PathError)) // prints "got *errors.errorString; want **fs.PathError"
	assert.AsErr[F](t, err, new(*fs.PathError)) // prints "got *errors.errorString; want **fs.PathError" and stops the test
}
