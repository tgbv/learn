package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tgbv/learn/server1/internal/routes"
)

func main() {

	// setup router
	routerEngine := gin.Default()

	// binds routes to controllers
	routes.Handle(routerEngine)

	// run server
	routerEngine.Run()
}
