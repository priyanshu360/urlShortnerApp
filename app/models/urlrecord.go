package models

type URLRecord struct {
	Hash    string
	LongURL string
}

type CreateShortURLReq struct {
	LongURL string `json:"long_url"`
}