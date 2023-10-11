package controller

import (
	"encoding/json"

	"github.com/priyanshu360/urlShortnerApp.git/models"
	"github.com/priyanshu360/urlShortnerApp.git/storage"
)

type Controller struct {
	DB storage.Storage
}

func NewController(db storage.Storage) *Controller {
	return &Controller{
		DB : db,
	}
}

func response(status int, val interface{}) models.APIResult {
	data, _ := json.Marshal(val)
	return models.APIResult{
		Status: status,
		Body : struct{Data string "json:\"data\""}{
			Data: string(data),
		},
	}
}