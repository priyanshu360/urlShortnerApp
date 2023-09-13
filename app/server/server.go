package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priyanshu360/urlShortnerApp.git/config"
	"github.com/priyanshu360/urlShortnerApp.git/controller"
	"github.com/priyanshu360/urlShortnerApp.git/utils"
)

type Server struct {
	Address string
	Router  *mux.Router
	Controller controller.Controller
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
	router.HandleFunc("/create-short-url", utils.Handle(server.Controller.CreateShort)).Methods("POST")
	router.HandleFunc("/{shortUrl}", utils.Handle(server.Controller.ServeShort)).Methods("GET")

	server.Router = router
	return server
}


