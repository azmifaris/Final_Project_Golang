package utility

import "time"

type NewComment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	PhotoID   string     `json:"photoId"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"createdAt"`
}

type UpdatedComment struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photoUrl"`
	UserID    string     `json:"userId"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
