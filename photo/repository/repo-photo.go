package repository

import (
	"context"
	"fmt"
	"mygram/domain"
	"time"

	nano "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (photoRepository *photoRepository) Fetch(ctx context.Context, photos *[]domain.Photo) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = photoRepository.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Find(&photos).Error
	if err != nil {
		return err
	}
	return
}

func (photoRepository *photoRepository) Store(ctx context.Context, photo *domain.Photo) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ID, _ := nano.New(16)
	photo.ID = fmt.Sprintf("photo-%s", ID)
	err = photoRepository.db.WithContext(ctx).Create(&photo).Error
	if err != nil {
		return err
	}
	return
}

func (photoRepository *photoRepository) GetByID(ctx context.Context, photo *domain.Photo, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = photoRepository.db.WithContext(ctx).First(&photo, &id).Error
	if err != nil {
		return err
	}
	return
}

func (photoRepository *photoRepository) Update(ctx context.Context, photo domain.Photo, id string) (p domain.Photo, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	p = domain.Photo{}
	err = photoRepository.db.WithContext(ctx).First(&p, &id).Error
	if err != nil {
		return p, err
	}
	err = photoRepository.db.WithContext(ctx).Model(&p).Updates(photo).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (photoRepository *photoRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = photoRepository.db.WithContext(ctx).First(&domain.Photo{}).Error
	if err != nil {
		return err
	}
	err = photoRepository.db.WithContext(ctx).Delete(&domain.Photo{}, &id).Error
	if err != nil {
		return err
	}
	return
}
