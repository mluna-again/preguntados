package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mluna-again/pregunta2/models"
)

var db *models.Queries
var dbPool *pgxpool.Pool

func index(ctx *gin.Context) {
		questions, err := allQuestions(ctx)
		questions, err = withAnswers(ctx, questions)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, questions)
}

func create(ctx *gin.Context) {
	var data QuestionData
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	errors, hasErrors := validateQuestionForCreate(data)
	if hasErrors {
		ctx.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	data, err = createQuestion(ctx, data)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	data, err = withAnswersForOne(ctx, data)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

func update(ctx *gin.Context) {
	var data QuestionData
	err := ctx.ShouldBind(&data)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	errors, hasErrors := validateQuestionForUpdate(data)
	if hasErrors {
		ctx.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	data, err = updateQuestion(ctx, data)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	data, err = withAnswersForOne(ctx, data)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func Setup(router *gin.Engine, d *models.Queries, pool *pgxpool.Pool) {
	db = d
	dbPool = pool

	router.GET("/admin/questions", index)
	router.POST("/admin/questions", create)
	router.PATCH("/admin/questions", update)
}
