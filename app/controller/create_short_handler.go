package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/priyanshu360/urlShortnerApp.git/models"
	"github.com/priyanshu360/urlShortnerApp.git/utils"
)

func (ctrl *Controller) CreateShort(rw http.ResponseWriter, r *http.Request) models.APIResult {
	var reqBody models.CreateShortURLReq
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return response(http.StatusBadRequest, err.Error())
	}
	
	err := validator.New().Struct(reqBody)
	if err != nil {
		return response(http.StatusBadRequest, err.Error())
	}

	urlRecord, err := ctrl.createNewUrlRecord(&reqBody)
	if err != nil {
		return response(http.StatusInternalServerError, err.Error())
	}

	if err := ctrl.DB.CreateURLRecord(urlRecord); err != nil {
		return response(http.StatusInternalServerError, err.Error())
	}

	return response(http.StatusOK, urlRecord)
}

func (ctrl *Controller) createNewUrlRecord(request *models.CreateShortURLReq) (*models.URLRecord, error) {
	for retries := 0; retries < 5; retries += 1 {
		hash := utils.GenerateRandomHash()

		longUrl, _ := ctrl.DB.GetLongUrl(hash)
		if longUrl == "" {
			return &models.URLRecord{
				Hash:    hash,
				LongURL: request.LongURL,
			}, nil
		}

	}
	return nil, fmt.Errorf("Fail to create unique hash")
}
