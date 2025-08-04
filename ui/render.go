package ui

import (
	"fmt"
	"strings"

	"shooter-duel/game"

	"github.com/nsf/termbox-go"
)

// GameStates for game state management
const (
	StateMenu = iota
	StateWaitingForClient
	StateConnecting
	StateGameRunning
	StateGameOver
)

// MenuOptions for menu options
const (
	MenuOptionCreate = iota
	MenuOptionJoin
	MenuOptionExit
)

// DrawGame renders the game state
func DrawGame(gs *game.GameState) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw players
	for i, player := range gs.Players {
		if !player.Alive {
			continue
		}

		color := termbox.ColorYellow
		if i == 1 {
			color = termbox.ColorMagenta
		}

		DrawSprite(int(player.X), int(player.Y), player.Sprite, color, termbox.ColorDefault)

		// Draw health bar
		healthBar := fmt.Sprintf("P%d: %d", player.ID, player.Health)
		DrawText(0, i*2, healthBar, color, termbox.ColorDefault)
	}

	// Draw bullets
	for _, b := range gs.Bullets {
		DrawSprite(int(b.X), int(b.Y), game.Params.BulletSprite, termbox.ColorWhite, termbox.ColorDefault)
	}

	// Draw instructions
	instructions := "A/D: Move, J: Shoot, Q: Quit"
	DrawText(0, gs.ScreenHeight-1, instructions, termbox.ColorCyan, termbox.ColorDefault)

	// Draw game over message
	if gs.IsGameOver {
		msg := "GAME OVER"
		if gs.Winner > 0 {
			msg = fmt.Sprintf("Player %d Wins!", gs.Winner)
		}
		DrawCenteredText(gs.ScreenWidth/2, gs.ScreenHeight/2, msg, termbox.ColorRed, termbox.ColorDefault)
	}

	termbox.Flush()
}

// DrawSprite draws a sprite at the specified position
func DrawSprite(x, y int, sprite []string, fg, bg termbox.Attribute) {
	w, h := termbox.Size()
	for row, line := range sprite {
		for col, char := range line {
			if x+col >= 0 && x+col < w && y+row >= 0 && y+row < h {
				termbox.SetCell(x+col, y+row, char, fg, bg)
			}
		}
	}
}

// DrawText draws text at the specified position
func DrawText(x, y int, text string, fg, bg termbox.Attribute) {
	w, h := termbox.Size()
	for i, char := range text {
		if x+i >= 0 && x+i < w && y >= 0 && y < h {
			termbox.SetCell(x+i, y, char, fg, bg)
		}
	}
}

// DrawCenteredText draws horizontally centered text
func DrawCenteredText(centerX, y int, text string, fg, bg termbox.Attribute) {
	x := centerX - len(text)/2
	DrawText(x, y, text, fg, bg)
}

// DrawMenu draws the main menu
func DrawMenu(selectedOption, w, h int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	title := "ONLINE SHOOTER DUEL"
	menuOptions := []string{
		"Create Room (Host)",
		"Join Room (Client)",
		"Exit Game",
	}
	xTitle := (w - len(title)) / 2
	yTitle := h/2 - 4
	for i, r := range title {
		termbox.SetCell(xTitle+i, yTitle, r, termbox.ColorCyan, termbox.ColorDefault)
	}
	yMenu := h / 2
	for i, option := range menuOptions {
		xOption := (w - len(option)) / 2
		color := termbox.ColorWhite
		if i == selectedOption {
			color = termbox.ColorGreen
		}
		for j, r := range option {
			termbox.SetCell(xOption+j, yMenu+i, r, color, termbox.ColorDefault)
		}
	}
	startMsg := "Use Arrows to select, Enter to confirm"
	xStart := (w - len(startMsg)) / 2
	yStart := h/2 + 4
	for i, r := range startMsg {
		termbox.SetCell(xStart+i, yStart, r, termbox.ColorYellow, termbox.ColorDefault)
	}
	termbox.Flush()
}

// DrawWaitingScreen draws the waiting screen
func DrawWaitingScreen(message string, w, h int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Split the message into lines
	lines := strings.Split(message, "\n")

	// Calculate initial position to center all lines
	startY := h/2 - len(lines)/2

	for lineIndex, line := range lines {
		xMsg := (w - len(line)) / 2
		yMsg := startY + lineIndex

		for i, r := range line {
			termbox.SetCell(xMsg+i, yMsg, r, termbox.ColorYellow, termbox.ColorDefault)
		}
	}

	termbox.Flush()
}

// DrawGameOver draws the game over screen
func DrawGameOver(score int, message, restartMsg string, w, h int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	msg := message
	if score > 0 {
		msg = fmt.Sprintf("%s - Final Score: %d", message, score)
	}
	restart := restartMsg
	xMsg := (w - len(msg)) / 2
	yMsg := h / 2
	xRestart := (w - len(restart)) / 2
	yRestart := h/2 + 1
	for i, r := range msg {
		termbox.SetCell(xMsg+i, yMsg, r, termbox.ColorRed, termbox.ColorDefault)
	}
	for i, r := range restart {
		termbox.SetCell(xRestart+i, yRestart, r, termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}
