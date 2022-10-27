package utility

import "time"

type NewSocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"socialMediaUrl"`
	UserID         string     `json:"userId"`
	CreatedAt      *time.Time `json:"createdAt"`
}

type UpdatedSocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"socialMediaUrl"`
	UserID         string     `json:"userId"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}
