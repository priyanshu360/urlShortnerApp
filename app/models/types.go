package models

// URLRecord represents a URL record in the system.
//
// swagger:model
type URLRecord struct {
	// Hash for the short URL.
	//
	// required: true
	// example: abc123
	Hash string `json:"hash"`
	
	// LongURL is the original URL.
	//
	// required: true
	// example: http://www.youtube.com
	LongURL string `json:"long_url"`
}

// CreateShortURLReq represents the request body parameter for creating a short URL.
//
// swagger:model 
type CreateShortURLReq struct {
	// LongURL is the original URL.
	// required: true
	// example: http://www.youtube.com
	LongURL string `json:"long_url" validate:"required"`
}


// APIResult represents the API response.
//
// swagger:response APIResult
// APIResult represents a standard API response.
type APIResult struct {
	// in: body
    Body struct {
		// Data is the response data, which can be either a string or an URLRecord.
		// required: true
		// example: "Not Found"
		Data string `json:"data"`
	}

	// Status is the HTTP status code of the response.
	// required: true
	// example: 404
	Status int `json:"status"`
}



