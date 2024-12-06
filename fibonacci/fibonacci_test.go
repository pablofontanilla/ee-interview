package fibonacci_test

import (
	"mywsapp/fibonacci"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIndex(t *testing.T) {
	t.Run("should return the fibonacci index", func(tt *testing.T) {
		i := "9"

		resp, err := fibonacci.ParseIndex(i)
		assert.Nil(tt, err)
		assert.Equal(tt, resp, "34")
	})

	t.Run("should only accept integer inputs", func(tt *testing.T) {
		i := "I am clearly not an integer"

		resp, err := fibonacci.ParseIndex(i)
		assert.NotNil(tt, err)
		assert.Equal(tt, resp, "")
	})
}
