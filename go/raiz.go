package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func raizQuadradaHandler(c *fiber.Ctx) error {
	opStr := c.Params("op")
	op, err := strconv.ParseFloat(opStr, 64)

	// Checando parâmetro.
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro Inválido:\":%s\"", opStr))
	}

	// Calculando raiz quadrada.
	raiz, err := raizQuadrada(op)
	if err != nil {
		msg := fmt.Sprintf("Erro ao calcular raiz quadrada: %s", err)
		return c.Status(http.StatusBadRequest).SendString(msg)
	}

	// formatando retorno para ter 2 casas decimais e enviando para o cliente.
	return c.SendString(fmt.Sprintf("%.2f", raiz))
}

func raizQuadrada(op float64) (float64, error) {
	if op < 0 {
		return 0, fmt.Errorf("Operador negativo: %f", op)
	}
	return math.Sqrt(op), nil
}
