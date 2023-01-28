package gophengine

import "path/filepath"

func GetAsset(path string) string {
	return filepath.Join(G.AssetsDir, path)
}
