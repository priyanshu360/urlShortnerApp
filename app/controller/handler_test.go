package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/priyanshu360/urlShortnerApp.git/models"
	"github.com/priyanshu360/urlShortnerApp.git/storage"
	"github.com/priyanshu360/urlShortnerApp.git/utils"
)

func TestCreateShort(t *testing.T){
	mongo, err := storage.InitMongo()
	if err != nil {
		t.Error(err)
	}
	ctrl := NewController(mongo)
	server := httptest.NewServer(utils.Handle(ctrl.CreateShort))

	payload := strings.NewReader(`{"long_url" : "www.google.com"}`)
	resp, err := http.Post(server.URL, "application/json", payload)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but received %d", resp.StatusCode)
	}
}

func TestGetLongUrl(t *testing.T){
	mongo, err := storage.InitMongo()
	if err != nil {
		t.Error(err)
	}
	ctrl := NewController(mongo)

	router := mux.NewRouter()
    router.HandleFunc("/{shortUrl}", utils.Handle(ctrl.GetLongUrl)).Methods("GET")
    server := httptest.NewServer(router)

	ctrl.DB.CreateURLRecord(&models.URLRecord{
		Hash : "testHash",
		LongURL: "http://www.test.com",
	})

	// url := "localhost:8081/testHash"
	resp, err := http.Get(server.URL + "/testHash")
	fmt.Println(server.URL + "/testHash")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but received %d", resp.StatusCode)
	}
}