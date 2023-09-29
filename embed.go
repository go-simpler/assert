//go:build copier

package assert

import _ "embed"

// These files are used by the cmd/copier tool. Ignore them.
var (
	//go:embed assert.go
	MainFile string
	//go:embed EF/alias.go
	AliasFile string
)
