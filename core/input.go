package core

import (
	"github.com/nsf/termbox-go"
)

// ReadInputFromTerminal reads user input from the terminal
func ReadInputFromTerminal(inputChan chan string) {
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
				inputChan <- "quit"
				return
			}
			if ev.Ch == 'a' {
				inputChan <- "move_left"
			}
			if ev.Ch == 'd' {
				inputChan <- "move_right"
			}
			if ev.Ch == 'j' {
				inputChan <- "shoot"
			}
		}
	}
}

// WaitForRestart waits for the user to press R to restart or Q to exit
func WaitForRestart() bool {
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Ch == 'r' || ev.Ch == 'R' {
				return true
			}
			if ev.Ch == 'q' || ev.Ch == 'Q' {
				return false
			}
			if ev.Key == termbox.KeyEsc {
				return false
			}
		}
	}
}

// HandleMenuInput handles user input in the menu
func HandleMenuInput(selectedOption int) (int, bool) {
	ev := termbox.PollEvent()
	if ev.Type == termbox.EventKey {
		switch ev.Key {
		case termbox.KeyEnter:
			return selectedOption, true
		case termbox.KeyArrowUp:
			return (selectedOption - 1 + 3) % 3, false
		case termbox.KeyArrowDown:
			return (selectedOption + 1) % 3, false
		case termbox.KeyEsc, termbox.KeyCtrlQ:
			return 2, true // Select Exit Game
		}
	}
	return selectedOption, false
}
