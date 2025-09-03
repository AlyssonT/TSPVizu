package main

import (
	"github.com/AlyssonT/tsp-visual-backend/configs"
	"github.com/AlyssonT/tsp-visual-backend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	configs.BuildConfigs()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{configs.GetConfigs().FrontendURL}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	server.Use(cors.New((config)))

	server.POST("/solve", handlers.TSP_Solve)

	server.Run()
}
