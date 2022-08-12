package models

type User struct {
	Firstname     string `json:"first_name"`
	Lastname      string `json:"last_name"`
	ID            string `json:"id"`
	OwnsPlaylists bool   `json:"ownsplaylists"`
}
