//go:build installer

package assert

import _ "embed"

// These files are used by the `cmd/installer` tool. Ignore them.
var (
	//go:embed assert.go
	MainFile string
	//go:embed dotimport/alias.go
	SupportFile string
)
