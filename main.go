package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type vec struct{ x, y int }

type ship struct {
	pos vec
}

type alien struct {
	pos   vec
	typ   int
	alive bool
}

type bullet struct {
	pos    vec
	active bool
}

const (
	frameWidth  = 60
	frameHeight = 20
	alienRows   = 3
	alienCols   = 6
	fps         = 20 // Frame per seconds
)

var (
	player        ship
	aliens        []alien
	playerBullets []bullet
	alienBullets  []bullet
	score         int
	gameOver      bool
	gameWon       bool
	frame         [][]rune
	prevFrame     [][]rune
	alienSpeed    = 500 * time.Millisecond
	lastAlienMove time.Time
	alienDir      = 1
	alienSprites  = []string{
		"/o  o\\",
		"  oo  \n<xxxx>",
		" xxxx \n /oo\\",
	}
)

func init() {
	// Initialize frame buffer
	frame = make([][]rune, frameHeight)
	prevFrame = make([][]rune, frameHeight) // Initialize previous frame buffer
	for i := range frame {
		frame[i] = make([]rune, frameWidth)
		prevFrame[i] = make([]rune, frameWidth)
	}
}

func initGame() {
	player = ship{pos: vec{x: frameWidth / 2, y: frameHeight - 2}}
	aliens = make([]alien, 0, alienRows*alienCols)
	playerBullets = make([]bullet, 0, 3)
	alienBullets = make([]bullet, 0, 5)
	alienDir = 1
	lastAlienMove = time.Now()

	// Create aliens in grid pattern
	for y := range alienRows {
		for x := range alienCols {
			aliens = append(aliens, alien{
				pos:   vec{x: 10 + x*7, y: 3 + y*3},
				typ:   y,
				alive: true,
			})
		}
	}
	score = 0
	gameOver = false
}

func drawSprite(x, y int, sprite string) {
	lines := strings.Split(sprite, "\n")
	for i, line := range lines {
		row := y + i
		if row < 0 || row >= frameHeight {
			continue
		}
		for j, ch := range line {
			col := x + j
			if col >= 0 && col < frameWidth {
				frame[row][col] = ch
			}
		}
	}
}

func clearFrame() {
	for y := range frame {
		for x := range frame[y] {
			if y == frameHeight-1 {
				frame[y][x] = '-'
			} else {
				frame[y][x] = ' '
			}
		}
	}
}

func draw() {
	clearFrame()

	// Draw player
	if !gameOver {
		drawSprite(player.pos.x, player.pos.y, "/#\\")
	}

	// Draw bullets
	for _, b := range playerBullets {
		if b.active {
			drawSprite(b.pos.x, b.pos.y, "o")
		}
	}
	for _, b := range alienBullets {
		if b.active {
			drawSprite(b.pos.x, b.pos.y, "x")
		}
	}

	// Draw aliens
	for _, a := range aliens {
		if a.alive {
			drawSprite(a.pos.x, a.pos.y, alienSprites[a.typ])
		}
	}

	// Draw score
	scoreStr := fmt.Sprintf("Score: %d", score)
	if len(scoreStr) > frameWidth {
		scoreStr = scoreStr[:frameWidth]
	}
	for i, ch := range scoreStr {
		if i < frameWidth {
			frame[0][i] = ch
		}
	}

	// Draw game over
	if gameOver {
		var msg string
		if gameWon {
			msg = "You Won! Press 'r' to play again or 'q' to quit"
		} else {
			msg = "Game Over! Press 'r' to restart or 'q' to quit"
		}

		startX := max(frameWidth/2-len(msg)/2, 0)
		for i, ch := range msg {
			x := startX + i
			if x < frameWidth {
				frame[frameHeight/2][x] = ch
			}
		}
	}

	// Efficient rendering: only update changed cells
	fmt.Print("\033[H") // Move cursor to top-left
	for y := range frameHeight {
		for x := range frameWidth {
			// Only print if cell changed
			if frame[y][x] != prevFrame[y][x] {
				fmt.Printf("\033[%d;%dH%c", y+1, x+1, frame[y][x])
				prevFrame[y][x] = frame[y][x]
			}
		}
	}
}

// FIXED: Proper bounding box collision detection
func isColliding(bulletPos vec, alienPos vec) bool {
	// Alien bounding box: 6 characters wide, 2 tall
	return bulletPos.x >= alienPos.x &&
		bulletPos.x <= alienPos.x+5 &&
		bulletPos.y >= alienPos.y &&
		bulletPos.y <= alienPos.y+1
}

func updatePlayerBullets() {
	newBullets := playerBullets[:0]
	for _, b := range playerBullets {
		if !b.active || b.pos.y <= 0 {
			continue
		}

		newPos := vec{x: b.pos.x, y: b.pos.y - 1}
		hit := false

		// Check collision with aliens using bounding boxes
		for i := range aliens {
			if !aliens[i].alive {
				continue
			}
			if isColliding(newPos, aliens[i].pos) {
				aliens[i].alive = false
				score += (aliens[i].typ + 1) * 10
				hit = true
				break
			}
		}

		if !hit {
			b.pos = newPos
			newBullets = append(newBullets, b)
		}
	}
	playerBullets = newBullets
}

func updateAlienBullets() {
	newBullets := alienBullets[:0]
	for _, b := range alienBullets {
		if !b.active || b.pos.y >= frameHeight-1 {
			continue
		}

		b.pos.y++
		// Check collision with player (player is 3 characters wide)
		if b.pos.y == player.pos.y &&
			b.pos.x >= player.pos.x &&
			b.pos.x <= player.pos.x+2 {
			gameOver = true
		}
		newBullets = append(newBullets, b)
	}
	alienBullets = newBullets
}

func moveAliens() {
	if time.Since(lastAlienMove) < alienSpeed {
		return
	}
	lastAlienMove = time.Now()

	// Check if we need to change direction and move down
	changeDir := false
	for _, a := range aliens {
		if !a.alive {
			continue
		}
		if (alienDir == 1 && a.pos.x >= frameWidth-7) ||
			(alienDir == -1 && a.pos.x <= 1) {
			changeDir = true
			break
		}
	}

	if changeDir {
		alienDir *= -1
		// Move all aliens down
		for i := range aliens {
			if aliens[i].alive {
				aliens[i].pos.y++
				// Check if alien reached bottom
				if aliens[i].pos.y >= frameHeight-3 {
					gameOver = true
					return
				}
			}
		}
	} else {
		// Move horizontally
		for i := range aliens {
			if aliens[i].alive {
				aliens[i].pos.x += alienDir
			}
		}
	}
}

func alienShoot() {
	if len(alienBullets) >= 5 {
		return
	}

	// Find lowest alien in each column
	columnAliens := make(map[int]int)
	for i, a := range aliens {
		if !a.alive {
			continue
		}
		if idx, ok := columnAliens[a.pos.x]; !ok || aliens[idx].pos.y < a.pos.y {
			columnAliens[a.pos.x] = i
		}
	}

	if len(columnAliens) == 0 {
		return
	}

	keys := make([]int, 0, len(columnAliens))
	for k := range columnAliens {
		keys = append(keys, k)
	}

	alienIdx := columnAliens[keys[rand.Intn(len(keys))]]
	typ := aliens[alienIdx].typ

	// Adjusted shooting chance: top aliens shoot more often
	var chance int
	switch typ {
	case 0: // Top row
		chance = 15 // Highest chance (15%)
	case 1: // Middle row
		chance = 10 // Medium chance (10%)
	case 2: // Bottom row
		chance = 5 // Lowest chance (5%)
	}

	// Only shoot if random value is within the chance threshold
	if rand.Intn(100) > chance {
		return
	}

	alienBullets = append(alienBullets, bullet{
		pos:    vec{x: aliens[alienIdx].pos.x + 2, y: aliens[alienIdx].pos.y + 2},
		active: true,
	})
}

func checkWin() {
	for _, a := range aliens {
		if a.alive {
			return
		}
	}
	gameOver = true
	gameWon = true
}

func updateGame() {
	updatePlayerBullets()
	updateAlienBullets()
	moveAliens()
	alienShoot()
	checkWin()

	// Increase speed as aliens are eliminated
	aliveCount := 0
	for _, a := range aliens {
		if a.alive {
			aliveCount++
		}
	}
	if aliveCount < len(aliens)/2 {
		alienSpeed = 300 * time.Millisecond
	} else {
		alienSpeed = 500 * time.Millisecond
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Clear screen and hide cursor at start
	fmt.Print("\033[2J\033[H\033[?25l")
	defer func() {
		// Show cursor when exiting
		fmt.Print("\033[?25h")
	}()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	initGame()
	frameTime := time.Second / time.Duration(fps)
	ticker := time.NewTicker(frameTime)
	defer ticker.Stop()

	keyChan := make(chan keyboard.Key, 10)
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			if key != 0 {
				keyChan <- key
			} else {
				switch char {
				case 'q', 'r', 'a', 'd', ' ':
					keyChan <- keyboard.Key(char)
				}
			}
		}
	}()

	for {
		select {
		case <-ticker.C:
			if !gameOver {
				updateGame()
			}
			draw()

		case key := <-keyChan:
			if gameOver {
				switch key {
				case 'r':
					initGame()
				case 'q':
					fmt.Print("\033[2J\033[H")
					return
				}
				continue
			}

			switch key {
			case keyboard.KeyArrowLeft, 'a':
				if player.pos.x > 1 {
					player.pos.x--
				}
			case keyboard.KeyArrowRight, 'd':
				if player.pos.x < frameWidth-4 {
					player.pos.x++
				}
			case keyboard.KeySpace:
				if len(playerBullets) < 3 {
					playerBullets = append(playerBullets, bullet{
						pos:    vec{x: player.pos.x + 1, y: player.pos.y - 1},
						active: true,
					})
				}
			case 'q':
				fmt.Print("\033[2J\033[H")
				return
			}
		}
	}
}
