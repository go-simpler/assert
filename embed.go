//go:build installer

package assert

import _ "embed"

// These files are embedded for the cmd/installer tool to work. Do not use them.
var (
	//go:embed LICENSE
	LICENSE string
	//go:embed assert.go
	MainFile string
	//go:embed dotimport/alias.go
	SupportFile string
)
