package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:",omitempty" bson:"password,omitempty"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = hash(user.Password) //hash user password
	collection := client.Database("instagram").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.String(), "/")
	if len(params) == 3 && params[2] != "" {
		id, _ := primitive.ObjectIDFromHex(params[2])
		var user User
		collection := client.Database("instagram").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		user.Password = "" //not returning password here
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "No id provided" }`))
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	switch r.Method {
	case http.MethodGet:
		getUserById(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		fmt.Fprint(w, "method not allowed")
	}
}
