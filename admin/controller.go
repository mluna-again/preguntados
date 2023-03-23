package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mluna-again/pregunta2/models"
)

var db *models.Queries

func index(ctx *gin.Context) {
		questions, err := allQuestions(ctx)
		questions, err = withAnswers(ctx, questions)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"questions": questions,
		})
}

func Setup(router *gin.Engine, d *models.Queries) {
	db = d

	router.GET("/questions", index)
}
