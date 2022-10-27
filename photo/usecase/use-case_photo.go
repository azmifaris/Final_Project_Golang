package usecase

import (
	"context"
	"mygram/domain"
)

type photoUseCase struct {
	photoRepository domain.PhotoRepository
}

func NewPhotoUsecase(photoRepository domain.PhotoRepository) *photoUseCase {
	return &photoUseCase{photoRepository}
}
func (photoUseCase *photoUseCase) Fetch(ctx context.Context, photos *[]domain.Photo) (err error) {
	err = photoUseCase.photoRepository.Fetch(ctx, photos)
	if err != nil {
		return err
	}
	return
}
func (photoUseCase *photoUseCase) Store(ctx context.Context, photo *domain.Photo) (err error) {
	err = photoUseCase.photoRepository.Store(ctx, photo)
	if err != nil {
		return err
	}
	return
}
func (photoUseCase *photoUseCase) GetByID(ctx context.Context, photo *domain.Photo, id string) (err error) {
	err = photoUseCase.photoRepository.GetByID(ctx, photo, id)
	if err != nil {
		return err
	}
	return
}
func (photoUseCase *photoUseCase) Update(ctx context.Context, photo domain.Photo, id string) (p domain.Photo, err error) {
	p, err = photoUseCase.photoRepository.Update(ctx, photo, id)
	if err != nil {
		return p, err
	}
	return p, nil
}
func (photoUseCase *photoUseCase) Delete(ctx context.Context, id string) (err error) {
	err = photoUseCase.photoRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return
}
