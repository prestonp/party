package server

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLobbyOption(t *testing.T) {
	l := NewLobby(SetMaxRooms(42))
	assert.Equal(t, 42, l.MaxRooms)
}

func TestLobbyReserve(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	l := NewLobby(SetRng(rng))

	room, err := l.Reserve()
	assert.NoError(t, err)

	assert.Equal(t, 317516, room.ID)

	_, err = l.GetRoom(room.ID)
	assert.NoError(t, err)

	err = l.Release(room.ID)
	assert.NoError(t, err)

	_, err = l.GetRoom(room.ID)
	assert.NotEmpty(t, err)
}
