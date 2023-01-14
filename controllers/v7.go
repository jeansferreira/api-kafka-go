package controllers

import (
	"collect-server/schemas"
	"collect-server/services"
	"collect-server/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

const (
	TOPIC_HOME    string = "event_home"
	TOPIC_PRODUCT string = "event_product"
)

type V7Controller struct {
	kafka *services.KafkaService
}

func NewV7Controller(kafka *services.KafkaService) V7Controller {
	return V7Controller{
		kafka: kafka,
	}
}

func (controller *V7Controller) HandleHome(c *fiber.Ctx) error {
	data := schemas.HomeEvent{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	go func() {
		if err := controller.kafka.Publish(TOPIC_HOME, []kafka.Message{
			{
				Value: utils.JsonToByes(data),
			},
		}); err != nil {
			log.Println("Error publishing message to kafka")
			log.Println(err)
		}
	}()

	return c.Status(204).JSON(fiber.Map{})
}

func (controller *V7Controller) HandleProduct(c *fiber.Ctx) error {
	data := schemas.ProductEvent{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	go func() {
		if err := controller.kafka.Publish(TOPIC_PRODUCT, []kafka.Message{
			{
				Value: utils.JsonToByes(data),
			},
		}); err != nil {
			log.Println("Error publishing message to kafka")
			log.Println(err)
		}
	}()

	return c.Status(204).JSON(fiber.Map{})
}
