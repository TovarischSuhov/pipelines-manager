package models

import "time"

type Commit struct {
	ID          string    `json:"id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	Title       string    `json:"title"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	ShortID     string    `json:"short_id"`
}
