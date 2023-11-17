package main

import (
	"belajar-api-goleng/config"
	"belajar-api-goleng/databases"
	"belajar-api-goleng/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// config env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.InitConfigEnv()

	// databases connection
	databases.InitMysql()

	// aplikasi back end fremwork gin
	app := gin.Default()

	// Router
	routes.RouterIndex(app)

	// port
	app.Run(config.APP_PORT)
}
