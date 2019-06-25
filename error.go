package npireg

// Error ...
type Error struct {
	Message  string `json:"message,omitempty"`
	Response struct {
		StatusCode int    `json:"status_code,omitempty"`
		Body       string `json:"body,omitempty"`
	} `json:"response,omitempty"`
}

// Error ...
func (e Error) Error() string {
	return e.Message
}
