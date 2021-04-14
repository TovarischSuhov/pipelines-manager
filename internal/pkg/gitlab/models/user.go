package models

import "time"

type User struct {
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	ID           int64     `json:"id"`
	State        string    `json:"state"`
	AvatarURL    string    `json:"avatar_url"`
	WebURL       string    `json:"web_url"`
	CreatedAt    time.Time `json:"created_at"`
	Bio          *string   `json:"bio"`
	Location     *string   `json:"location"`
	PublicEmail  string    `json:"public_email"`
	Skype        string    `json:"skype"`
	Linkedin     string    `json:"linkedin"`
	Twitter      string    `json:"twitter"`
	WebsiteURL   string    `json:"website_url"`
	Organization string    `json:"organization"`
}
