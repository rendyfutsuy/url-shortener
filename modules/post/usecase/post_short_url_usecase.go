package usecase

import (
	"context"

	"github.com/rendyfutsuy/base-go/models"
	"github.com/rendyfutsuy/base-go/modules/post/dto"
)

func (u *postUsecase) CreatePostUrlShortener(ctx context.Context, req *dto.ReqUrlShortener) (*models.PostShortUrl, error) {
	c, err := u.repo.CreateShortener(ctx, dto.ToDBPostShortener{
		FullUrl: req.Url,
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (u *postUsecase) GetShortedURLByCode(ctx context.Context, code string) (*models.PostShortUrl, error) {
	return u.repo.GetShortedURLByCode(ctx, code)
}
