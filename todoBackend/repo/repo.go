package repo

import (
	"context"
	"todolist/models"
	"todolist/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

func CreateTaskInDB(task models.ToDoList) error {
	_, err := mongodb.ToDoCollection.InsertOne(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTask() (tasks []models.ToDoList, err error) {
	var task models.ToDoList
	cursor, err := mongodb.ToDoCollection.Find(ctx, bson.D{})
	if err != nil {
		defer cursor.Close(ctx)
		return tasks, err
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func SetTaskAsCompleted(id string) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", "done"}}}}
	_, err := mongodb.ToDoCollection.UpdateOne(ctx, filter, update)
	return err
}

func SetTaskAsNotCompleted(id string) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", "not done"}}}}
	_, err := mongodb.ToDoCollection.UpdateOne(ctx, filter, update)
	return err
}

func DeleteSingleTask(id string) error {
	filter := bson.D{{"_id", id}}
	_, err := mongodb.ToDoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// func DeleteAllTasks() (int64, error) {
// 	d, err := mongodb.ToDoCollection.DeleteMany(ctx, bson.D{{}}, nil)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return d.DeletedCount, err
// }
