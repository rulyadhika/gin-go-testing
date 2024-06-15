package dto

type APIResponse struct {
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
