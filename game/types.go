package game

// =============================================================================
// GAME STRUCTURES
// =============================================================================

type Hitbox struct {
	Width, Height int
}

type Player struct {
	X, Y   float64
	Sprite []string
	Speed  float64
	Hitbox Hitbox
	ID     int
	Health int
	Alive  bool
}

type Bullet struct {
	X, Y    float64
	Sprite  []string
	Speed   float64
	Hitbox  Hitbox
	OwnerID int
}

type GameState struct {
	Players      []*Player
	Bullets      []*Bullet
	ScreenWidth  int
	ScreenHeight int
	IsGameOver   bool
	Message      string
	Winner       int
}

// =============================================================================
// CONFIGURATION PARAMETERS
// =============================================================================

var Params = struct {
	Player1Sprite []string
	Player2Sprite []string
	PlayerSpeed   float64
	PlayerHitbox  Hitbox
	BulletSprite  []string
	BulletSpeed   float64
	BulletHitbox  Hitbox
	PlayerHealth  int
}{
	Player1Sprite: []string{
		` /^\ `,
		` |'| `,
		` /-\ `,
	},
	Player2Sprite: []string{
		` \_/ `,
		` |'| `,
		` / \ `,
	},
	PlayerSpeed:  2,
	PlayerHitbox: Hitbox{Width: 5, Height: 3},
	PlayerHealth: 3,

	BulletSprite: []string{`^`},
	BulletSpeed:  1.0,
	BulletHitbox: Hitbox{Width: 1, Height: 1},
}
