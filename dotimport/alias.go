// Package dotimport provides type aliases for the parent [assert] package. It
// is intended to be imported using dot syntax so that [E] and [F] can be used
// as if they were local types.
//
//	package foo_test
//
//	import (
//		"testing"
//
//		"assert"
//		. "assert/dotimport"
//	)
//
//	func TestFoo(t *testing.T) {
//		assert.NoErr[E](t, foo.Foo())
//	}
package dotimport

import "github.com/junk1tm/assert"

type (
	E = assert.E
	F = assert.F
)
