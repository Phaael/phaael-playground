package main

import (
	"log"
	"os"

	routers "github.com/phaael/phaael-playground/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func main() {
	StartServer()
}

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := Run(port); err != nil {
		log.Panic("error running server", err)
	}
}

func Run(port string) error {
	Router = gin.Default()
	routers.MapURL(Router)
	return Router.Run(":" + port)
}
