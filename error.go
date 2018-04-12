package npireg

// Error ...
type Error struct {
	Message  string `json:"message"`
	Response struct {
		StatusCode int    `json:"status_code"`
		Body       string `json:"body"`
	} `json:"response"`
}

// Error ...
func (e Error) Error() string {
	return e.Message
}
