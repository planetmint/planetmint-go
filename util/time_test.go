package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2UnixTime(t *testing.T) {
	t.Parallel()
	input := "2024-03-26T11:10:41"
	unixTime, err := String2UnixTime(input)
	assert.NoError(t, err)
	assert.Equal(t, int64(1711451441), unixTime)
}
