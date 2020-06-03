package models

// Blog models
type Blog struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"authorId"`
}
