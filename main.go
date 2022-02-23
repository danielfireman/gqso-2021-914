package main

import (
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Utilizando o framework fiber.
	// https://github.com/gofiber/fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bem-vindo(a) a calculadora API.")
	})

	// Cria um listener padrão de produção. Importante para configuração
	// de parâmetros específicos, bem como facilitar testes fim-a-fim.
	ln, err := net.Listen(fiber.NetworkTCP, ":8080")
	if err != nil {
		log.Fatal(err)
	}
	listen(app, ln)
}

func listen(app *fiber.App, ln net.Listener) error {
	registerHandlers(app)
	return app.Listener(ln)
}

func registerHandlers(app *fiber.App) {
	app.Get("/raiz/:op", raizQuadradaHandler)
	app.Get("/potencia/:base/:exp", potenciaHandler)
}
