package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priyanshu360/urlShortnerApp.git/config"
	"github.com/priyanshu360/urlShortnerApp.git/handlers"
	"github.com/priyanshu360/urlShortnerApp.git/storage"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func GetHandler(f apiFunc) http.HandlerFunc {
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
	Handler handlers.Handler
}



func (s Server) Run() {
	log.Println("Server Runnning : ", s)
	log.Fatal(http.ListenAndServe(s.Address, s.Router))
}

func NewServer() *Server {
	router := mux.NewRouter()
	server := &Server{
		Address: ":" + config.PORT,
	}
	router.HandleFunc("/create-short-url", GetHandler(server.Handler.CreateShort)).Methods("POST")
	router.HandleFunc("/{shortUrl}", GetHandler(server.Handler.ServeShort)).Methods("GET")
	
	server.Router = router
	return server
}



func main() {
	apiServer := NewServer()

	mongo, err := storage.InitMongo()
	if err != nil {
		log.Fatal("Error initializing MongoDB:", err)
	}

	apiServer.Handler.DB = mongo
	apiServer.Run()
}
