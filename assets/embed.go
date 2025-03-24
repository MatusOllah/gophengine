// Package assets provides access to the embedded FNF assets.
package assets

import "embed"

//go:embed [^_]*
var FS embed.FS
