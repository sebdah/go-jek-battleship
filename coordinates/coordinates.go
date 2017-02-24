package coordinates

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrInvalidInput is returned if the input of coordinates is invalid.
	ErrInvalidInput = errors.New("Invalid coordinate input")
)

// Coordinate represents an coordinate.
type Coordinate struct {
	x int
	y int
}

// NewCoordinate instantiates a new coordinate.
func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{
		x: x,
		y: y,
	}
}

// Equal checks if the coordinate is equal to the input coordinate.
func (c *Coordinate) Equal(c2 *Coordinate) bool {
	return c.x == c2.x && c.y == c2.y
}

// String returns the string representation of the coordinate, x:y.
func (c *Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

// Coordinates represents a slice of coordinates.
type Coordinates []Coordinate

// Include check if the input coordinate is included in the coordinates slice.
func (c *Coordinates) Include(c1 *Coordinate) bool {
	for _, coordinate := range *c {
		if coordinate.Equal(c1) {
			return true
		}
	}

	return false
}

// String returns the string representation of the coordinates. Format
// x1,y1:x2:y2.
func (c *Coordinates) String() string {
	var coordinates []string
	for _, coodinate := range *c {
		coordinates = append(coordinates, coodinate.String())
	}

	return strings.Join(coordinates, ":")
}

// ParseCoordinates is reading a string of the format x1,y1:x2,y2 and converts
// that to a slice of coordinates.
func ParseCoordinates(input string) (Coordinates, error) {
	var coordinates Coordinates

	for _, coordinateString := range strings.Split(input, ":") {
		coordinateString = strings.TrimSuffix(coordinateString, "\n")
		cs := strings.SplitN(coordinateString, ",", 2)

		if len(cs) != 2 {
			return []Coordinate{}, ErrInvalidInput
		}

		x, err := strconv.Atoi(cs[0])
		if err != nil {
			return []Coordinate{}, ErrInvalidInput
		}

		y, err := strconv.Atoi(cs[1])
		if err != nil {
			return []Coordinate{}, ErrInvalidInput
		}

		coordinates = append(coordinates, Coordinate{
			x: x,
			y: y,
		})
	}

	return coordinates, nil
}
