package pokemon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Fairy    = 17
	Dark     = 15
	Ghost    = 13
	Dragon   = 14
	Electric = 3
)

func TestEffectivenessAccessTable(t *testing.T) {
	attack := 1
	defense := 4

	assert.Equal(t, 20, Effectiveness[attack][defense])
	assert.Equal(t, 20, Effectiveness[Fairy][Dragon])
}

func TestCover(t *testing.T) {
	fairy := NewNode(nil, Fairy)
	fairyDark := NewNode(fairy, Dark)

	assert.False(t, fairyDark.Cover(Electric), "Electric")
	assert.True(t, fairyDark.Cover(Dragon), "Dragon")
	assert.True(t, fairyDark.Cover(Ghost), "Ghost")
}

func TestExport(t *testing.T) {
	fairy := NewNode(nil, Fairy)
	fairyDark := NewNode(fairy, Dark)
	s := fairyDark.Export()
	assert.Equal(t, 2, len(s))
	assert.Equal(t, "Dark", s[0])
	assert.Equal(t, "Fairy", s[1])
}
