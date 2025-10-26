package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    api := r.Group("/api/v1")
    api.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    addr := ":8080"
    log.Printf("DevSkill Backend listening on %s", addr)
    if err := r.Run(addr); err != nil {
        log.Fatal(err)
    }
}
