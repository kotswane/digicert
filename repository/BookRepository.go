package repository

import (
	"context"
	"digicert/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	db  *mongo.Database
	ctx context.Context
}

func NewBookRepository(ctx context.Context, db *mongo.Database) IBookRepository {
	return &BookRepository{db: db, ctx: ctx}
}

func (b BookRepository) Create(book model.Book) (model.Book, error) {

	_, err := b.db.Collection("Book").InsertOne(b.ctx, book)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b BookRepository) Read() ([]model.Book, error) {

	filter := bson.D{}
	res, err := b.db.Collection("Book").Find(b.ctx, filter)
	if err != nil {
		return []model.Book{}, err
	}

	var books []model.Book
	if err = res.All(context.TODO(), &books); err != nil {
		return []model.Book{}, err
	}
	return books, nil
}

func (b BookRepository) Update(id primitive.ObjectID, book model.Book) int64 {

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
		return 0
	}

	return res.ModifiedCount
}

func (b BookRepository) Delete(id primitive.ObjectID) int64 {
	//TODO implement me
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	res, err := b.db.Collection("Book").DeleteOne(b.ctx, filter)
	if err != nil {
		return 0
	}

	return res.DeletedCount
}

func (b BookRepository) GetById(id primitive.ObjectID) (model.Book, error) {
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	book := model.Book{}
	err := b.db.Collection("Book").FindOne(b.ctx, filter).Decode(&book)
	if err != nil {
		return model.Book{}, errors.New("not found")
	}

	return book, nil
}

func (b BookRepository) IsbnAvailable(isbn string) bool {
	filter := bson.D{
		primitive.E{Key: "isbn", Value: isbn},
	}
	book := model.Book{}
	err := b.db.Collection("Book").FindOne(b.ctx, filter).Decode(&book)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return true
	}
	return false
}
