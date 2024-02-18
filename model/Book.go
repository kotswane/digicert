package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Author      string             `bson:"author"`
	Publication string             `bson:"publication"`
	ISBN        string             `bson:"isbn"`
}
