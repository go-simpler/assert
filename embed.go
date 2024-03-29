//go:build cp

package assert

import _ "embed"

var (
	//go:embed assert.go
	MainFile string
	//go:embed EF/alias.go
	AliasFile string
)
