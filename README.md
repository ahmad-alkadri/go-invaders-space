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

   ```sh
   git clone https://github.com/ahmad-alkadri/go-invaders-space.git
   ```

2. Open a terminal in the project directory and run.

   ```sh
   go run main.go
   ```

3. Alternatively, you can build it:
   ```sh
   go build -o go-invaders-space
   ./go-invaders-space
   ```

## Notes
- Tested extensively on `bash`, not yet on Windows
- Uses the [eiannone/keyboard](https://github.com/eiannone/keyboard) package for keyboard input.

Enjoy blasting some aliens!
