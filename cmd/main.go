package main

import (
	"log"
	"net/http"

	"github.com/fhva29/go-vault/internal/config"
	"github.com/fhva29/go-vault/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configs and db
	config.LoadEnv()
	_ = db.Connect()

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
