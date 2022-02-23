package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func potenciaHandler(c *fiber.Ctx) error {
	base := c.Params("base")
	exp := c.Params("exp")

	// Calculando raiz quadrada.
	resultado, err := potencia(base, exp)
	if err != nil {
		msg := fmt.Sprintf("Erro ao calcular potência: %s", err)
		return c.Status(http.StatusBadRequest).SendString(msg)
	}

	// formatando retorno para ter 2 casas decimais e enviando para o cliente.
	return c.SendString(fmt.Sprintf("%.2f", resultado))
}

func potencia(base, expoente string) (float64, error) {
	b, err := strconv.ParseFloat(base, 64)
	if err != nil {
		return 0, fmt.Errorf("base inválida (%s):%q", base, err)
	}
	e, err := strconv.ParseFloat(expoente, 64)
	if err != nil {
		return 0, fmt.Errorf("expoente inválido (%s):%q", base, err)
	}
	return math.Pow(b, e), nil
}
