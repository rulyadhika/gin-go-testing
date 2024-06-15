package dto

type APIResponse struct {
	Status     string `json:"status"`
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
