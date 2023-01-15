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

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/rizalgowandy/go-swag-sample/docs/fibersimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
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

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", HealthCheck)

	// V5
	setupV5(app.Get("/v5", middlewares.ValidateClientMiddleware(&clientsService)))

	// V7
	setupV7(app.Get("/v7", middlewares.ValidateClientMiddleware(&clientsService)))

	app.Get("/swagger/*", swagger.HandlerDefault)

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

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
