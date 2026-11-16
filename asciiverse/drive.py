"""Plays End of Eden for the catalog preview (asciiverse/record-preview.sh).

Reads the tmux pane and presses keys at a watchable pace: New Game from
the main menu, the first choice of each event, Enter through the reward
screens, and in a fight plays the cards in hand from left to right,
picking the first target, ending the turn (Space) when nothing more can be
played. Quits to the main menu and out after about 50 seconds.
    python3 -I drive.py <tmux socket> <session> [debug]
"""
import subprocess
import sys
import time

SOCK, SESS = sys.argv[1], sys.argv[2]
DEBUG = len(sys.argv) > 3
LENGTH = 50.0
T0 = time.time()


def tmux(*args):
    return subprocess.run(["tmux", "-L", SOCK, *args], capture_output=True,
                          text=True).stdout


def screen():
    return tmux("capture-pane", "-p", "-t", SESS)


def key(k, wait=0.9):
    if DEBUG:
        print(round(time.time() - T0, 1), k, file=sys.stderr)
    tmux("send-keys", "-t", SESS, k)
    time.sleep(wait)


def wait_for(text, limit=15.0):
    end = time.time() + limit
    while time.time() < end:
        if text in screen():
            return True
        time.sleep(0.2)
    return False


wait_for("New Game")
time.sleep(2.0)
key("Down", 0.5)
key("Down", 0.7)
key("Enter", 4.0)          # New Game: waking up in the cryo facility
stuck = 0
last = ""
while time.time() - T0 < LENGTH:
    s = screen()
    if "GAME OVER" in s.upper() and "Action Points" not in s:
        break
    if "Action Points" in s:                       # a fight
        if "Select a target" in s:
            key("Enter", 1.2)
        elif "Action Points: (0)" in s or stuck >= 2:
            key("Space", 2.5)
            stuck = 0
        else:
            key("Enter", 1.2)
    else:                                          # events and rewards
        key("Enter", 2.0)
    now = screen()
    stuck = stuck + 1 if now == s else 0
    last = now

for _ in range(4):                                 # the overview menu
    if "Overview" in screen():
        break
    key("Escape", 1.5)
for _ in range(4):
    key("Down", 0.3)
key("Enter", 1.5)                                  # Quit: back to the menu
key("q", 1.0)
s = screen()
if "Main Menu" in s or "New Game" in s:
    key("C-c", 0.5)
