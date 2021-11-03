package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Utilizando o framework fiber.
	// https://github.com/gofiber/fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bem-vindo(a) a calculadora API.")
	})

	app.Get("/raiz/:op", raizQuadradaHandler)

	app.Listen(":8000")
}
