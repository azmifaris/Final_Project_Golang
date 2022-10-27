package domain

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Title     string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"title" json:"title" example:"Title"`
	Caption   string     `form:"caption" json:"caption"`
	PhotoUrl  string     `gorm:"not null" valid:"required" form:"photoUrl" json:"photoUrl" example:"https://www.example.com/image.jpg"`
	UserID    string     `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User      User       `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt *time.Time `gorm:"not null" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"not null" json:"updatedAt,omitempty"`
	Comments  []Comment  `json:"comments"`
}

func (p *Photo) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

func (p *Photo) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

type PhotoUseCase interface {
	Fetch(context.Context, *[]Photo) error
	Store(context.Context, *Photo) error
	GetByID(context.Context, *Photo, string) error
	Update(context.Context, Photo, string) (Photo, error)
	Delete(context.Context, string) error
}

type PhotoRepository interface {
	Fetch(context.Context, *[]Photo) error
	Store(context.Context, *Photo) error
	GetByID(context.Context, *Photo, string) error
	Update(context.Context, Photo, string) (Photo, error)
	Delete(context.Context, string) error
}
