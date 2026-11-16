# End of Eden

**RISE AND SHINE. MAY YOUR DEATH BE A GLORIOUS ONE.**

You wake in a dimly lit room under a red emergency light, the metallic smell of the cryo chamber still in your nose. A panel on the wall blinks *Cryo Sleep Malfunction*. Every other capsule is broken, the world above has had five hundred years of climate collapse and nuclear war, and shortly after you sit up you realise you are not alone. Grab whatever weapon you can find, or gather your strength and go at it with your bare hands.

End of Eden is a roguelite deck-builder in the spirit of Slay the Spire, drawn entirely in your terminal. Every fight is a hand of cards and three action points: stab with the VIBRO Knife, throw up a Block, down an ENRGY Drink X91 for one more action, then end your turn and take what's coming. You can always see what's coming, because every enemy shows its intent. The Plasma Golem spends a turn charging up, and you had better have your shield ready when it lets go.

- **Things that live in the facility**: Rust Mites eating the walls ("Tasty metals..."), a Nanobot Swarm that keeps on growing, a Cleaning Bot that never stopped cleaning and greets you with "Corpse. Clean. Engage.", a Cyber Spider waiting on the ceiling, Laser Drones, Repair Drones, and a Cyber Slime that splits in two when you kill it
- **Weapons and gear**: a rusty Crowbar, the VIBRO Knife, an LZR Pistol, the HAR-II assault rifle and an Arm Mounted Gun, plus Stim Packs, Smoke Bombs, Flash Bangs and Adrenaline Shots
- **Artifacts** that change the rules: a Bio Recycler that heals you on every kill, a Gold Scrapper, Reflective Armor, Combat Gloves and a Headbut Helmet
- **Strange rooms between the fights**: an old vending machine that still sells health, an upgrade station, a barrier guarding a chest, a symbiotic parasite offering you more life, and a strange figure with wares beyond your wildest imagination
- **Seed of the Day**: everyone gets the same run today
- **Mods in Lua**: every card, enemy, artifact and event in the game is a Lua script, and yours can sit beside them

## Controls

- **Left** and **Right**: choose a card; **Tab**: next card, or next target
- **Enter**: play the card, pick its target, take a choice
- **Space**: end your turn
- **s**: look at yourself; **x**: inspect the enemy under the mouse
- **Esc**: the overview of your character, cards, artifacts and log, and quit to the menu
- **Mouse**: click a card to play it, an enemy to target or inspect it, End Turn to end your turn
- **Tutorial** in the main menu: a first fight against training dummies

## History

BigJk (Daniel Schmidt) started End of Eden in April 2023 as a deck-builder that would run entirely in a console, with the game's rules in Go and all of its content in Lua so that anyone could add to it. To bring it to players who never open a terminal, he went as far as writing his own terminal emulator, CRT, for a windowed version with a retro shader, and a fuzzy tester that throws random moves at the game until something breaks. It has grown through twenty-six releases since, with new cards, enemies and artifacts arriving in the summer of 2024.

## Where things live

Run `end_of_eden` from anywhere; it reads the `assets` folder beside it.
Settings (`settings_term.toml`), the save (`session.save`), the logs
(`logs/global.log`, and one `logs/S <date> <time>.txt` per run), the
image cache and your mods (`mods/<name>`) are kept in
`$XDG_STATE_HOME/end-of-eden` (default `~/.local/state/end-of-eden`),
or `%LOCALAPPDATA%\end-of-eden` on Windows. `--audio=false` starts
without sound. The game wants a true colour terminal of at least
112x36.

## Credits

- End of Eden: BigJk (Daniel Schmidt), https://github.com/BigJk/end_of_eden
  and https://bigjk.itch.io/end-of-eden
- Sound effects: CC0 (assets/audio/README.md)
- Licence: MIT (LICENSE)

## This build

The Asciiverse packaging of the terminal version, from
https://github.com/theasciiverse/end-of-eden. What we changed, and which
of upstream's assets are not included and why, is in PATCHES.md. Build
it with Go 1.23 or newer:

    CGO_ENABLED=0 go build -trimpath -tags no_audio -o end_of_eden ./cmd/game
    sh asciiverse/stage.sh <folder>

(Leave out `-tags no_audio` and set `CGO_ENABLED=1`, with ALSA headers on
Linux, for sound.)
