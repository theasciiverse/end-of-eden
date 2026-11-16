# End of Eden (Asciiverse packaging)

BigJk's End of Eden (Go, Bubble Tea, Lua, MIT), packaged for the
Asciiverse catalog (id `end-of-eden`, kind `game`, genre `card`). Not our
code: the game stays as its author made it; list every change in
PATCHES.md and keep each one offerable upstream (UPSTREAM.md).

- `upstream` is github.com/BigJk/end_of_eden (branch `master`); ours is
  branch `asciiverse` (from master `bdfd8e3`, v0.2.2 plus seven
  commits), `origin` github.com/theasciiverse/end-of-eden (private).
- Build: `CGO_ENABLED=0 go build -trimpath -tags no_audio ./cmd/game`
  (terminal version; Windows without `no_audio`). Never `cmd/game_win`
  (`_gl`), `cmd/game_ssh` or WASM.
- Assets: only what `asciiverse/stage.sh` copies may ship (PATCHES.md,
  "Assets": no music, no images, no faces, no Bleeoop sounds).
- Release: tag `v0.2.2-av.<n>` and push it;
  `.github/workflows/asciiverse-release.yml` cross-builds Linux, macOS
  and Windows on one runner.
- Settings, save, logs, cache and mods live in
  `$XDG_STATE_HOME/end-of-eden`. Tests and scripted runs always set
  scratch `HOME` and `XDG_STATE_HOME` in the same command, with
  `--audio=false`, under `systemd-run --user --scope -p MemoryMax=4G`.
- Smallest whole screen: 112x36.
- Preview: `asciiverse/record-preview.sh <folder with end_of_eden and assets>`.
- Catalog drafts: `asciiverse/end-of-eden.toml` and `.md`.
