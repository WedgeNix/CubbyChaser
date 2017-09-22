package ship

import "net/http"

// Control controller data for shipstaion calls.
type Control struct {
	shipURL  string
	username string
	password string
	client   http.Client
}
