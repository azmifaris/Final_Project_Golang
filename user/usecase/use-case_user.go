package usecase

import (
	"context"
	"mygram/domain"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (UserUseCase *UserUseCase) Register(ctx context.Context, user *domain.User) (err error) {
	err = UserUseCase.userRepository.Register(ctx, user)
	if err != nil {
		return err
	}
	return
}

func (UserUseCase *UserUseCase) Login(ctx context.Context, user *domain.User) (err error) {
	err = UserUseCase.userRepository.Login(ctx, user)
	if err != nil {
		return nil
	}
	return
}
func (UserUseCase *UserUseCase) Update(ctx context.Context, user domain.User) (u domain.User, err error) {
	u, err = UserUseCase.userRepository.Update(ctx, user)
	if err != nil {
		return u, err
	}
	return u, nil
}
func (UserUseCase *UserUseCase) Delete(ctx context.Context, id string) (err error) {
	err = UserUseCase.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return
}
