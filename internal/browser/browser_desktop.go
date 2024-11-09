//go:build !wasm

package browser

import pkg_browser "github.com/pkg/browser"

func OpenURL(url string) error {
	return pkg_browser.OpenURL(url)
}
