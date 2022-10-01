package assert_test

import (
	"errors"
	"os"
	"testing"

	"github.com/junk1tm/assert"
	. "github.com/junk1tm/assert/dotimport"
)

var t *testing.T

func ExampleEqual() {
	// testing only:
	if 1 != 2 {
		t.Errorf("got %v; want %v", 1, 2)
	}
	// with assert:
	assert.Equal[E](t, 1, 2)

	// testing only:
	if 1 != 2 {
		t.Fatalf("got %v; want %v", 1, 2)
	}
	// with assert:
	assert.Equal[F](t, 1, 2)
}

var err error

func ExampleNoErr() {
	// testing only:
	if err != nil {
		t.Errorf("got %v; want no error", err)
	}
	// with assert:
	assert.NoErr[E](t, err)

	// testing only:
	if err != nil {
		t.Fatalf("got %v; want no error", err)
	}
	// with assert:
	assert.NoErr[F](t, err)
}

func ExampleIsErr() {
	// testing only:
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("got %v; want %v", err, os.ErrNotExist)
	}
	// with assert:
	assert.IsErr[E](t, err, os.ErrNotExist)

	// testing only:
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("got %v; want %v", err, os.ErrNotExist)
	}
	// with assert:
	assert.IsErr[F](t, err, os.ErrNotExist)
}

func ExampleAsErr() {
	// testing only:
	if !errors.As(err, new(*os.PathError)) {
		t.Errorf("got %T; want %T", err, new(*os.PathError))
	}
	// with assert:
	assert.AsErr[E](t, err, new(*os.PathError))

	// testing only:
	if !errors.As(err, new(*os.PathError)) {
		t.Fatalf("got %T; want %T", err, new(*os.PathError))
	}
	// with assert:
	assert.AsErr[F](t, err, new(*os.PathError))
}
