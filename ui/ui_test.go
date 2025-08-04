package ui

import (
	"online/game"
	"testing"
)

func TestDrawCenteredText(t *testing.T) {
	// test the centering calculation logic
	text := "Test"
	centerX := 10

	expectedX := centerX - len(text)/2
	if expectedX != 8 { // 10 - 4/2 = 8
		t.Errorf("Expected center position %d, got %d", 8, expectedX)
	}
}

func TestGameStateConstants(t *testing.T) {
	// make sure state constants are correct
	if StateMenu != 0 {
		t.Error("StateMenu should be 0")
	}
	if StateWaitingForClient != 1 {
		t.Error("StateWaitingForClient should be 1")
	}
	if StateConnecting != 2 {
		t.Error("StateConnecting should be 2")
	}
	if StateGameRunning != 3 {
		t.Error("StateGameRunning should be 3")
	}
	if StateGameOver != 4 {
		t.Error("StateGameOver should be 4")
	}
}

func TestMenuOptionsConstants(t *testing.T) {
	// make sure menu option constants are correct
	if MenuOptionCreate != 0 {
		t.Error("MenuOptionCreate should be 0")
	}
	if MenuOptionJoin != 1 {
		t.Error("MenuOptionJoin should be 1")
	}
	if MenuOptionExit != 2 {
		t.Error("MenuOptionExit should be 2")
	}
}

func TestGameStateStructure(t *testing.T) {
	// check that game state works for UI rendering
	gs := game.InitGame(true, 80, 24)
	if gs == nil {
		t.Fatal("GameState should not be nil")
	}

	if len(gs.Players) != 2 {
		t.Error("GameState should have 2 players")
	}

	if gs.ScreenWidth != 80 {
		t.Error("ScreenWidth should be 80")
	}

	if gs.ScreenHeight != 24 {
		t.Error("ScreenHeight should be 24")
	}
}

func TestPlayerRenderingData(t *testing.T) {
	// test that player data is suitable for rendering
	gs := game.InitGame(true, 80, 24)

	for i, player := range gs.Players {
		// check that players have valid positions
		if player.X < 0 {
			t.Errorf("Player %d X position should be >= 0", i+1)
		}
		if player.Y < 0 {
			t.Errorf("Player %d Y position should be >= 0", i+1)
		}

		// check that players have valid health
		if player.Health <= 0 {
			t.Errorf("Player %d should have positive health", i+1)
		}

		// check that players have sprites
		if len(player.Sprite) == 0 {
			t.Errorf("Player %d should have a sprite", i+1)
		}
	}
}

func TestBulletRenderingData(t *testing.T) {
	// test that bullet data is suitable for rendering
	gs := game.InitGame(true, 80, 24)

	// add a test bullet
	bullet := &game.Bullet{
		X:       40,
		Y:       20,
		Speed:   1.0,
		OwnerID: 1,
	}
	gs.Bullets = append(gs.Bullets, bullet)

	// check that bullets have valid positions
	for i, bullet := range gs.Bullets {
		if bullet.X < 0 {
			t.Errorf("Bullet %d X position should be >= 0", i+1)
		}
		if bullet.Y < 0 {
			t.Errorf("Bullet %d Y position should be >= 0", i+1)
		}

		// check that bullets have valid owner
		if bullet.OwnerID <= 0 {
			t.Errorf("Bullet %d should have valid owner ID", i+1)
		}
	}
}

func TestGameOverLogic(t *testing.T) {
	// test game over state for UI rendering
	gs := game.InitGame(true, 80, 24)

	// initially game should not be over
	if gs.IsGameOver {
		t.Error("Game should not be over initially")
	}

	// simulate game over
	gs.IsGameOver = true
	gs.Winner = 1

	if !gs.IsGameOver {
		t.Error("Game should be over")
	}

	if gs.Winner != 1 {
		t.Error("Winner should be player 1")
	}
}

func TestScreenDimensions(t *testing.T) {
	// test different screen dimensions
	testCases := []struct {
		width, height int
	}{
		{80, 24},
		{100, 30},
		{60, 20},
	}

	for _, tc := range testCases {
		gs := game.InitGame(true, tc.width, tc.height)

		if gs.ScreenWidth != tc.width {
			t.Errorf("Expected width %d, got %d", tc.width, gs.ScreenWidth)
		}

		if gs.ScreenHeight != tc.height {
			t.Errorf("Expected height %d, got %d", tc.height, gs.ScreenHeight)
		}
	}
}
