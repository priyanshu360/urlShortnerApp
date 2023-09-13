package models

type URLRecord struct {
	Hash    string
	LongURL string
}

type CreateShortURLReq struct {
	LongURL string `json:"long_url" validate:"required"`
}

type APIResult struct {
	Status  int
	Message interface{}
}