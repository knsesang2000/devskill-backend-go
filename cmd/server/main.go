package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    handlers "github.com/knsesang2000/devskill-backend-go/internal/handlers"
)

func main() {
    r := gin.Default()

    // Top-level health
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // Wire routes under /api/v1 (problems, submissions)
    handlers.Router(r)

    addr := ":8080"
    log.Printf("DevSkill Backend listening on %s", addr)
    if err := r.Run(addr); err != nil {
        log.Fatal(err)
    }
}
