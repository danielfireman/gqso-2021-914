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

	// Calculando raiz quadrada.
	raiz, err := raizQuadrada(opStr)
	if err != nil {
		msg := fmt.Sprintf("Erro ao calcular raiz quadrada: %s", err)
		return c.Status(http.StatusBadRequest).SendString(msg)
	}

	// formatando retorno para ter 2 casas decimais e enviando para o cliente.
	return c.SendString(fmt.Sprintf("%.2f", raiz))
}

func raizQuadrada(opStr string) (float64, error) {
	op, err := strconv.ParseFloat(opStr, 64)
	if err != nil {
		return 0, fmt.Errorf("parâmetro não é um número: %s", opStr)
	}
	if op < 0 {
		return 0, fmt.Errorf("operador negativo: %f", op)
	}
	return math.Sqrt(op), nil
}
