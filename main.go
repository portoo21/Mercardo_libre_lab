package main

import (
	"fmt"
	"mercado-libre/controllers"
	"mercado-libre/middlewares"
	"mercado-libre/repositories"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
		os.Exit(1)
	}
	clientRepository, err := repositories.InitClientRepository()

	if err != nil {
		fmt.Printf("Bad thing happened! %v", err)
		os.Exit(1)
	}

	controller := controllers.InitClientController(clientRepository)

	defer controller.Close()

	r := gin.Default()

	protected := r.Group("/clients")

	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/", controller.GetClients)
	protected.GET("/sensitive", controller.GetClientSensitive)

	r.Run(":80")

}
