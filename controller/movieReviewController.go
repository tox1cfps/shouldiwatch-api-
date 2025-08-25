package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tox1cfps/shouldiwatch-api/model"
	"github.com/tox1cfps/shouldiwatch-api/service"
)

type MovieReviewController struct {
	service service.MovieReviewService
}

func NewMovieReviewController(service service.MovieReviewService) MovieReviewController {
	return MovieReviewController{
		service: service,
	}
}

func (mrc *MovieReviewController) GetReviews(ctx *gin.Context) {
	reviews, err := mrc.service.GetReviews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, reviews)
}

func (mrc *MovieReviewController) GetReviewByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	review, err := mrc.service.GetReviewByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ctx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, review)
}

func (mrc *MovieReviewController) CreateReview(ctx *gin.Context) {
	var review model.MovieReview
	err := ctx.BindJSON(&review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertReview, err := mrc.service.CreateReview(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertReview)
}

func (mrc *MovieReviewController) UpdateReview(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var review model.MovieReview
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	review.ID = id

	_, err = mrc.service.UpdateReview(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Review atualizada com sucesso",
		"id":      id,
	})
}

func (mrc *MovieReviewController) DeleteReview(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	review := model.MovieReview{ID: id}

	_, err = mrc.service.DeleteReview(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "review deletada com sucesso!",
	})
}
