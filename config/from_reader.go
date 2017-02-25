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
	var p1Ships, p1Moves coordinates.Coordinates
	var p2Ships, p2Moves coordinates.Coordinates
	config := &Config{}

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

			p1Ships = cs

		case 3:
			cs, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			p2Ships = cs

		case 4:
			t, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			config.T = t

		case 5:
			cs, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			p1Moves = cs

		case 6:
			cs, err := coordinates.ParseCoordinates(scanner.Text())
			if err != nil {
				return nil, err
			}

			p2Moves = cs
		}
	}

	config.Player1 = player.NewPlayer(p1Ships, p1Moves)
	config.Player2 = player.NewPlayer(p2Ships, p2Moves)

	return config, nil
}
