// Package player package is used to support player use cases.
package player

import "github.com/sebdah/go-jek-battleship/coordinates"

// Player represents a player in the program.
type Player struct {
	// moves is a slice of the player moves on the battleground.
	moves coordinates.Coordinates

	// shipPositions is a slice of the player ship positions.
	shipPositions coordinates.Coordinates

	// points is the number of points the player has.
	points int
}

// NewPlayer returns a new Player.
func NewPlayer(shipPositions, moves coordinates.Coordinates) *Player {
	return &Player{
		shipPositions: shipPositions,
		moves:         moves,
		points:        0,
	}
}

// Hit is called when the player hits the opponents ship.
func (p *Player) Hit() {
	p.points++
}

// Miss is called when the player launches a missile, but misses.
func (p *Player) Miss() {
	if p.points > 0 {
		p.points--
	}
}

// Moves returns the coordinates for the players missiles.
func (p *Player) Moves() coordinates.Coordinates {
	return p.moves
}

// Points returns the number of points the user has.
func (p *Player) Points() int {
	return p.points
}

// ShipPositions returns the coordinates for the players ships.
func (p *Player) ShipPositions() coordinates.Coordinates {
	return p.shipPositions
}
