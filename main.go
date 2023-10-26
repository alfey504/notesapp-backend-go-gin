package main

import (
	"fmt"
	"notes-appapi/config"
	"notes-appapi/middleware"
	"notes-appapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("REST API for Notes App")

	config.ConnectMongo()
	defer config.MongoDisconnect(config.DB)

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	routes.UserRoutes(r)

	r.Run()
	return
}
