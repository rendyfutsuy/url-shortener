package repository

import (
	"context"
	"time"

	"github.com/rendyfutsuy/base-go/models"
	"github.com/rendyfutsuy/base-go/modules/post/dto"
	"github.com/rendyfutsuy/base-go/utils"
)

func (r *postRepository) CreateShortener(ctx context.Context, data dto.ToDBPostShortener) (*models.PostShortUrl, error) {

	code := utils.GenerateRandomString(6)
	// // hashed short url code
	// hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), 1)
	// if err != nil {
	// 	utils.Logger.Error(err.Error())
	// 	return nil, err
	// }

	now := time.Now().UTC()
	c := &models.PostShortUrl{
		FullUrl: data.FullUrl,
		// Code:      string(hashedCode),
		Code:      code,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := r.DB.WithContext(ctx).Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *postRepository) GetShortedURLByCode(ctx context.Context, code string) (*models.PostShortUrl, error) {
	c := &models.PostShortUrl{}
	if err := r.DB.WithContext(ctx).
		Table("post_short_urls psu").
		Select("*").
		Where("psu.code = ?", code).
		First(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
