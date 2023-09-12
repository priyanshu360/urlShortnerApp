package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Handler)ServeShort(rw http.ResponseWriter, r *http.Request) error {
    shortUrl := mux.Vars(r)["shortUrl"]
	longUrl, err := s.DB.GetLongUrl(shortUrl)
	if err != nil {
		return err
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(longUrl))

	return nil
}