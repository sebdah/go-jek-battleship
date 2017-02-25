package player

import "github.com/sebdah/go-jek-battleship/coordinates"

// Playerer is the interface for a player type.
type Playerer interface {
	Hit()
	Miss()
	Moves() coordinates.Coordinates
	Points() int
	ShipPositions() coordinates.Coordinates
}
