package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption     string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageUrl    string             `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	UserId      primitive.ObjectID `json:"userId,omitempty" bson: "userId,omitempty"`
	TimeCreated time.Time          `json:"timeCreated,omitempty" bson:"timeCreated,omitempty"`
}

//create post expects userId from request.body
func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	post.TimeCreated = time.Now()
	id, _ := primitive.ObjectIDFromHex(post.UserId.Hex())
	post.UserId = id
	collection := client.Database("instagram").Collection("post")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func getPostById(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.String(), "/")
	if len(params) == 3 && params[2] != "" {
		id, _ := primitive.ObjectIDFromHex(params[2])
		var post Post
		collection := client.Database("instagram").Collection("post")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(w).Encode(post)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "No id provided" }`))
	}
}

func getPostsOfUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("content-type", "application/json")
		params := strings.Split(r.URL.String(), "/")
		if len(params) == 4 && params[3] != "" {
			id, _ := primitive.ObjectIDFromHex(params[3])
			collection := client.Database("instagram").Collection("post")
			ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

			findOptions := options.Find()
			pageString := r.URL.Query().Get("page")

			if pageString == "" {
				pageString = "1"
			}
			page, _ := strconv.Atoi(pageString) //pagination

			var perPage int64 = 10
			findOptions.SetSkip((int64(page) - 1) * perPage)
			findOptions.SetLimit(perPage)

			cursor, _ := collection.Find(ctx, bson.M{"userid": id}, findOptions)
			var results []Post
			defer cursor.Close(ctx)

			for cursor.Next(ctx) {
				var post Post
				cursor.Decode(&post)
				results = append(results, post)
			}
			// if err := cursor.All(context.TODO(), &results); err != nil {
			// 	panic(err)
			// }
			json.NewEncoder(w).Encode(results)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "No id provided" }`))
		}
	} else {
		fmt.Fprint(w, "method not allowed")
	}

}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	switch r.Method {
	case http.MethodGet:
		getPostById(w, r)
	case http.MethodPost:
		createPost(w, r)
	default:
		fmt.Fprint(w, "method not allowed")
	}
}
