package usecase

import (
	"context"
	"mygram/domain"
)

type commentUseCase struct {
	commentRepository domain.CommentRepository
}

func NewCommentUseCase(commentRepository domain.CommentRepository) *commentUseCase {
	return &commentUseCase{commentRepository}
}

func (commentUseCase *commentUseCase) Fetch(ctx context.Context, comments *[]domain.Comment, userID string) (err error) {
	err = commentUseCase.commentRepository.Fetch(ctx, comments, userID)
	if err != nil {
		return err
	}
	return
}

func (commentUseCase *commentUseCase) Store(ctx context.Context, comment *domain.Comment) (err error) {
	err = commentUseCase.commentRepository.Store(ctx, comment)
	if err != nil {
		return err
	}
	return
}

func (commentUseCase *commentUseCase) GetByID(ctx context.Context, comment *domain.Comment, id string) (err error) {
	err = commentUseCase.commentRepository.GetByID(ctx, comment, id)
	if err != nil {
		return err
	}
	return
}

func (commentUseCase *commentUseCase) Update(ctx context.Context, comment domain.Comment, id string) (photo domain.Photo, err error) {
	photo, err = commentUseCase.commentRepository.Update(ctx, comment, id)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (commentUseCase *commentUseCase) Delete(ctx context.Context, id string) (err error) {
	err = commentUseCase.commentRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return
}
