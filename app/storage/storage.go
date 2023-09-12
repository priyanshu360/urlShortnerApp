package storage

import "github.com/priyanshu360/urlShortnerApp.git/models"

type Storage interface {
	GetLongUrl(string) (string, error)
	GetHashValue(string) (string, error)
	DeleteURLRecord(string) error
	CreateURLRecord(*models.URLRecord) error
}
