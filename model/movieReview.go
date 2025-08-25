package model

import "time"

type MovieReview struct {
	ID           int       `json:"id"`
	Rating       int       `json:"rating"`
	ReviewerName string    `json:"reviewername"`
	Movie        string    `json:"movie"`
	Review       string    `json:"review"`
	IsFavorite   bool      `json:"isfavorite"`
	Created_At   time.Time `json:"created_at"`
}
