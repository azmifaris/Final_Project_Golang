package utility

import "time"

type UserUpdated struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       uint       `json:"age"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
