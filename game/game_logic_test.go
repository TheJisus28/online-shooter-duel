package game

import (
	"testing"
)

func TestInitGame(t *testing.T) {
	gs := InitGame(true, 80, 24)

	if gs == nil {
		t.Fatal("GameState should not be nil")
	}

	if len(gs.Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(gs.Players))
	}

	if gs.ScreenWidth != 80 {
		t.Errorf("Expected screen width 80, got %d", gs.ScreenWidth)
	}

	if gs.ScreenHeight != 24 {
		t.Errorf("Expected screen height 24, got %d", gs.ScreenHeight)
	}

	// Verify that players are alive
	for i, player := range gs.Players {
		if !player.Alive {
			t.Errorf("Player %d should be alive", i+1)
		}
		if player.Health != Params.PlayerHealth {
			t.Errorf("Player %d should have %d health, got %d", i+1, Params.PlayerHealth, player.Health)
		}
	}
}

func TestHandlePlayerInput(t *testing.T) {
	gs := InitGame(true, 80, 24)
	player := gs.Players[0]
	initialX := player.X

	// Test left movement
	HandlePlayerInput(gs, player, "move_left")
	if player.X >= initialX {
		t.Error("Player should move left")
	}

	// Test right movement (after moving left)
	leftX := player.X
	HandlePlayerInput(gs, player, "move_right")
	if player.X <= leftX {
		t.Error("Player should move right")
	}

	// Test shooting
	initialBullets := len(gs.Bullets)
	HandlePlayerInput(gs, player, "shoot")
	if len(gs.Bullets) != initialBullets+1 {
		t.Error("Should create a new bullet when shooting")
	}
}

func TestCheckCollisions(t *testing.T) {
	gs := InitGame(true, 80, 24)
	player := gs.Players[0]

	// Create a bullet that hits the player
	bullet := &Bullet{
		X:       player.X + 2, // Center of the player
		Y:       player.Y + 1, // Inside the player's hitbox
		OwnerID: 2,            // From the other player
	}
	gs.Bullets = append(gs.Bullets, bullet)

	initialHealth := player.Health
	CheckCollisions(gs)

	if player.Health != initialHealth-1 {
		t.Errorf("Player health should decrease by 1, expected %d, got %d", initialHealth-1, player.Health)
	}

	if len(gs.Bullets) != 0 {
		t.Error("Bullet should be removed after collision")
	}
}

func TestCheckGameOver(t *testing.T) {
	gs := InitGame(true, 80, 24)

	// Simulate that a player dies
	gs.Players[0].Alive = false
	gs.Players[0].Health = 0

	CheckGameOver(gs)

	if !gs.IsGameOver {
		t.Error("Game should be over when one player dies")
	}

	if gs.Winner != 2 {
		t.Errorf("Player 2 should be the winner, got %d", gs.Winner)
	}
}

func TestUpdateGame(t *testing.T) {
	gs := InitGame(true, 80, 24)

	// Create a bullet
	bullet := &Bullet{
		X:       40,
		Y:       20,
		Speed:   1.0,
		OwnerID: 1,
	}
	gs.Bullets = append(gs.Bullets, bullet)

	initialY := bullet.Y
	UpdateGame(gs)

	if bullet.Y != initialY+bullet.Speed {
		t.Errorf("Bullet should move down, expected %f, got %f", initialY+bullet.Speed, bullet.Y)
	}
}
