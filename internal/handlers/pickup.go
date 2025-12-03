package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/bibashjaprel/devloveapi/internal/data"
)

func PickupLineHandler(category string) gin.HandlerFunc {
    return func(c *gin.Context) {
        cat, line := data.GetRandomLine(category)
        c.JSON(http.StatusOK, gin.H{
            "category": cat,
            "line":     line,
        })
    }
}
