package config

import "github.com/sebdah/go-jek-battleship/player"

// errInvalidConfig is thrown if the validation fails.
type errInvalidConfig struct {
	message string
}

func (e *errInvalidConfig) Error() string {
	return e.message
}

// Config holds the configuration for the game.
type Config struct {
	// M defines the size of the battleground. 0 < M < 10.
	M int

	// S is the total number of ships. 0 < S <= M/2
	S int

	// T is the number of missiles per player. 0 < T < 100.
	T int

	// Player1 represents the first player.
	Player1 player.Player

	// Player2 represents the first player.
	Player2 player.Player
}

// Validate checks the integrity of the configuration.
func (c *Config) Validate() error {
	if len(c.Player1.Moves) != c.T {
		return &errInvalidConfig{
			message: "Player 1 has wrong number of moves",
		}
	}

	if len(c.Player2.Moves) != c.T {
		return &errInvalidConfig{
			message: "Player 2 has wrong number of moves",
		}
	}

	if len(c.Player1.ShipPositions) != c.S {
		return &errInvalidConfig{
			message: "Player 1 has wrong number of ships",
		}
	}

	if len(c.Player2.ShipPositions) != c.S {
		return &errInvalidConfig{
			message: "Player 2 has wrong number of ships",
		}
	}

	return nil
}
