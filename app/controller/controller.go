package controller

import (
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

func response(status int, message interface{}) models.APIResult {
	return models.APIResult{
		Status: status,
		Message: message,
	}
}