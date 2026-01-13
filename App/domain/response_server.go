package domain

type Response struct {
	Status     bool    `json:"status"`
	StatusCode int     `json:"statusCode"`
	Message    string  `json:"message"`
	Error      *string `json:"error,omitempty"`
	Data       any     `json:"data,omitempty"`
}
