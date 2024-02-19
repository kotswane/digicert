package service

import (
	"digicert/dto"
	"digicert/model"
	"digicert/repository"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService struct {
	repo         repository.IBookRepository
	customLogger *log.Logger
}

func NewBookService(repo repository.IBookRepository, customLogger *log.Logger) IBookService {
	return &BookService{repo: repo, customLogger: customLogger}
}

func (b BookService) Create(createBook dto.CreateBook) (dto.BookResponse, error) {
	b.customLogger.Infof("In a BookService::Creating a book %s", createBook)
	isbnAvailable := b.repo.IsbnAvailable(createBook.ISBN)
	if !isbnAvailable {
		b.customLogger.Errorf("In a BookService::Creating a book ISBN %s already exists", createBook.ISBN)
		return dto.BookResponse{}, fmt.Errorf("ISBN %s already exists", createBook.ISBN)
	}

	book, err := b.repo.Create(model.Book{
		ID:          primitive.NewObjectID(),
		Author:      createBook.Author,
		Publication: createBook.Publication,
		ISBN:        createBook.ISBN,
	})

	if err != nil {
		b.customLogger.Errorf("In a BookService::Creating a book %s", err.Error())
		return dto.BookResponse{}, err
	}
	b.customLogger.Info("In a BookService::Creating a book success")
	return dto.BookResponse{
		ID:          book.ID,
		Author:      book.Author,
		Publication: book.Publication,
		ISBN:        book.ISBN,
	}, nil
}

func (b BookService) Read() ([]dto.BookResponse, error) {
	b.customLogger.Info("In a BookService::Listing all the books")
	books, err := b.repo.Read()
	var booksResponse []dto.BookResponse

	if err != nil {
		b.customLogger.Errorf("In a BookService::Listing all the books %s", err.Error())
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
		b.customLogger.Error("In a BookService::Listing all the books not found")
		return booksResponse, errors.New("not found")
	}

	b.customLogger.Infof("In a BookService::Listing all the books %s", booksResponse)
	return booksResponse, nil
}

func (b BookService) Update(id string, updateBook dto.UpdateBook) error {
	b.customLogger.Infof("In a BookService::updating a book with id %s", id)
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		b.customLogger.Errorf("In a BookService::updating a book %s", err.Error())
		return err
	}

	modifiedCount := b.repo.Update(bookId, model.Book{
		Author:      updateBook.Author,
		Publication: updateBook.Publication,
	})

	if modifiedCount == 0 {
		b.customLogger.Errorf("In a BookService::updating a book no modified record")
		return errors.New("not found")
	}
	b.customLogger.Info("In a BookService::updating a book success")
	return nil
}

func (b BookService) Delete(id string) error {
	b.customLogger.Infof("In a BookService::deleting a book with id %s", id)
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		b.customLogger.Errorf("In a BookService::deleting a book %s", err.Error())
		return err
	}

	deletedRows := b.repo.Delete(bookId)
	if deletedRows == 0 {
		b.customLogger.Errorf("In a BookService::deleting a book not found")
		return errors.New("not found")
	}
	b.customLogger.Info("In a BookService::deleting a book success")
	return nil
}

func (b BookService) GetById(id string) (dto.BookResponse, error) {
	b.customLogger.Infof("In a BookService::getbyid %s", id)
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		b.customLogger.Errorf("In a BookService::getbyid %s", err.Error())
		return dto.BookResponse{}, err
	}

	book, err := b.repo.GetById(bookId)
	if err != nil {
		b.customLogger.Errorf("In a BookService::getbyid %s", err.Error())
		return dto.BookResponse{}, err
	}

	b.customLogger.Infof("In a BookService::getbyid success %s", book)
	return dto.BookResponse{
		ID:          book.ID,
		Author:      book.Author,
		Publication: book.Publication,
		ISBN:        book.ISBN,
	}, err
}
