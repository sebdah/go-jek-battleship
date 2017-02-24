package battleground_test

import (
	"testing"

	"github.com/sebdah/go-jek-battleship/battleground"
	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/sebdah/go-jek-battleship/player"
	"github.com/stretchr/testify/assert"
)

func TestBattlegroundPlay(t *testing.T) {
	tests := map[string]struct {
		battleground   *battleground.Battleground
		result         []string
		opponentPoints int
	}{
		"missile misses": {
			battleground: &battleground.Battleground{
				M: 2,
				Player: &player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
				},
				Opponent: &player.Player{
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 1),
					},
				},
			},
			result:         []string{"B", "_", "O", "_"},
			opponentPoints: 0,
		},
		"missile hits": {
			battleground: &battleground.Battleground{
				M: 2,
				Player: &player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
				},
				Opponent: &player.Player{
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
					},
				},
			},
			opponentPoints: 1,
			result:         []string{"X", "_", "_", "_"},
		},
		"multiple hits and misses": {
			battleground: &battleground.Battleground{
				M: 2,
				Player: &player.Player{
					ShipPositions: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
						*coordinates.NewCoordinate(1, 1),
					},
				},
				Opponent: &player.Player{
					Moves: coordinates.Coordinates{
						*coordinates.NewCoordinate(0, 0),
						*coordinates.NewCoordinate(0, 1),
						*coordinates.NewCoordinate(1, 0),
					},
				},
			},
			opponentPoints: 1,
			result:         []string{"X", "O", "X", "B"},
		},
		"defaults to all _": {
			battleground: battleground.NewBattleground(2, &player.Player{}, &player.Player{}),
			result:       []string{"_", "_", "_", "_"},
		},
	}

	for _, test := range tests {
		result := test.battleground.Play()
		assert.Equal(t, test.result, result)
		assert.Equal(t, test.opponentPoints, test.battleground.Opponent.Points())
	}
}
