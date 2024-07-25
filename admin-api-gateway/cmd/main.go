package main

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	serviceClients, err := server.InitServiceClients()
	if err != nil {
		log.Fatal("error occured while initializing service clients, error:", err)
	}

	server.InitRoutes(serviceClients, app)

	err = app.Listen(config.EnvValues.Port)
	if err != nil {
		log.Fatal("error occured while running the server, error:", err)
	}
}
