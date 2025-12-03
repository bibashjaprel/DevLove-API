package server

import (
  "os"
  "github.com/gin-gonic/gin"
  "github.com/bibashjaprel/devloveapi/internal/handlers"
)

func Start() {
    r := gin.Default()

    r.GET("/api/health", func(c *gin.Context) {
      c.JSON(200, gin.H{"status": "ok"})
    })

    r.GET("/api/golang", handlers.PickupLineHandler("golang"))
    r.GET("/api/git", handlers.PickupLineHandler("git"))
    r.GET("/api/docker", handlers.PickupLineHandler("docker"))
    r.GET("/api/kubernetes", handlers.PickupLineHandler("kubernetes"))
    r.GET("/api/romantic", handlers.PickupLineHandler("romantic"))
    r.GET("/api/random", handlers.PickupLineHandler("random"))

    port := os.Getenv("PORT")
    if port == "" {
      port = "7043"
    }
    r.Run(":" + port)
}
