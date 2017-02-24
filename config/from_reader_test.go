package config_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sebdah/go-jek-battleship/config"
	"github.com/sebdah/go-jek-battleship/coordinates"
	"github.com/stretchr/testify/assert"
)

func TestFromReader(t *testing.T) {
	tests := map[string]struct {
		input string
		err   error
	}{
		"valid input": {
			input: "5\n5\n1,1:2,0:2,3:3,4:4,3\n0,1:2,3:3,0:3,4:4,1\n5\n0,1:4,3:2,3:3,1:4,1\n0,1:0,0:1,2:2,3:4,3\n",
			err:   nil,
		},
		"bad M": {
			input: "a\n5\n1,1\n1,1\n5\n1,1\n1,1\n",
			err:   &strconv.NumError{},
		},
		"bad S": {
			input: "5\na\n1,1\n1,1\n5\n1,1\n1,1\n",
			err:   &strconv.NumError{},
		},
		"bad player1 ship moves": {
			input: "5\n5\na,b\n1,1\n5\n1,1\n1,1\n",
			err:   coordinates.ErrInvalidInput,
		},
		"bad player2 ship moves": {
			input: "5\n5\n1,1\na,b\n5\n1,1\n1,1\n",
			err:   coordinates.ErrInvalidInput,
		},
		"bad T": {
			input: "5\n5\n1,1\n1,1\na\n1,1\n1,1\n",
			err:   &strconv.NumError{},
		},
		"bad player1 moves": {
			input: "5\n5\n1,1\n1,1\n5\na,b\n1,1\n",
			err:   coordinates.ErrInvalidInput,
		},
		"bad player2 moves": {
			input: "5\n5\n1,1\n1,1\n5\n1,1\na,b\n",
			err:   coordinates.ErrInvalidInput,
		},
	}

	for _, test := range tests {
		_, err := config.FromReader(strings.NewReader(test.input))
		assert.IsType(t, test.err, err)
	}
}
