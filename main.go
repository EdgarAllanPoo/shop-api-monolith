package main

import (
	"log"
	"os"

	"github.com/EdgarAllanPoo/shop-api-monolith/controllers"
	"github.com/EdgarAllanPoo/shop-api-monolith/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, falling back to system environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	database.InitDatabase()

	router := gin.Default()
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	log.Fatal(router.Run(":" + port))
}
