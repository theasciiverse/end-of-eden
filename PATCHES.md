# Changes from upstream

This is the Asciiverse packaging of
[End of Eden](https://github.com/BigJk/end_of_eden) by BigJk (Daniel
Schmidt), a roguelite deck-builder in Go on Bubble Tea and Lua.
Upstream is the `upstream` remote. This build is upstream's `master` at
`bdfd8e3` (2024-09-02): the `v0.2.2` release tag (2024-08-31) plus the
seven commits after it, which fix saves breaking after a while and add
the Nanobot Swarm, the Rest and Long Rest cards and a few more artifacts
and cards. Our tags are therefore `v0.2.2-av.N`. The `asciiverse`
branch carries the changes below, each small enough to offer upstream
(UPSTREAM.md).

We ship the terminal version only (`./cmd/game`, upstream's
`end_of_eden_term`): no `_gl` window, no SSH server, no WASM build.

## Code

- **State folder** (`cmd/game/statedir.go`, `cmd/game/main.go`,
  `internal/fs`): the game read `./assets` and wrote
  `settings_term.toml`, `session.save`, `logs/` and `cache/` in the
  working directory, so it had to be started from its own folder. Now
  the assets are read from beside the executable (`fs.AssetRoot`) and
  the working directory becomes `$XDG_STATE_HOME/end-of-eden` (default
  `~/.local/state/end-of-eden`; `%LOCALAPPDATA%\end-of-eden` on
  Windows), which holds the settings, the save, the logs, the image
  cache and any mods (`mods/<name>`). From a source checkout, with no
  assets beside the binary, `./assets` is read as before.
- **Image cache** (`system/image/cache.go`): the `cache` folder was made
  in an `init()`, before the working directory was chosen; it is now
  made when the first image is cached.
- **Face data optional** (`system/gen/faces`, `cmd/game/main.go`): the
  game starts without `assets/gen/faces`; merchants then have no face.
- **Text logo** (`ui/menus/mainmenu`): without `assets/images/title.jpg`
  the main menu shows the game's own text logo (`ui.Title`) in red.

Session logging is upstream's own: `logs/global.log` for the program and
one `logs/S <date> <time>.txt` per run with the run's events, both in
the state folder.

## Assets

Upstream's `assets/audio/README.md` and README credits were checked file
by file (2026-11-16). `asciiverse/stage.sh` copies into each archive only
what is clearly licensed for redistribution:

- **Shipped**: the Lua content (`assets/scripts`, `assets/tutorial`),
  the German translation (`assets/locals`), the merchant and loading
  lines (`assets/gen/*.txt`), all part of the MIT repository; and six
  CC0 sound effects, `damage_1.wav`, `damage_2.mp3`, `damage_3.mp3`,
  `death_scream_1.wav`, `game_over.mp3`, `btn_deny.wav` (credited in the
  archive's `assets/audio/README.md`).
- **Left out**:
  - `assets/audio/music/` (two tracks by synthroton): "permission of use
    is only granted for End of Eden and its forks", not an open licence.
  - `btn_complete.wav`, `btn_menu.wav`, `interface_move.wav` (Bleeoop's
    Interface Bleeps): royalty free for use in projects, but the licence
    forbids sharing the files themselves and sublicensing.
  - `new_artifacts.wav`, `new_cards.wav`: no licence stated (added with
    synthroton's "additional audio work").
  - `assets/images/` (29 event and title pictures): no licence stated;
    the README says assets are licensed per folder and this folder has
    no README.
  - `assets/gen/faces/` (ASCII face parts from Huw Millward's Warsim
    Generator Toolbox): used with the author's permission given in a
    comment, no licence for others.
  - `assets/fonts/`, `assets/shader/`: only for the `_gl` window.

The files stay in the repository's history as upstream has them; they
are not in our release archives.

## Sound

The sound library (beep on oto 0.7) needs cgo on Linux and macOS, so
those builds use the game's own `no_audio` build tag and are silent.
Windows builds keep their sound.

## Packaging

- `.github/workflows/asciiverse-release.yml` (new): av archives for Linux
  (static, x86-64 and ARM64), macOS (Intel and Apple silicon) and Windows
  (x86-64 and ARM64), plus SHA256SUMS, cross-built with cgo off on one
  runner.
- `asciiverse/stage.sh` (new): the assets and notes for an archive.
- `asciiverse/README.md`, `asciiverse/audio-README.md` (new): the
  archive's README and its audio credits.
- `asciiverse/record-preview.sh`, `asciiverse/drive.py` (new): the
  catalog preview.
- `asciiverse/end-of-eden.toml`, `asciiverse/end-of-eden.md`,
  `asciiverse/end-of-eden.cast` (new): the catalog manifest, store page
  and preview, as drafts for hub.
- `CLAUDE.md`, `PATCHES.md`, `UPSTREAM.md` (new): notes.
