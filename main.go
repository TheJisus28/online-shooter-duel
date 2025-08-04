package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"online/core"
	"online/game"
	"online/network"
	"online/ui"

	"github.com/nsf/termbox-go"
)

// =============================================================================
// GLOBAL CONSTANTS & GAME STATES
// =============================================================================

const (
	gameOverMsg = "GAME OVER"
	restartMsg  = "Press R to restart or Q to quit"
)

// =============================================================================
// MAIN GAME LOOP & STATE MACHINE
// =============================================================================

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	rand.Seed(time.Now().UnixNano())

	w, h := termbox.Size()

	currentState := ui.StateMenu
	menuOptionSelected := ui.MenuOptionCreate

	for {
		switch currentState {
		case ui.StateMenu:
			ui.DrawMenu(menuOptionSelected, w, h)
			newOption, selected := core.HandleMenuInput(menuOptionSelected)
			menuOptionSelected = newOption
			if selected {
				if menuOptionSelected == ui.MenuOptionCreate {
					currentState = ui.StateWaitingForClient
				} else if menuOptionSelected == ui.MenuOptionJoin {
					currentState = ui.StateConnecting
				} else if menuOptionSelected == ui.MenuOptionExit {
					return // Exit the game
				}
			}

		case ui.StateWaitingForClient:
			ui.DrawWaitingScreen("Creating room...", w, h)
			conn, err := network.RunAsHost(w, h)

			// It will always return an error with the IP, so handle it directly
			if err != nil && len(err.Error()) > 20 && err.Error()[:20] == "waiting_for_connecti" {
				hostIP := err.Error()[21:] // Extract IP from message
				ui.DrawWaitingScreen("✅ Room created successfully!\n\n📡 Connection Info:\nIP: "+hostIP+"\nPort: "+network.Port+"\n\n⏳ Waiting for player to connect...", w, h)

				// Now accept the connection
				conn, err = network.AcceptConnection()
				if err != nil {
					ui.DrawGameOver(0, fmt.Sprintf("Connection Error: %s", err.Error()), restartMsg, w, h)
					if core.WaitForRestart() {
						currentState = ui.StateMenu
					} else {
						return
					}
				} else {
					// If connection was successful, continue to the game
					currentState = ui.StateGameRunning
					gameLoop(conn, true, w, h)
					currentState = ui.StateMenu
				}
			} else if err != nil {
				// Real error
				ui.DrawGameOver(0, fmt.Sprintf("Error: %s", err.Error()), restartMsg, w, h)
				if core.WaitForRestart() {
					currentState = ui.StateMenu
				} else {
					return
				}
			}

		case ui.StateConnecting:
			// Close termbox temporarily to allow console input
			termbox.Close()
			conn, err := network.RunAsClient(w, h)
			if err != nil {
				// Reinitialize termbox to show error
				termbox.Init()
				termbox.SetInputMode(termbox.InputEsc)
				ui.DrawGameOver(0, fmt.Sprintf("Error: %s", err.Error()), restartMsg, w, h)
				if core.WaitForRestart() {
					currentState = ui.StateMenu
				} else {
					return
				}
				break
			}

			// Reinitialize termbox
			termbox.Init()
			termbox.SetInputMode(termbox.InputEsc)

			currentState = ui.StateGameRunning
			gameLoop(conn, false, w, h)
			currentState = ui.StateMenu

		case ui.StateGameOver:
			ui.DrawGameOver(0, gameOverMsg, restartMsg, w, h)
			if core.WaitForRestart() {
				currentState = ui.StateMenu
			} else {
				return
			}
		}
	}
}

// =============================================================================
// GAME LOOP FUNCTIONS
// =============================================================================

func gameLoop(conn net.Conn, isHost bool, w, h int) {
	gs := game.InitGame(isHost, w, h)
	if isHost {
		hostGameLoop(gs, conn)
	} else {
		clientGameLoop(gs, conn)
	}
}

func hostGameLoop(gs *game.GameState, conn net.Conn) {
	player1Input := make(chan string)
	player2Input := make(chan string)

	go core.ReadInputFromTerminal(player1Input)
	go network.ReadInputFromNetwork(conn, player2Input)

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for !gs.IsGameOver {
		select {
		case ev := <-player1Input:
			if ev == "quit" {
				gs.IsGameOver = true
				conn.Close()
				break
			}
			game.HandlePlayerInput(gs, gs.Players[0], ev)
		case ev := <-player2Input:
			if ev == "quit" {
				gs.IsGameOver = true
				conn.Close()
				break
			}
			game.HandlePlayerInput(gs, gs.Players[1], ev)
		case <-ticker.C:
			game.UpdateGame(gs)
			game.CheckCollisions(gs)
			game.CheckGameOver(gs)
			network.SendGameState(conn, gs)
			ui.DrawGame(gs)
		}
	}
}

func clientGameLoop(gs *game.GameState, conn net.Conn) {
	playerInput := make(chan string)
	go core.ReadInputFromTerminal(playerInput)

	go func() {
		for {
			if err := network.ReadGameStateFromNetwork(conn, gs); err != nil {
				gs.IsGameOver = true
				return
			}
		}
	}()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for !gs.IsGameOver {
		select {
		case ev := <-playerInput:
			if ev == "quit" {
				gs.IsGameOver = true
				conn.Close()
				break
			}
			conn.Write([]byte(ev + "\n"))
		case <-ticker.C:
			ui.DrawGame(gs)
		}
	}
}
