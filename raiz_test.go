package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_raizQuadrada(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		result, err := raizQuadrada("4")
		assert.NoError(t, err)
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

func TestRaizQuadradaHandler(t *testing.T) {
	app := fiber.New()
	registerHandlers(app)

	t.Run("Ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/raiz/4", nil)
		resp, err := app.Test(req, 1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "2.00", string(body))
	})
	t.Run("Error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/raiz/aa", nil)
		resp, _ := app.Test(req, 1)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
