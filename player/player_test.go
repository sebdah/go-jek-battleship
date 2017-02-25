package player_test

import (
	"testing"

	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/sebdah/go-jek-battleship/player"
	"github.com/stretchr/testify/assert"
)

func TestPlayerPoints(t *testing.T) {
	player := &player.Player{}
	assert.Equal(t, 0, player.Points())
	player.Miss()
	assert.Equal(t, 0, player.Points())
	player.Hit()
	assert.Equal(t, 1, player.Points())
	player.Hit()
	assert.Equal(t, 2, player.Points())
	player.Miss()
	assert.Equal(t, 1, player.Points())
}

func TestPlayerMoves(t *testing.T) {
	tests := []struct {
		coords coordinates.Coordinates
	}{
		{
			coords: coordinates.Coordinates{
				*coordinates.NewCoordinate(1, 0),
				*coordinates.NewCoordinate(1, 1),
			},
		},
		{
			coords: coordinates.Coordinates{
				*coordinates.NewCoordinate(1, 1),
			},
		},
	}

	for _, test := range tests {
		p := player.NewPlayer(test.coords, test.coords)
		assert.Equal(t, test.coords, p.Moves())
	}
}

func TestPlayerShipPositions(t *testing.T) {
	tests := []struct {
		coords coordinates.Coordinates
	}{
		{
			coords: coordinates.Coordinates{
				*coordinates.NewCoordinate(1, 0),
				*coordinates.NewCoordinate(1, 1),
			},
		},
		{
			coords: coordinates.Coordinates{
				*coordinates.NewCoordinate(1, 1),
			},
		},
	}

	for _, test := range tests {
		p := player.NewPlayer(test.coords, test.coords)
		assert.Equal(t, test.coords, p.ShipPositions())
	}
}
