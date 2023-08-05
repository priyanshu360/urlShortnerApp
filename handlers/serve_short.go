package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *Handler) ServeShort(rw http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	longUrl, err := handler.DB.GetLongUrl(shortUrl)
	if err != nil {
		return err
	}

	http.Redirect(rw, r, longUrl, http.StatusFound)
	return nil
}
