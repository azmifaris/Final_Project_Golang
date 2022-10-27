package repository

import (
	"context"
	"errors"
	"fmt"
	"mygram/domain"
	"mygram/helpers"
	"time"

	nano "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (UserRepository *UserRepository) Register(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ID, _ := nano.New(16)
	user.ID = fmt.Sprintf("user-%s", ID)
	err = UserRepository.db.Debug().WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return
}

func (UserRepository *UserRepository) Login(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	password := user.Password
	err = UserRepository.db.Debug().WithContext(ctx).Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		return errors.New("email not found")
	}
	isValid := helpers.Compare([]byte(user.Password), []byte(password))
	if !isValid {
		return errors.New("invalid credential")
	}
	return
}

func (UserRepository *UserRepository) Update(ctx context.Context, user domain.User) (u domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	u = domain.User{}
	err = UserRepository.db.Debug().WithContext(ctx).First(&u).Error
	if err != nil {
		return u, err
	}
	err = UserRepository.db.Debug().WithContext(ctx).Model(&u).Updates(user).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func (UserRepository *UserRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = UserRepository.db.Debug().WithContext(ctx).Where("id=?").Delete(&domain.User{}, id).Error
	if err != nil {
		return err
	}
	return
}
