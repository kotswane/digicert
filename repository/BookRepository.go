package repository

import (
	"context"
	"digicert/model"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	db           *mongo.Database
	ctx          context.Context
	customLogger *log.Logger
}

func NewBookRepository(ctx context.Context, db *mongo.Database, customLogger *log.Logger) IBookRepository {
	return &BookRepository{db: db, ctx: ctx, customLogger: customLogger}
}

func (b BookRepository) Create(book model.Book) (model.Book, error) {
	b.customLogger.Infof("In a BookRepository::Creating a book %s", book)
	_, err := b.db.Collection("Book").InsertOne(b.ctx, book)
	if err != nil {
		b.customLogger.Errorf("In a BookRepository::Creating a book %s", err.Error())
		return model.Book{}, err
	}
	b.customLogger.Infof("In a BookRepository::Creating a book success %s", book)
	return book, nil
}

func (b BookRepository) Read() ([]model.Book, error) {
	b.customLogger.Info("In a BookRepository::Listing all the books")
	filter := bson.D{}
	res, err := b.db.Collection("Book").Find(b.ctx, filter)
	if err != nil {
		b.customLogger.Errorf("In a BookRepository::Listing all the books %s", err.Error())
		return []model.Book{}, err
	}

	var books []model.Book
	if err = res.All(context.TODO(), &books); err != nil {
		b.customLogger.Errorf("In a BookRepository::Listing all the books %s", err.Error())
		return []model.Book{}, err
	}
	b.customLogger.Infof("In a BookRepository::Listing all the books %s", books)
	return books, nil
}

func (b BookRepository) Update(id primitive.ObjectID, book model.Book) int64 {
	b.customLogger.Infof("In a BookRepository::updating a book with id %s", id)
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}

	update := bson.M{
		"$set": bson.M{
			"author":      book.Author,
			"publication": book.Publication,
		},
	}

	res, err := b.db.Collection("Book").UpdateOne(b.ctx, filter, update)
	if err != nil {
		b.customLogger.Errorf("In a BookRepository::updating a book no modified record")
		return 0
	}
	b.customLogger.Info("In a BookRepository::updating a book success")
	return res.ModifiedCount
}

func (b BookRepository) Delete(id primitive.ObjectID) int64 {
	b.customLogger.Infof("In a BookRepository::deleting a book with id %s", id)
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	res, err := b.db.Collection("Book").DeleteOne(b.ctx, filter)
	if err != nil {
		b.customLogger.Errorf("In a BookService::deleting a book %s", err.Error())
		return 0
	}
	b.customLogger.Info("In a BookService::deleting a book success")
	return res.DeletedCount
}

func (b BookRepository) GetById(id primitive.ObjectID) (model.Book, error) {
	b.customLogger.Infof("In a BookRepository::getbyid %s", id)
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	book := model.Book{}
	err := b.db.Collection("Book").FindOne(b.ctx, filter).Decode(&book)
	if err != nil {
		b.customLogger.Errorf("In a BookRepository::getbyid %s", err.Error())
		return model.Book{}, errors.New("not found")
	}
	b.customLogger.Infof("In a BookRepository::getbyid success %s", book)
	return book, nil
}

func (b BookRepository) IsbnAvailable(isbn string) bool {
	b.customLogger.Infof("In a BookRepository::isbnAvailable %s", isbn)
	filter := bson.D{
		primitive.E{Key: "isbn", Value: isbn},
	}
	book := model.Book{}
	err := b.db.Collection("Book").FindOne(b.ctx, filter).Decode(&book)
	if errors.Is(err, mongo.ErrNoDocuments) {
		b.customLogger.Errorf("In a BookRepository::isbnAvailable %s", err.Error())
		return true
	}
	b.customLogger.Info("In a BookRepository::isbnAvailable success")
	return false
}
