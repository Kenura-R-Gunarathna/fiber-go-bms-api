package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		hexStr := "48656C6C00011234567803E807D00F00"

		frame, err := DecodeDataFrame(hexStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error decoding",
				"err":     err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Successfully decoded",
			"data": fiber.Map{
				"frame":           frame,
				"isChecksumValid": ValidateChecksum(frame),
			},
		})
	})

	log.Fatal(app.Listen(":8080"))
}
