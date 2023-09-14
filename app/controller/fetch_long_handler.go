package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priyanshu360/urlShortnerApp.git/models"
)


func (s *Controller)GetLongUrl(rw http.ResponseWriter, r *http.Request) models.APIResult {
	// swagger:operation GET /{shortUrl} getLongUrl
	//
	// Get Long URL
	//
	// Retrieves the long URL associated with a short URL.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: shortUrl
	//   in: path
	//   description: Short URL to retrieve the long URL for.
	//   required: true
	//   type: string
	// responses:
	//   404:
	//     description: Not Found
	//     schema:
	//       "$ref": "#/responses/APIResult"
	//   200:
	//     description: Success response
	//     schema:
	//       "$ref": "#/responses/APIResult"
	//   400:
	//     description: Client Side Error
	//     schema:
	//       "$ref": "#/responses/APIResult"
	//   503:
	//     description: Server Side Error
	//     schema:
	//       "$ref": "#/responses/APIResult"

    shortUrl := mux.Vars(r)["shortUrl"]
	longUrl, err := s.DB.GetLongUrl(shortUrl)
	if err != nil {
		return response(http.StatusNotFound, err.Error())
	}

	return response(http.StatusOK, longUrl)
}