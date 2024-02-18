package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateBook model info
// @Description Creates a book
type CreateBook struct {
	Author      string `json:"author" binding:"required"`
	Publication string `json:"publication" binding:"required"`
	ISBN        string `isbn:"isbn" binding:"required"`
}

// UpdateBook model info
// @Description Updates a book
type UpdateBook struct {
	Author      string `json:"author" binding:"required"`
	Publication string `json:"publication" binding:"required"`
}

// BookResponse model info
// @Description Book response
type BookResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Author      string             `json:"author"`
	Publication string             `json:"publication"`
	ISBN        string             `isbn:"isbn"`
}
