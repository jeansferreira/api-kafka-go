package middlewares

import (
	"collect-server/services"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type genericBody struct {
	ApiKey string `json:"apiKey"`
}

func ValidateClientMiddleware(clientsService *services.ClientsService) func(*fiber.Ctx) error {
	reWhitespace := regexp.MustCompile(`\s+`)

	return func(c *fiber.Ctx) error {
		body := genericBody{}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid body",
			})
		}

		client, err := clientsService.GetClient(body.ApiKey)

		if err != nil {

			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid apiKey",
			})
		}

		origin, ok := getOrigin(c.GetReqHeaders())

		if !ok {

			return c.Status(400).JSON(fiber.Map{
				"error": "Missing origin header",
			})
		}

		rawHosts, ok := client.Info["hostWhitelist"]

		if !ok {

			return c.Status(400).JSON(fiber.Map{
				"error": "Client settings does not have hostWhitelist",
			})
		}

		hosts := strings.Split(reWhitespace.ReplaceAllString(rawHosts, ""), ",")
		allowed := false

		for _, host := range hosts {
			if host == origin {
				allowed = true
				break
			}
		}

		if !allowed {

			return c.Status(403).JSON(fiber.Map{
				"error": "Origin not allowed",
			})
		}

		return c.Next()
	}
}

func getOrigin(headers map[string]string) (string, bool) {
	if origin, ok := headers["Origin"]; ok {
		return origin, true
	}

	if origin, ok := headers["origin"]; ok {
		return origin, true
	}

	if origin, ok := headers["X-Host"]; ok {
		return origin, true
	}

	if origin, ok := headers["x-host"]; ok {
		return origin, true
	}

	return "", false
}
