package main

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/middleware"
	"github.com/AbdulRahimOM/gov-services-app/user-api-gateway/internal/server"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	// app.Use(middleware.CustomLogger)
	app.Static("/", "./assets")
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	serviceClients, err := server.InitServiceClients()
	if err != nil {
		log.Fatal("error occured while initializing service clients, error:", err)
	}

	server.InitRoutes(serviceClients, app.Use(middleware.CustomLogger))

	err = app.Listen(config.EnvValues.Port)
	if err != nil {
		log.Fatal("error occured while running the server, error:", err)
	}
}
