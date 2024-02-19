package gophengine

import (
	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/MatusOllah/gophengine/internal/fsutil"
)

func ExtractAssets() error {
	return fsutil.Extract(assets.FS, "assets", flagutil.MustGetBool(FlagSet, "gui"))
}
