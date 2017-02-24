package config

import (
	"testing"

	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/sebdah/go-jek-battleship/player"
	"github.com/stretchr/testify/assert"
)

func TestConfigValidate(t *testing.T) {
	tests := map[string]struct {
		config *Config
		valid  bool
	}{
		"valid": {
			config: &Config{
				S: 1,
				T: 2,
				Player1: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
				Player2: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
			},
			valid: true,
		},
		"wrong player1 moves": {
			config: &Config{
				S: 1,
				T: 2,
				Player1: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 1),
					},
				},
				Player2: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
			},
			valid: false,
		},
		"wrong player2 moves": {
			config: &Config{
				S: 1,
				T: 2,
				Player1: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
				Player2: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
				},
			},
			valid: false,
		},
		"wrong player1 ships": {
			config: &Config{
				S: 1,
				T: 2,
				Player1: player.Player{
					ShipPositions: coordinates.Coordinates{},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
				Player2: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
			},
			valid: false,
		},
		"wrong player2 ships": {
			config: &Config{
				S: 1,
				T: 2,
				Player1: player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
				Player2: player.Player{
					ShipPositions: coordinates.Coordinates{},
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
					},
				},
			},
			valid: false,
		},
	}

	for _, test := range tests {
		err := test.config.Validate()
		if test.valid {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestErrInvalidConfig(t *testing.T) {
	e := errInvalidConfig{
		message: "something",
	}
	assert.Equal(t, "something", e.Error())
}
