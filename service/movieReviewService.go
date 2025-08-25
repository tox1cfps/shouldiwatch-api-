package service

import (
	"log"

	"github.com/tox1cfps/shouldiwatch-api/model"
	"github.com/tox1cfps/shouldiwatch-api/repository"
)

type MovieReviewService struct {
	repository repository.MovieReviewRepository
}

func NewMovieReviewService(repo repository.MovieReviewRepository) MovieReviewService {
	return MovieReviewService{
		repository: repo,
	}
}

func (mr MovieReviewService) GetReviews() ([]model.MovieReview, error) {
	return mr.repository.GetReviews()
}

func (mr MovieReviewService) GetReviewByID(id int) (model.MovieReview, error) {
	return mr.repository.GetReviewByID(id)
}

func (mr MovieReviewService) CreateReview(review model.MovieReview) (model.MovieReview, error) {
	reviewID, err := mr.repository.CreateReview(review)
	if err != nil {
		log.Println(err)
		return model.MovieReview{}, nil
	}

	review.ID = reviewID

	return review, nil
}

func (mr MovieReviewService) UpdateReview(review model.MovieReview) (model.MovieReview, error) {
	err := mr.repository.UpdateReview(review)
	if err != nil {
		log.Println(err)
		return model.MovieReview{}, err
	}

	return review, nil
}

func (mr MovieReviewService) DeleteReview(review model.MovieReview) (model.MovieReview, error) {
	err := mr.repository.DeleteReviews(review)
	if err != nil {
		log.Println(err)
		return model.MovieReview{}, err
	}

	return review, nil
}
