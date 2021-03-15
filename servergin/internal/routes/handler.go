package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tgbv/learn/server1/internal/controllers"
)

// Handle function
// Binds routes to controller
func Handle(engineContext *gin.Engine) {
	// bind routes to controllers
	engineContext.GET("/util/ping", controllers.Ping)
	engineContext.GET("/:target", controllers.GetFile)

}
