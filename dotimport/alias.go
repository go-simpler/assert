// Package dotimport provides type aliases for the parent [assert] package.
// It is intended to be dot-imported so that [E] and [F] can be used as local types.
package dotimport

import "go-simpler.org/assert"

type (
	E = assert.E
	F = assert.F
)
