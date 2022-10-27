package repository

import (
	"context"
	"fmt"
	"mygram/domain"
	"time"

	nano "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (socialMediaRepository *socialMediaRepository) Fetch(ctx context.Context, socialMedias *[]domain.SocialMedia, userID string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = socialMediaRepository.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username", "ProfileImageUrl")
	}).Find(&socialMedias).Error
	if err != nil {
		return err
	}
	return
}

func (socialMediaRepository *socialMediaRepository) Store(ctx context.Context, socialMedia *domain.SocialMedia) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ID, _ := nano.New(16)
	socialMedia.ID = fmt.Sprintf("socialmedia-%s", ID)
	err = socialMediaRepository.db.WithContext(ctx).Create(&socialMedia).Error
	if err != nil {
		return err
	}
	return
}

func (socialMediaRepository *socialMediaRepository) GetByID(ctx context.Context, socialMedia *domain.SocialMedia, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = socialMediaRepository.db.WithContext(ctx).First(&socialMedia, &id).Error
	if err != nil {
		return err
	}
	return
}

func (socialMediaRepository *socialMediaRepository) Update(ctx context.Context, socialMedia domain.SocialMedia, id string) (socmed domain.SocialMedia, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	sosmed := domain.SocialMedia{}
	err = socialMediaRepository.db.WithContext(ctx).First(&sosmed, &id).Error
	if err != nil {
		return sosmed, err
	}
	err = socialMediaRepository.db.WithContext(ctx).Model(&sosmed).Updates(socialMedia).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (socialMediaRepository *socialMediaRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = socialMediaRepository.db.WithContext(ctx).First(&domain.SocialMedia{}).Error
	if err != nil {
		return err
	}
	err = socialMediaRepository.db.WithContext(ctx).Delete(&domain.SocialMedia{}, &id).Error
	if err != nil {
		return err
	}
	return
}
