package main

import (
	"log"

	"github.com/alexsasharegan/dotenv"
	"github.com/gin-gonic/gin"
	"github.com/tgbv/learn/server1/internal/routes"
)

func main() {

	// setup dotenv
	error := dotenv.Load("../.env")
	if error != nil {
		log.Fatalf("Error loading .env file: %v", error)
	}

	// setup router
	routerEngine := gin.Default()

	// binds routes to controllers
	routes.Handle(routerEngine)

	// run server
	routerEngine.Run()

}
