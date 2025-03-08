package database

import (
   "fmt"
   "log"
   "time"
   "os"
   "context"
   "github.com/joho/godotenv"
   "go.mongo.org/mongo-driver/mongo"
   "go/mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Unable to load .env file")
    }

    MongoDB := os.Getenv("MONGODB_URL")

    client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.withTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB.")

    return client
}

var Client *mongo.client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    var collection *mongo.Collection = client.Database("cluster0").Colletion(collectionName)
    return collection
}



