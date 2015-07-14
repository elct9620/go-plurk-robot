package plurk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BoolToInt(t *testing.T) {
	assert.Equal(t, 1, BoolToInt(true))
	assert.Equal(t, 0, BoolToInt(false))
}
