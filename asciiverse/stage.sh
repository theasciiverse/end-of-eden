#!/bin/sh
# Copies what an Asciiverse archive of End of Eden ships into a folder:
# the game's assets, minus the files whose licence does not clearly allow
# us to pass them on (see PATCHES.md, "Assets"), plus the notes.
#   asciiverse/stage.sh <folder>
set -eu
here=$(cd "$(dirname "$0")/.." && pwd)
out=$(mkdir -p "$1" && cd "$1" && pwd)   # absolute: the copies below run from assets/
mkdir -p "$out/assets/audio"
cd "$here/assets"
cp -r gen locals scripts tutorial "$out/assets/"
rm -rf "$out/assets/gen/faces"
# CC0 sound effects only (assets/audio/README.md)
for f in damage_1.wav damage_2.mp3 damage_3.mp3 death_scream_1.wav \
         game_over.mp3 btn_deny.wav; do
  cp "audio/$f" "$out/assets/audio/"
done
cp "$here/asciiverse/audio-README.md" "$out/assets/audio/README.md"
cp "$here/LICENSE" "$here/PATCHES.md" "$out/"
cp "$here/asciiverse/README.md" "$out/README.md"
