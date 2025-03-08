package routes

import (
    controller "cfw/controllers"
    "cfw/middleware"
    "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.User(middleware.Authenticate())
    incomingRoutes.GET("/users", controller.GetUsers())
    incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
