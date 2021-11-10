package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPotencia_sucesso(t *testing.T) {
	r, err := potencia("2", "3")
	assert.Nil(t, err)
	assert.Equal(t, 8., r)
}

func TestPotencia_erro(t *testing.T) {
	t.Run("baseInválida", func(t *testing.T) {
		_, err := potencia("a", "3")
		assert.Error(t, err)
	})
	t.Run("expoenteInválido", func(t *testing.T) {
		_, err := potencia("2", "b")
		assert.Error(t, err)
	})
}
