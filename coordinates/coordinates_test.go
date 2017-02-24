package coordinates_test

import (
	"testing"

	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/stretchr/testify/assert"
)

func TestCoordinateEquals(t *testing.T) {
	tests := []struct {
		c1    *coordinates.Coordinate
		c2    *coordinates.Coordinate
		equal bool
	}{
		{
			c1:    coordinates.NewCoordinate(1, 1),
			c2:    coordinates.NewCoordinate(1, 1),
			equal: true,
		},
		{
			c1:    coordinates.NewCoordinate(1, 1),
			c2:    coordinates.NewCoordinate(2, 1),
			equal: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.equal, test.c1.Equal(test.c2))
	}
}

func TestCoordinatesInclude(t *testing.T) {
	tests := []struct {
		cs      coordinates.Coordinates
		c       *coordinates.Coordinate
		include bool
	}{
		{
			cs: coordinates.Coordinates{
				*coordinates.NewCoordinate(1, 1),
				*coordinates.NewCoordinate(1, 2),
			},
			c:       coordinates.NewCoordinate(1, 1),
			include: true,
		},
		{
			cs: coordinates.Coordinates{
				*coordinates.NewCoordinate(1, 1),
				*coordinates.NewCoordinate(1, 2),
			},
			c:       coordinates.NewCoordinate(2, 1),
			include: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.include, test.cs.Include(test.c))
	}
}

func TestParseCoordinates(t *testing.T) {
	tests := map[string]struct {
		input string
		err   error
	}{
		"basic valid input ": {
			input: "1,2:2,3:3,4",
			err:   nil,
		},
		"invalid input": {
			input: "12:2,3:3,4",
			err:   coordinates.ErrInvalidInput,
		},
		"x is not an integer": {
			input: "a,2:2,3:3,4",
			err:   coordinates.ErrInvalidInput,
		},
		"y is not an integer": {
			input: "1,b:2,3:3,4",
			err:   coordinates.ErrInvalidInput,
		},
	}

	for _, test := range tests {
		b, err := coordinates.ParseCoordinates(test.input)
		assert.Equal(t, test.err, err)

		if err == nil {
			assert.Equal(t, test.input, b.String())
		}
	}
}
