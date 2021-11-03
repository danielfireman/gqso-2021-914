package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_raizQuadrada(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		result, err := raizQuadrada("4")
		assert.Nil(t, err)
		assert.Equal(t, 2., result)
	})
	t.Run("ParâmetroInválido_NúmeroNegativo", func(t *testing.T) {
		_, err := raizQuadrada("-4")
		assert.Error(t, err)
	})
	t.Run("ParâmetroInválido_Letra", func(t *testing.T) {
		_, err := raizQuadrada("a")
		assert.Error(t, err)
	})
}
