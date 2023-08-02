package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/priyanshu360/urlShortnerApp.git/models"
	"github.com/priyanshu360/urlShortnerApp.git/utils"
)


func (handler *Handler) CreateShort(rw http.ResponseWriter, r *http.Request) error {
	var reqBody models.CreateShortURLReq
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return err
	}


	urlRecord := models.URLRecord{
		Hash:    utils.GenerateRandomHash(),
		LongURL: reqBody.LongURL,
	}

	if err := handler.DB.CreateURLRecord(urlRecord); err != nil {
		return err
	}

	responseJSON, err := json.Marshal(urlRecord)
	if err != nil {
		return err
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(responseJSON)
	return nil
}
