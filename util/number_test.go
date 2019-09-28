package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 5), 1)
	assert.Equal(t, Min(22, 5), 5)
}
func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 5), 5)
	assert.Equal(t, Max(22, 5), 22)
}
