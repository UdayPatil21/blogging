package main

import (
	"log"
	"os"

	"github.com/UdayPatil21/blogging/db"
	"github.com/UdayPatil21/blogging/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func startServer(r *gin.Engine) {
	//init routes
	routes.Init(r)

	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	//init database
	db.Init()
}
func main() {
	//Create gin framework instance
	r := gin.Default()

	//Start server and initialize the data base and routes
	startServer(r)
	r.Run(":" + os.Getenv("ServerPort"))
}
