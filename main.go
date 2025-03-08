package main

import (
    routes "cfw/routes"
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    port = os.Getenv("PORT")
    if port == "" {
        port = "8090"
    }

    router := gin.New()
    router.Use(gin.Logger())

    routes.AuthRoutes(router)
    routes.UserRoutes(router)

    router.GET("/api/v1", func(c *gin.Context){
        c.JSON(200, gin.M{"success": "Access granted for api/v1"})
    })

    router.GET("/api/v2", func(c *gin.Context) {
        c.JSON(200, gin.M{"success": "Access granted for api/v2"})
    })

    router.Run(":" + port)
}
