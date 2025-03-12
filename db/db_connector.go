package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"griffin/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBConnector struct {
    client_URI string
    database string
    collection string
    timeout int
}

func NewDBConnector(client_URI string, database string, collection string, timeout int) DBConnector {
    return DBConnector{client_URI, database, collection, timeout}

}

// Creates a connection to the datbase
func (connector DBConnector) connect() *mongo.Client{
    client, err := mongo.Connect(options.Client().ApplyURI(connector.client_URI))
    if err != nil {
        err := fmt.Sprintf("error creating client: %v", err)
        fmt.Println(err)
        os.Exit(1)
	}
    return client
}

// Closes the connection to the database
func close_connection(client *mongo.Client) {
    if err := client.Disconnect(context.TODO()); err != nil {
        panic(err)
    }
}

// Writes a new task (must not be initialized or writen before as the ObjecID will be saved automatically)
func (connector DBConnector) WriteTask(task *models.Task) string {
    client := connector.connect()
    defer close_connection(client)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(connector.timeout))
    defer cancel()
    collection := client.Database(connector.database).Collection(connector.collection)

    res, err := collection.InsertOne(ctx, task)
    if err != nil {
        panic(err)
    }
    id := res.InsertedID.(bson.ObjectID).Hex()
    task.ID = res.InsertedID.(bson.ObjectID)
    return id
}

// Queries all available tasks
func (connector DBConnector) GetAllTasks() []models.Task {
    client := connector.connect()
    defer close_connection(client)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(connector.timeout))
    defer cancel()
    collection := client.Database(connector.database).Collection(connector.collection)

    cursor, err := collection.Find(ctx, bson.D{})
    if err != nil {
        panic(err)
    }

    defer cursor.Close(ctx)

    total := make([]models.Task, 0)

    for cursor.Next(ctx){
        var result models.Task
        cursor.Decode(&result)
        total = append(total, result)
    }

    return total

}

// Updates a task based in the new object
func (connector DBConnector) UpdateTask(task *models.Task) {
    client := connector.connect()
    defer close_connection(client)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(connector.timeout))
    defer cancel()
    collection := client.Database(connector.database).Collection(connector.collection)

    res := collection.FindOneAndReplace(ctx, bson.D{{Key: "_id", Value: task.ID}}, task)
    res.Decode(task)
}


// Delete a task from the database (it will still remain in the program)
func (connector DBConnector) DeleteTask(task *models.Task) {
    client := connector.connect()
    defer close_connection(client)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(connector.timeout))
    defer cancel()
    collection := client.Database(connector.database).Collection(connector.collection)

    res := collection.FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: task.ID}})
    res.Decode(task)
}
