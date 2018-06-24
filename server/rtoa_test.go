package server

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRtoa(t *testing.T) {
	assert.Equal(t, "AABA", rtoa(26, 4))
	assert.Equal(t, "AABB", rtoa(27, 4))
	assert.Equal(t, "AAAA", rtoa(0, 4))
	assert.Equal(t, "AAAB", rtoa(1, 4))
	assert.Equal(t, "AAAC", rtoa(2, 4))
	assert.Equal(t, "AAAD", rtoa(3, 4))
	assert.Equal(t, "ABVM", rtoa(1234, 4))
	assert.Equal(t, "ASGV", rtoa(12345, 4))
	assert.Equal(t, "ZZZZ", rtoa(int(math.Pow(26, 4))-1, 4))
}
