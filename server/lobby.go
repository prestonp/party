package server

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Lobby manages creating and releasing rooms
type Lobby struct {
	mu              sync.Mutex
	MaxGamePlayers  int
	MaxRooms        int
	reservedRooms   map[int]*Room
	unreservedRooms []int
	Rng             *rand.Rand
}

func NewLobby(opts ...LobbyOption) *Lobby {
	lobby := &Lobby{
		MaxGamePlayers: 8,
		MaxRooms:       int(math.Pow(26, 4)), // 4 letter rooms,
		Rng:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// apply functional options
	for _, opt := range opts {
		opt(lobby)
	}

	// generate rooms
	lobby.unreservedRooms = lobby.Rng.Perm(lobby.MaxRooms)
	lobby.reservedRooms = make(map[int]*Room)
	return lobby
}

type LobbyOption func(lobby *Lobby)

func SetMaxRooms(max int) LobbyOption {
	return func(lobby *Lobby) {
		lobby.MaxRooms = max
	}
}

func SetRng(rng *rand.Rand) LobbyOption {
	return func(lobby *Lobby) {
		lobby.Rng = rng
	}
}

var ErrNoRoomsLeft = errors.New("no rooms left")

// Reserve checks out a random, unused room
func (l *Lobby) Reserve() (*Room, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.unreservedRooms) == 0 {
		return nil, ErrNoRoomsLeft
	}

	var id int
	id, l.unreservedRooms = l.unreservedRooms[0], l.unreservedRooms[1:]
	room := Room{ID: id}
	l.reservedRooms[id] = &room
	return &room, nil
}

type ErrRoomNotFound int

func (e ErrRoomNotFound) Error() string {
	return fmt.Sprintf("room %d not found", int(e))
}

// Release marks a room unused
func (l *Lobby) Release(id int) error {
	l.mu.Lock()
	l.mu.Unlock()

	if _, ok := l.reservedRooms[id]; !ok {
		return ErrRoomNotFound(id)
	}
	delete(l.reservedRooms, id)

	l.unreservedRooms = append(l.unreservedRooms, id)
	return nil
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func rtoa(rid, padding int) string {
	out := ""
	base := len(charset)
	for rid > 0 {
		rem := rid % base
		out = string(charset[rem]) + out
		rid /= base
	}
	for len(out) < padding {
		out = string(charset[0]) + out
	}
	return out
}

func (l *Lobby) GetRoom(id int) (*Room, error) {
	room, ok := l.reservedRooms[id]
	if !ok {
		return nil, ErrRoomNotFound(id)
	}
	return room, nil
}
