package battleground

import (
	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/sebdah/go-jek-battleship/player"
)

const (
	// NONE is for unused postitions.
	NONE = "_"

	// MISS is for missile misses.
	MISS = "O"

	// HIT is for hit battleships.
	HIT = "X"

	// ALIVE is for alive battleships.
	ALIVE = "B"
)

// Battleground defines the battleground for the game.
type Battleground struct {
	// M defines the size of the battleground. 0 < M < 10.
	M int

	// Player represents the first player.
	Player *player.Player

	// Opponent represents the first player.
	Opponent *player.Player
}

// NewBattleground sets up a new battleground.
func NewBattleground(m int, player, opponent *player.Player) *Battleground {
	return &Battleground{
		M:        m,
		Player:   player,
		Opponent: opponent,
	}
}

// Play is starting the game.
func (b *Battleground) Play() []string {
	matrix := []string{}

	for y := 0; y < b.M; y++ {
		for x := 0; x < b.M; x++ {
			coordinate := coordinates.NewCoordinate(x, y)

			isMissile := false
			if b.Opponent.Moves.Include(coordinate) {
				isMissile = true
			}

			isShip := false
			if b.Player.ShipPositions.Include(coordinate) {
				isShip = true
			}

			switch {
			case !isMissile && isShip:
				matrix = append(matrix, ALIVE)
			case isMissile && isShip:
				b.Opponent.Hit()
				matrix = append(matrix, HIT)
			case isMissile && !isShip:
				b.Opponent.Miss()
				matrix = append(matrix, MISS)
			default:
				matrix = append(matrix, NONE)
			}
		}
	}

	return matrix
}
