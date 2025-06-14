# InvadersSpace

A simple Space Invaders-style game written in Go for the terminal.

## Features
- Playable in the terminal (cross-platform)
- Move your ship left/right and shoot aliens
- Aliens move and shoot back
- Score tracking
- Game over and restart functionality

## Controls
- `a` or Left Arrow: Move left
- `d` or Right Arrow: Move right
- Spacebar: Shoot
- `q`: Quit
- `r`: Restart (after game over)

## Requirements
- Go 1.16 or newer
- Terminal/console with basic ANSI support

## How to Run
1. Clone or download this repository.
2. Open a terminal in the project directory.
3. Run:
   ```sh
   go run main.go
   ```

## Notes
- On Windows, the screen will clear using `cls`.
- Uses the [eiannone/keyboard](https://github.com/eiannone/keyboard) package for keyboard input.

Enjoy blasting some aliens!
