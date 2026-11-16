package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BigJk/end_of_eden/internal/fs"
)

// useStateDir keeps the game's data beside the binary and everything it
// writes in a per-user folder. The assets are read from the folder that
// holds the executable (when an assets folder is there), and the working
// directory becomes the state folder, which then holds settings_term.toml,
// session.save, logs, the image cache and mods:
//
//   - Windows: %LOCALAPPDATA%\end-of-eden
//   - elsewhere: $XDG_STATE_HOME/end-of-eden (default ~/.local/state/end-of-eden)
//
// Run from a source checkout (go run ./cmd/game) there is no assets folder
// beside the binary, so the game reads ./assets as before.
func useStateDir() error {
	assetRoot, _ := os.Getwd()
	if exe, err := os.Executable(); err == nil {
		if exe, err := filepath.EvalSymlinks(exe); err == nil {
			if st, err := os.Stat(filepath.Join(filepath.Dir(exe), "assets")); err == nil && st.IsDir() {
				assetRoot = filepath.Dir(exe)
			}
		}
	}
	fs.AssetRoot = assetRoot

	var dir string
	if runtime.GOOS == "windows" && os.Getenv("LOCALAPPDATA") != "" {
		dir = filepath.Join(os.Getenv("LOCALAPPDATA"), "end-of-eden")
	} else if xdg := os.Getenv("XDG_STATE_HOME"); xdg != "" {
		dir = filepath.Join(xdg, "end-of-eden")
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		dir = filepath.Join(home, ".local", "state", "end-of-eden")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.Chdir(dir)
}
