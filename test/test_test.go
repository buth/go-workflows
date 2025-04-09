package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	assert.True(t, true, "This test should always pass")
	assert.False(t, false, "This test should also always pass")
}
