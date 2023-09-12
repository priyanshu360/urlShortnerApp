package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/priyanshu360/urlShortnerApp.git/models"
	"github.com/priyanshu360/urlShortnerApp.git/utils"
)

func (handler *Handler) CreateShort(rw http.ResponseWriter, r *http.Request) error {
	var reqBody models.CreateShortURLReq
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return err
	}

	urlRecord, err := handler.createNewUrlRecord(&reqBody)
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

func (handler Handler) createNewUrlRecord(request *models.CreateShortURLReq) (*models.URLRecord, error) {
	for retries := 0; retries < 5; retries += 1 {
		hash := utils.GenerateRandomHash()

		if longUrl, err := handler.DB.GetLongUrl(hash); err != nil {
			return nil, err
		} else if longUrl == "" {
			return &models.URLRecord{
				Hash:    utils.GenerateRandomHash(),
				LongURL: request.LongURL,
			}, nil
		}
	}
	return nil, fmt.Errorf("Fail to create unique hash")
}
