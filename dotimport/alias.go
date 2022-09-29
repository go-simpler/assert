// Package dotimport provides type aliases for the parent [assert] package. It
// is intended to be imported using dot syntax so that [E] and [F] can be used
// as if they were local types.
package dotimport

import "github.com/junk1tm/assert"

type (
	E = assert.E
	F = assert.F
)