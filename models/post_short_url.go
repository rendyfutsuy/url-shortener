package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rendyfutsuy/base-go/utils"
)

type PostShortUrl struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v7()" json:"id"`
	FullUrl   string    `gorm:"column:full_url;type:text;not null" json:"full_url"`
	Code      string    `gorm:"column:code;type:varchar(255);not null" json:"code"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (PostShortUrl) TableName() string {
	return "post_short_urls"
}

func (PostShortUrl PostShortUrl) GetShortedURL() string {
	baseURL := utils.ConfigVars.String("app_url") + ":" + utils.ConfigVars.String("app_port") + "/v1/post/url-shortener/"
	return baseURL + PostShortUrl.Code
}
