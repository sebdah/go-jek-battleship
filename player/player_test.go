package player_test

import (
	"testing"

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
