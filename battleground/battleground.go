package battleground

import (
	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/sebdah/go-jek-battleship/player"
)

const (
	// none is for unused postitions.
	none = "_"

	// miss is for missile misses.
	miss = "O"

	// hit is for hit battleships.
	hit = "X"

	// alive is for alive battleships.
	alive = "B"
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

			isMissile := b.Opponent.Moves.Include(coordinate)
			isShip := b.Player.ShipPositions.Include(coordinate)

			switch {
			case !isMissile && isShip:
				matrix = append(matrix, alive)
			case isMissile && isShip:
				b.Opponent.Hit()
				matrix = append(matrix, hit)
			case isMissile && !isShip:
				b.Opponent.Miss()
				matrix = append(matrix, miss)
			default:
				matrix = append(matrix, none)
			}
		}
	}

	return matrix
}
