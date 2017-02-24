package config

import (
	"bufio"
	"io"
	"strconv"

	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/sebdah/go-jek-battleship/player"
)

// FromReader builds a configuration object from a file input source.
func FromReader(reader io.Reader) (*Config, error) {
	config := &Config{
		Player1: player.Player{},
		Player2: player.Player{},
	}

	scanner := bufio.NewScanner(reader)
	for line := 0; scanner.Scan(); line++ {
		switch line {
		case 0:
			m, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.M = m

		case 1:
			s, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.S = s

		case 2:
			cs, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.Player1.ShipPositions = cs

		case 3:
			cs, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.Player2.ShipPositions = cs

		case 4:
			t, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.T = t

		case 5:
			moves, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.Player1.Moves = moves

		case 6:
			moves, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.Player2.Moves = moves
		}
	}

	return config, nil
}
