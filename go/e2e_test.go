package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	app := fiber.New()
	addr, err := sobeServidor(app)
	assert.NoError(t, err)

	resp, err := http.Get(addr + "/raiz/4")
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "2.00", string(body))
}

// Necessário para subir o servidor em uma porta que estará disponível.
func sobeServidor(app *fiber.App) (string, error) {
	// Setup listener
	ln, err := net.Listen(fiber.NetworkTCP, ":0")
	if err != nil {
		return "", err
	}
	go func() {
		listen(app, ln)
	}()
	return fmt.Sprintf("http://127.0.0.1:%s", strings.TrimLeft(ln.Addr().String(), "[::]")), nil
}
