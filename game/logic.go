package game

// =============================================================================
// GAME LOGIC FUNCTIONS
// =============================================================================

// InitGame initializes a new game state
func InitGame(isHost bool, w, h int) *GameState {
	// Host is always at the bottom, client at the top
	var player1, player2 *Player

	if isHost {
		// Host: player 1 at bottom
		player1 = &Player{
			X:      float64(w) / 4,
			Y:      float64(h - 4), // Bottom
			Sprite: Params.Player1Sprite,
			Speed:  Params.PlayerSpeed,
			Hitbox: Params.PlayerHitbox,
			ID:     1,
			Health: Params.PlayerHealth,
			Alive:  true,
		}
		// Client: player 2 at top
		player2 = &Player{
			X:      float64(w) / 4 * 3,
			Y:      2, // Top
			Sprite: Params.Player2Sprite,
			Speed:  Params.PlayerSpeed,
			Hitbox: Params.PlayerHitbox,
			ID:     2,
			Health: Params.PlayerHealth,
			Alive:  true,
		}
	} else {
		// Client: player 1 at top
		player1 = &Player{
			X:      float64(w) / 4,
			Y:      2, // Top
			Sprite: Params.Player1Sprite,
			Speed:  Params.PlayerSpeed,
			Hitbox: Params.PlayerHitbox,
			ID:     1,
			Health: Params.PlayerHealth,
			Alive:  true,
		}
		// Host: player 2 at bottom
		player2 = &Player{
			X:      float64(w) / 4 * 3,
			Y:      float64(h - 4), // Bottom
			Sprite: Params.Player2Sprite,
			Speed:  Params.PlayerSpeed,
			Hitbox: Params.PlayerHitbox,
			ID:     2,
			Health: Params.PlayerHealth,
			Alive:  true,
		}
	}

	players := []*Player{player1, player2}

	return &GameState{
		Players:      players,
		Bullets:      make([]*Bullet, 0),
		ScreenWidth:  w,
		ScreenHeight: h,
		IsGameOver:   false,
		Winner:       0,
	}
}

// UpdateGame updates the game state (positions, bullets, etc.)
func UpdateGame(gs *GameState) {
	// Update player positions
	for _, p := range gs.Players {
		if !p.Alive {
			continue
		}
		if p.X < 0 {
			p.X = 0
		}
		if p.X+float64(p.Hitbox.Width) > float64(gs.ScreenWidth) {
			p.X = float64(gs.ScreenWidth - p.Hitbox.Width)
		}
	}

	// Update bullets
	bulletsToKeep := []*Bullet{}
	for _, b := range gs.Bullets {
		b.Y += b.Speed // Negative speed goes up, positive speed goes down
		if b.Y >= -1 && b.Y < float64(gs.ScreenHeight)+1 {
			bulletsToKeep = append(bulletsToKeep, b)
		}
	}
	gs.Bullets = bulletsToKeep
}

// CheckCollisions checks collisions between bullets and players
func CheckCollisions(gs *GameState) {
	bulletsToKeep := []*Bullet{}

	for _, bullet := range gs.Bullets {
		hit := false
		for _, player := range gs.Players {
			if !player.Alive {
				continue
			}

			// Check collision between bullet and player
			if bullet.X >= player.X &&
				bullet.X <= player.X+float64(player.Hitbox.Width) &&
				bullet.Y >= player.Y &&
				bullet.Y <= player.Y+float64(player.Hitbox.Height) &&
				bullet.OwnerID != player.ID {

				player.Health--
				if player.Health <= 0 {
					player.Alive = false
				}
				hit = true
				break
			}
		}

		if !hit {
			bulletsToKeep = append(bulletsToKeep, bullet)
		}
	}

	gs.Bullets = bulletsToKeep
}

// CheckGameOver checks if the game has ended
func CheckGameOver(gs *GameState) {
	alivePlayers := 0
	lastAlivePlayer := 0

	for _, player := range gs.Players {
		if player.Alive {
			alivePlayers++
			lastAlivePlayer = player.ID
		}
	}

	if alivePlayers <= 1 {
		gs.IsGameOver = true
		if alivePlayers == 1 {
			gs.Winner = lastAlivePlayer
		}
	}
}

// HandlePlayerInput processes player input
func HandlePlayerInput(gs *GameState, p *Player, input string) {
	if !p.Alive {
		return
	}

	switch input {
	case "move_left":
		p.X -= p.Speed
	case "move_right":
		p.X += p.Speed
	case "shoot":
		// Determine bullet direction based on player position
		var bulletY float64
		var bulletSpeed float64

		if p.Y < 10 { // Player at the top (client)
			bulletY = float64(p.Y + float64(p.Hitbox.Height)) // Shoot from the bottom of the sprite
			bulletSpeed = Params.BulletSpeed                  // Positive speed to go down
		} else { // Player at the bottom (host)
			bulletY = float64(p.Y - 1)        // Shoot from the top of the sprite
			bulletSpeed = -Params.BulletSpeed // Negative speed to go up
		}

		bullet := &Bullet{
			X:       float64(p.X + float64(p.Hitbox.Width)/2 - 0.5), // Center the bullet horizontally
			Y:       bulletY,
			Sprite:  Params.BulletSprite,
			Speed:   bulletSpeed,
			Hitbox:  Params.BulletHitbox,
			OwnerID: p.ID,
		}
		gs.Bullets = append(gs.Bullets, bullet)
	}
}
