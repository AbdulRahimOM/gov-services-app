package main

import (
	"log"

	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/config"
	"github.com/AbdulRahimOM/gov-services-app/admin-api-gateway/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	serviceClients, err := server.InitServiceClients()
	if err != nil {
		log.Fatal("error occured while initializing service clients, error:", err)
	}

	server.InitRoutes(serviceClients, engine)

	err = engine.Run(config.EnvValues.Port)
	if err != nil {
		log.Fatal("error occured while running the server, error:", err)
	}
}
