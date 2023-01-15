package initializers

import (
	"collect-server/controllers"
	"collect-server/env"
	"collect-server/middlewares"
	"collect-server/schemas"
	"collect-server/services"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializeServer() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	prometheus := fiberprometheus.New("event-collect-server")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	clientsService := services.NewClientsService(services.ClientAuth{
		Username: env.PLAT_USER,
		Password: env.PLAT_PASSWORD,
	})

	// V5
	setupV5(app.Get("/v5", middlewares.ValidateClientMiddleware(&clientsService)))

	// V7
	setupV7(app.Get("/v7", middlewares.ValidateClientMiddleware(&clientsService)))

	return app
}

func setupV5(router fiber.Router) {

}

func setupV7(router fiber.Router) {
	kafkaService := services.NewKafkaService(&services.KafkaServiceConfig{
		Brokers: env.KAFKA_BROKERS,
	})

	controller := controllers.NewV7Controller(&kafkaService)

	router.Post(
		"events/views/home",
		middlewares.ValidateBodyMiddleware[schemas.HomeEvent](),
		controller.HandleHome,
	)

	router.Post(
		"events/views/product",
		middlewares.ValidateBodyMiddleware[schemas.ProductEvent](),
		controller.HandleProduct,
	)
}
