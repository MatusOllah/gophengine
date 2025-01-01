// Package assets provides access to the embedded FNF assets.
package assets

import "embed"

//go:embed *
var FS embed.FS
