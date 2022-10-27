package usecase

import (
	"context"
	"mygram/domain"
)

type socialMediaUseCase struct {
	socialMediaRepository domain.SocialMediaRepository
}

func NewSocialMediaUseCase(socialMediaRepository domain.SocialMediaRepository) *socialMediaUseCase {
	return &socialMediaUseCase{socialMediaRepository}
}

func (socialMediaUseCase *socialMediaUseCase) Fetch(ctx context.Context, socialMedias *[]domain.SocialMedia, userID string) (err error) {
	err = socialMediaUseCase.socialMediaRepository.Fetch(ctx, socialMedias, userID)
	if err != nil {
		return err
	}
	return
}

func (socialMediaUseCase *socialMediaUseCase) Store(ctx context.Context, socialMedia *domain.SocialMedia) (err error) {
	err = socialMediaUseCase.socialMediaRepository.Store(ctx, socialMedia)
	if err != nil {
		return err
	}
	return
}

func (socialMediaUseCase *socialMediaUseCase) GetByID(ctx context.Context, socialMedia *domain.SocialMedia, id string) (err error) {
	err = socialMediaUseCase.socialMediaRepository.GetByID(ctx, socialMedia, id)
	if err != nil {
		return err
	}
	return
}

func (socialMediaUseCase *socialMediaUseCase) Update(ctx context.Context, socialMedia domain.SocialMedia, id string) (socmed domain.SocialMedia, err error) {
	sosmed, err := socialMediaUseCase.socialMediaRepository.Update(ctx, socialMedia, id)
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (socialMediaUseCase *socialMediaUseCase) Delete(ctx context.Context, id string) (err error) {
	err = socialMediaUseCase.socialMediaRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return
}
