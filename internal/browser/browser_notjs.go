//go:build !js

package browser

import pkg_browser "github.com/pkg/browser"

// OpenURL opens a new browser window pointing to url.
func OpenURL(url string) error {
	return pkg_browser.OpenURL(url)
}
