# Go Invaders Space!

A simple Space Invaders-style game written in Go for the terminal.

| ![invaders](https://github.com/user-attachments/assets/acb05dbf-e65d-45ce-83e8-869cb76ff9ff) |
| --- |
| Game demo, played on terminal, recorded using [asciinema](https://docs.asciinema.org/) |

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

## Releases and Installers
- Automated packaging is configured via GoReleaser.
- On Linux, .deb and .rpm packages are created.
- On Windows and Linux, prebuilt archives (zip/tar.gz) are uploaded.

### How to trigger a release (GitHub Actions)
1. Go to the Actions tab, open the "Release" workflow, click "Run workflow".
2. Enter a new tag like `v1.0.0` and run. The workflow will tag and publish a release with artifacts.
3. Alternatively, push a tag `v*` to trigger the workflow automatically.

### Run GoReleaser locally (optional)
Requirements: Go and GoReleaser installed.

```sh
# snapshot build without publishing
goreleaser release --snapshot --clean
```

Enjoy blasting some aliens!
