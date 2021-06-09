package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type Url struct {
	Id       string `json:"id" binding:"required"`
	LongUrl  string `json:"long_url" binding:"required"`
	ShortUrl string `json:"short_url" binding:"required"`
}

var db *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	errorHandler(err)

	err = client.Ping(ctx, nil)
	errorHandler(err)

	db = client.Database("urlShorterDB").Collection("urls")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", MainPageHandler)
	router.HandleFunc("/create-short-url", LongUrlHandler).Methods("POST")
	router.HandleFunc("/{shortUrl}", ShortUrlHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func LongUrlHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var uid string
	cnt, _ := db.CountDocuments(ctx, bson.M{"longurl": r.FormValue("url")})

	if cnt != 0 {
		var url []Url
		cursor, err := db.Find(ctx, bson.M{"longurl": r.FormValue("url")})
		errorHandler(err)

		err = cursor.All(ctx, &url)
		errorHandler(err)

		fmt.Println(r.FormValue("url"))
		json.NewEncoder(w).Encode(url[0])
		return
	}

	for {
		uid = uuid.New().String()[:5]
		fmt.Println(uid)
		cnt, _ := db.CountDocuments(ctx, bson.M{"id": uid})
		if cnt == 0 {
			break
		}
	}

	host := "127.0.0.1:8081/"
	url := Url{
		Id:       uid,
		LongUrl:  r.FormValue("url"),
		ShortUrl: host + uid,
	}

	res, err := db.InsertOne(ctx, url)
	errorHandler(err)

	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
	json.NewEncoder(w).Encode(url)
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["shortUrl"]

	cnt, err := db.CountDocuments(ctx, bson.M{"id": urlId})
	errorHandler(err)

	if cnt == 0 {
		return
	}

	var url []Url
	cursor, err := db.Find(ctx, bson.M{"id": urlId})
	errorHandler(err)

	err = cursor.All(ctx, &url)
	errorHandler(err)

	http.Redirect(w, r, url[0].LongUrl, 307)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
