// Package player package is used to support player use cases.
package player

import "github.com/sebdah/go-jek-battleship/coordinates"

// Player represents a player in the program.
type Player struct {
	// Moves is a slice of the player moves on the battleground.
	Moves coordinates.Coordinates

	// ShipPositions is a slice of the player ship positions.
	ShipPositions coordinates.Coordinates

	// points is the number of points the player has.
	points int
}

// Points returns the number of points the user has.
func (p *Player) Points() int {
	return p.points
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
