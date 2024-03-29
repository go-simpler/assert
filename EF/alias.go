// Package EF provides type aliases for the parent [assert] package.
// It should be dot-imported so that [E] and [F] can be used as local types.
package EF

import "go-simpler.org/assert"

type (
	E = assert.E
	F = assert.F
)
