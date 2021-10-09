package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleRequests() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUsers)

	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", handlePosts)
	http.HandleFunc("/posts/users/", getPostsOfUser)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

var client *mongo.Client

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb://localhost:27017/"
	clientOptions := options.Client().ApplyURI(uri)
	client, _ = mongo.Connect(ctx, clientOptions)
	handleRequests()
}
