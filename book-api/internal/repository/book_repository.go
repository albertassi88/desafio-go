package repository

import (
	"book-api/internal/models"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func InitMongoDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI não está definida!")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("bookdb").Collection("books")
}

func GetBookByID(id string) (models.Book, error) {
	var book models.Book

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, fmt.Errorf("invalid ObjectID format")
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&book)
	if err != nil {
		return book, err
	}
	return book, nil
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

func UpdateBook(id string, book models.Book) (int64, error) {
	fmt.Println("Updating book with ID:", id)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	update := bson.M{
		"$set": bson.M{
			"title":    book.Title,
			"category": book.Category,
			"author":   book.Author,
			"synopsis": book.Synopsis,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func DeleteBook(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
