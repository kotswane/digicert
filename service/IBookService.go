package service

import (
	"digicert/dto"
)

type IBookService interface {
	Create(book dto.CreateBook) (dto.BookResponse, error)
	Read() ([]dto.BookResponse, error)
	Update(id string, book dto.UpdateBook) error
	Delete(id string) error
	GetById(id string) (dto.BookResponse, error)
}
