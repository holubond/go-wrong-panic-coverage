package temp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert.Panics(t, func() { foo() })
	assert.Equal(t, 0, shouldBeZero)
}
