package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidatorError struct {
	Name  string `json:"name"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateBodyMiddleware[T any]() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var data T

		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		v := validator.New()

		if err := v.Struct(data); err != nil {
			var errors []ValidatorError

			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, ValidatorError{
					Name:  err.Field(),
					Tag:   err.Tag(),
					Value: err.Param(),
				})
			}

			return c.Status(400).JSON(fiber.Map{
				"errors": errors,
			})
		}

		return c.Next()
	}
}
