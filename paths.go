package main

import "path/filepath"

func GetAsset(path string) string {
	return filepath.Join(g.AssetsDir, path)
}
