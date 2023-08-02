package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func CreateShort(rw http.ResponseWriter, r *http.Request) error {
	var reqBody CreateShortURLReq
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return err
	}

	// generate hash

	urlRecord := URLRecord{
		Hash : "some_value",
		LongUrl : reqBody.LongUrl,
	}

	responseJSON, err := json.Marshal(urlRecord)
	if err != nil {
		return err
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(responseJSON)

	// log.Println(urlRecord)

	return nil
}

func ServeShort(rw http.ResponseWriter, r *http.Request) error {
	return nil
}

func Handler(f apiFunc) http.HandlerFunc {
	log.Println(f)
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := f(rw, r); err != nil {
			log.Println(err.Error())
		}
	}
}


type Server struct {
	Address string
	Router  *mux.Router
	DB      *MongoDB
}

func NewServer() *Server {
	router := mux.NewRouter()
	router.HandleFunc("/create-short-url", Handler(CreateShort)).Methods("POST")
	router.HandleFunc("/{shortUrl}", Handler(ServeShort)).Methods("GET")
	return &Server{
		Address: ":8080",
		Router:  router,
	}
}

func (s Server) Run() {
	log.Println("Server Runnning : ", s)
	log.Fatal(http.ListenAndServe(s.Address, s.Router))
}

func (s *Server) initDB() error {
	log.Println("init DB ... ")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// credential := options.Credential{
	// 	Username: "user",
	// 	Password: "password",
	// }
	// clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	// if err := createDatabaseAndCollection(ctx, client, "url_shortner_db", "urlrecords"); err != nil {
	// 	return err
	// }

	s.DB = &MongoDB{client: client}
	return nil
}


type Storage interface {
	GetLongUrl(string) (string, error)
	GetHashValue(string) (string, error)
	UpdateHashValue(*URLRecord) error
	DeleteURLRecord(string) error
}

type MongoDB struct {
	client *mongo.Client
}




func createDatabaseAndCollection(ctx context.Context, client *mongo.Client, dbName, collectionName string) error {
	// Check if the desired collection already exists.
	// collections, err := client.Database(dbName).Collection(collectionName)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return err
	// }

	// If the collection doesn't exist, create it.
	collectionExists := false
	// for _, coll := range collections {
	// 	if coll == collectionName {
	// 		collectionExists = true
	// 		break
	// 	}
	// }

	// If the collection doesn't exist, create it.
	if !collectionExists {
		err := client.Database(dbName).CreateCollection(ctx, collectionName)
		if err != nil {
			return err
		}
		fmt.Printf("Collection \"%s\" created successfully in database \"%s\".\n", collectionName, dbName)
	}

	return nil
}

type URLRecord struct {
	Hash    string	`json:"hash"`
	LongUrl string	`json:"long_url"`
}

type CreateShortURLReq struct {
	LongUrl string `json:"long_url"`
}


func main() {
	apiServer := NewServer()

	if err := apiServer.initDB(); err != nil {
		log.Fatal(err)
	}
	
	apiServer.Run()
}
