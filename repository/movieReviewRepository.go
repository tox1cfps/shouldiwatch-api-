package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tox1cfps/shouldiwatch-api/model"
)

type MovieReviewRepository struct {
	connection *sql.DB
}

func NewMovieReviewRepository(connection *sql.DB) MovieReviewRepository {
	return MovieReviewRepository{
		connection: connection,
	}
}

func (m *MovieReviewRepository) GetReviews() ([]model.MovieReview, error) {
	query := "SELECT id, reviewername, movie, review, rating, isfavorite, created_at FROM reviews"
	rows, err := m.connection.Query(query)
	if err != nil {
		log.Println("erro ao fazer a busca das reviews:", err)
		return []model.MovieReview{}, nil
	}

	defer rows.Close()

	var reviewList []model.MovieReview

	for rows.Next() {
		var reviewObj model.MovieReview
		err = rows.Scan(
			&reviewObj.ID,
			&reviewObj.ReviewerName,
			&reviewObj.Movie,
			&reviewObj.Review,
			&reviewObj.Rating,
			&reviewObj.IsFavorite,
			&reviewObj.Created_At,
		)
		if err != nil {
			log.Println("Erro ao iterar pelas reviews:", err)
			return []model.MovieReview{}, err
		}

		reviewList = append(reviewList, reviewObj)

	}

	return reviewList, nil
}

func (m *MovieReviewRepository) GetReviewByID(id int) (model.MovieReview, error) {
	query := "SELECT id, reviewername, movie, review, rating, isfavorite, created_at FROM reviews WHERE id=$1"
	row := m.connection.QueryRow(query, id)

	var reviewObj model.MovieReview
	err := row.Scan(
		&reviewObj.ID,
		&reviewObj.ReviewerName,
		&reviewObj.Movie,
		&reviewObj.Review,
		&reviewObj.Rating,
		&reviewObj.IsFavorite,
		&reviewObj.Created_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.MovieReview{}, fmt.Errorf("nenhuma review encontrada com o id: %d", id)
		}
	}

	return reviewObj, nil
}

func (m *MovieReviewRepository) CreateReview(review model.MovieReview) (int, error) {
	var id int
	query, err := m.connection.Prepare("INSERT INTO reviews" + "(reviewername, movie, review, rating, isfavorite)" + "VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		log.Println("Erro ao criar nova review:", err)
		return 0, nil
	}

	err = query.QueryRow(review.ReviewerName, review.Movie, review.Review, review.Rating, review.IsFavorite).Scan(&id)
	if err != nil {
		log.Println("erro ao validar criação de review:", err)
		return 0, nil
	}

	query.Close()

	return id, nil
}

func (m *MovieReviewRepository) UpdateReview(review model.MovieReview) error {
	query, err := m.connection.Prepare("UPDATE reviews SET reviewername = $1, movie = $2, review = $3, rating = $4, isfavorite = $5 WHERE id = $6")
	if err != nil {
		log.Println("Erro ao atualizar review:", err)
		return err
	}
	defer query.Close()

	result, err := query.Exec(review.ReviewerName, review.Movie, review.Review, review.Rating, review.IsFavorite, review.ID)
	if err != nil {
		log.Println("erro ao executar atualização:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("erro:", err)
	}

	if rowsAffected == 0 {
		log.Println("Review não encontrada")
	}

	return nil
}

func (m *MovieReviewRepository) DeleteReviews(review model.MovieReview) error {
	query, err := m.connection.Prepare("DELETE FROM reviews WHERE id = $1")
	if err != nil {
		log.Println("erro ao localizar review:", err)
		return err
	}

	defer query.Close()

	result, err := query.Exec(review.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("erro:", err)
	}

	if rowsAffected == 0 {
		log.Println("Review não encontrada")
	}

	return nil
}
