package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
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

func TestPotenciaHandler(t *testing.T) {
	app := fiber.New()
	registerHandlers(app)

	t.Run("Ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/potencia/2/3", nil)
		resp, err := app.Test(req, 1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "8.00", string(body))
	})
	t.Run("Error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/potencia/2/aa", nil)
		resp, _ := app.Test(req, 1)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
