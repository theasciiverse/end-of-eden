//go:build !js
// +build !js

package fs

import (
	"path/filepath"
	"strings"
)

// AssetRoot is the folder that holds the game's assets folder. When it is
// set, every path starting with "assets" (as "./assets/..." does) is read
// from there instead of from the working directory, so the game can keep
// its data beside the binary while the working directory is the player's
// state folder (settings, save, logs, cache and mods).
var AssetRoot = ""

// resolve maps an assets path to AssetRoot, and leaves any other path alone.
func resolve(path string) string {
	if AssetRoot == "" || filepath.IsAbs(path) {
		return path
	}
	clean := filepath.ToSlash(filepath.Clean(path))
	if clean == "assets" || strings.HasPrefix(clean, "assets/") {
		return filepath.Join(AssetRoot, filepath.FromSlash(clean))
	}
	return path
}
