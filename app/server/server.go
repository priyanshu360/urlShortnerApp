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

	// swagger:route POST /create-short-url createShortURL
	// create short url
	//
	// responses:
	//  503: APIResult
	//  400: APIResult
	//  200: APIResult
	router.HandleFunc("/create-short-url", utils.Handle(server.Controller.CreateShort)).Methods("POST")
	
	
	// swagger:route GET /{shortUrl} getLongUrl
	// get long URL
	//
	// retrieves the long URL associated with a short URL.
	//
	// responses:
	//   503: APIResult
	//   404: APIResult
	//   400: APIResult
	//   200: APIResult
	router.HandleFunc("/{shortUrl}", utils.Handle(server.Controller.GetLongUrl)).Methods("GET")

	server.Router = router
	return server
}


