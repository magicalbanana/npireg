package npireg

import "net/url"

import "net/http"

// Client ...
type Client struct {
	client  *http.Client // HTTP client used to communicate with the API.
	BaseURL *url.URL
}
