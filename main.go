package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mluna-again/pregunta2/models"
)

func main() {
	router := gin.Default()
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db := models.New(dbpool)

	router.GET("/questions", func(ctx *gin.Context) {
		questions, err := db.GetQuestions(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"questions": questions,
		})
	})

	router.Run()
	log.Println("Hello new world!")
}
