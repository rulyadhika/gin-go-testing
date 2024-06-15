package dto

type NewBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookResponse struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
