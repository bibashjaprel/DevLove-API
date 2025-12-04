package server

import (
  "os"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "github.com/bibashjaprel/devloveapi/internal/handlers"
)

func Start() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
      AllowOrigins:     []string{"https://devlove.vercel.app", "https://devlove.bibash.info.np", "http://localhost:3000"},
      AllowMethods:     []string{"GET", "POST", "OPTIONS"},
      AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
      ExposeHeaders:    []string{"Content-Length"},
      AllowCredentials: true,
    }))

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
