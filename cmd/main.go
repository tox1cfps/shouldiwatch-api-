package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tox1cfps/shouldiwatch-api/config"
	"github.com/tox1cfps/shouldiwatch-api/controller"
	"github.com/tox1cfps/shouldiwatch-api/db"
	"github.com/tox1cfps/shouldiwatch-api/repository"
	"github.com/tox1cfps/shouldiwatch-api/service"
)

func main() {
	// conexão .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Init()
	setting := config.Settings
	dbconnection, err := db.ConnectToDB(
		setting.Database.Host,
		setting.Database.Port,
		setting.Database.User,
		setting.Database.Password,
		setting.Database.Dbname,
		setting.Database.Sslmode,
	)
	if err != nil {
		panic(err)
	}

	// injeções
	movieReviewRepository := repository.NewMovieReviewRepository(dbconnection)
	movieReviewService := service.NewMovieReviewService(movieReviewRepository)
	movieReviewController := controller.NewMovieReviewController(movieReviewService)

	// colocar em outro lugar depois xara
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	})
	r.GET("/reviews", movieReviewController.GetReviews)
	r.GET("/reviews/:id", movieReviewController.GetReviewByID)
	r.POST("/review", movieReviewController.CreateReview)
	r.PUT("/review/:id", movieReviewController.UpdateReview)
	r.DELETE("/review/:id", movieReviewController.DeleteReview)
	r.Run(":5000")
}
