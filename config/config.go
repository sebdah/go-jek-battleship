package config

import "github.com/sebdah/go-jek-battleship/player"

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
