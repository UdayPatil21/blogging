package main

import (
	"github.com/UdayPatil21/blogging/db"
	"github.com/UdayPatil21/blogging/routes"
	"github.com/gin-gonic/gin"
)

func startServer(r *gin.Engine) {
	//init routes
	routes.Init(r)

	//init database
	db.Init()
}
func main() {
	//Create gin framework instance
	r := gin.Default()

	//Start server and initialize the data base and routes
	startServer(r)
	r.Run(":3000")
}
