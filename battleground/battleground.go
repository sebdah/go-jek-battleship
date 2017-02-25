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
	// m defines the size of the battleground. 0 < m < 10.
	m int

	// Player represents the first player.
	Player player.Playerer

	// Opponent represents the first player.
	Opponent player.Playerer
}

// NewBattleground sets up a new battleground.
func NewBattleground(m int, player, opponent player.Playerer) *Battleground {
	return &Battleground{
		m:        m,
		Player:   player,
		Opponent: opponent,
	}
}

// M returns the size of the one side of the battleground.
func (b *Battleground) M() int {
	return b.m
}

// Play is starting the game.
func (b *Battleground) Play() []string {
	matrix := []string{}

	for y := 0; y < b.m; y++ {
		for x := 0; x < b.m; x++ {
			coordinate := coordinates.NewCoordinate(x, y)

			moves := b.Opponent.Moves()
			isMissile := moves.Include(coordinate)
			ships := b.Player.ShipPositions()
			isShip := ships.Include(coordinate)

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
