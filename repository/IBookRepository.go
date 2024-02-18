package repository

import (
	"digicert/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IBookRepository interface {
	Create(book model.Book) (model.Book, error)
	Read() ([]model.Book, error)
	Update(id primitive.ObjectID, book model.Book) int64
	Delete(id primitive.ObjectID) int64
	GetById(id primitive.ObjectID) (model.Book, error)
	IsbnAvailable(isbn string) bool
}
