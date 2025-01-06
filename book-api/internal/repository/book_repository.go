package repository

import (
	"book-api/internal/models"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func InitMongoDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI não está definida!")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("bookdb").Collection("books")
}

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	result, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
