# Changes for upstream

Hi! The Asciiverse build of End of Eden carries a few small changes,
written so they can be taken or left as they are. PATCHES.md lists them
with the packaging files.

## Settings, save and logs in a per-user folder

`cmd/game/statedir.go` and `internal/fs/assets.go`: the terminal version
now reads its assets from beside the executable and makes a per-user
folder its working directory (`$XDG_STATE_HOME/end-of-eden`, default
`~/.local/state/end-of-eden`, or `%LOCALAPPDATA%\end-of-eden` on
Windows), so it can be started from anywhere, from a read-only install
folder, or by a launcher. `fs.AssetRoot` maps every `assets/...` path
read through `internal/fs`; everything else is untouched. With no assets
folder beside the binary (`go run ./cmd/game`) it reads `./assets` as
before.

## Smaller ones

- `system/image/cache.go`: the `cache` folder is made when an image is
  first cached instead of in `init()`, so it lands in the chosen
  working directory.
- `system/gen/faces`: a missing face folder no longer stops the game at
  start, and `Gen` returns an empty face for an unknown id.
- `ui/menus/mainmenu`: without `title.jpg` the main menu shows the text
  logo from `ui.Title`.

## Asset licences

We ship only assets with a licence that allows passing the files on
(PATCHES.md, "Assets"). If the event images in `assets/images` are your
own work under the repository's MIT licence, a line in a README there
would settle it for packagers, and the same for `new_artifacts.wav` and
`new_cards.wav`.

Thank you for End of Eden!
