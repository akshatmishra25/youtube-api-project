package db

import (
    "context"
    "log"
    "time"

    "youtube-fam/config"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI(config.AppConfig.MongoDBURI)
    var err error
    client, err = mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")

    createCollection("videos")
}

func GetCollection(collectionName string) *mongo.Collection {
    return client.Database(config.AppConfig.DatabaseName).Collection(collectionName)
}

func createCollection(collectionName string) {
    collection := client.Database(config.AppConfig.DatabaseName).Collection(collectionName)
    log.Printf("Collection %s is ready", collection.Name())
}