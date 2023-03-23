package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mluna-again/pregunta2/admin"
	"github.com/mluna-again/pregunta2/models"
)

func main() {
	router := gin.Default()
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db := models.New(dbpool)

	admin.Setup(router, db)

	router.Run()
	log.Println("Hello new world!")
}
