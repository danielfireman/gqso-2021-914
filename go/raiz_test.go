package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_raizQuadrada(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		result, err := raizQuadrada(4)
		assert.Nil(t, err)
		assert.Equal(t, 2., result)
	})
	t.Run("ParâmetroInválido", func(t *testing.T) {
		_, err := raizQuadrada(-4)
		assert.Error(t, err)
	})
}
