package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const databaseName = "todoList"

var ToDoCollection *mongo.Collection

func GetMongodbEngine(mongodbURL string) (err error) {
	// credential := options.Credential{
	// 	AuthSource: "admin",
	// 	Username:   "root",
	// 	Password:   "secret",
	// }
	clientOptions := options.Client().ApplyURI(mongodbURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")

	mDB := client.Database(databaseName)
	if mDB == nil {
		return errors.New("database not found")
	}
	ToDoCollection = mDB.Collection("task")

	return nil
}
