package service

import (
	"digicert/dto"
	"digicert/model"
	"digicert/repository"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService struct {
	repo repository.IBookRepository
}

func NewBookService(repo repository.IBookRepository) IBookService {
	return &BookService{repo: repo}
}

func (b BookService) Create(createBook dto.CreateBook) (dto.BookResponse, error) {

	isbnAvailable := b.repo.IsbnAvailable(createBook.ISBN)
	if !isbnAvailable {
		return dto.BookResponse{}, fmt.Errorf("ISBN %s already exists", createBook.ISBN)
	}

	book, err := b.repo.Create(model.Book{
		ID:          primitive.NewObjectID(),
		Author:      createBook.Author,
		Publication: createBook.Publication,
		ISBN:        createBook.ISBN,
	})

	if err != nil {
		return dto.BookResponse{}, err
	}

	return dto.BookResponse{
		ID:          book.ID,
		Author:      book.Author,
		Publication: book.Publication,
		ISBN:        book.ISBN,
	}, nil
}

func (b BookService) Read() ([]dto.BookResponse, error) {
	books, err := b.repo.Read()
	var booksResponse []dto.BookResponse

	if err != nil {
		return booksResponse, err
	}

	for _, book := range books {
		booksResponse = append(booksResponse, dto.BookResponse{
			ID:          book.ID,
			Author:      book.Author,
			Publication: book.Publication,
			ISBN:        book.ISBN,
		})
	}

	if len(booksResponse) == 0 {
		return booksResponse, errors.New("not found")
	}

	return booksResponse, nil
}

func (b BookService) Update(id string, updateBook dto.UpdateBook) error {
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	modifiedCount := b.repo.Update(bookId, model.Book{
		Author:      updateBook.Author,
		Publication: updateBook.Publication,
	})

	if modifiedCount == 0 {
		return errors.New("not found")
	}
	return nil
}

func (b BookService) Delete(id string) error {
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	deletedRows := b.repo.Delete(bookId)
	if deletedRows == 0 {
		return errors.New("not found")
	}
	return nil
}

func (b BookService) GetById(id string) (dto.BookResponse, error) {

	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.BookResponse{}, err
	}

	book, err := b.repo.GetById(bookId)
	if err != nil {
		return dto.BookResponse{}, err
	}

	return dto.BookResponse{
		ID:          book.ID,
		Author:      book.Author,
		Publication: book.Publication,
		ISBN:        book.ISBN,
	}, err
}
