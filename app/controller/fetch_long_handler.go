package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priyanshu360/urlShortnerApp.git/models"
)


func (s *Controller)GetLongUrl(rw http.ResponseWriter, r *http.Request) models.APIResult {
    shortUrl := mux.Vars(r)["shortUrl"]
	fmt.Println(shortUrl, "hellllllllllllllll")
	longUrl, err := s.DB.GetLongUrl(shortUrl)
	if err != nil {
		return response(http.StatusNotFound, err.Error())
	}

	return response(http.StatusOK, longUrl)
}