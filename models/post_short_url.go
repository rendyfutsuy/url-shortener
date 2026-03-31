package models

import (
	"time"

	"github.com/google/uuid"
)

type PostShortUrl struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v7()" json:"id"`
	CreatedBy uuid.UUID `gorm:"column:created_by;type:uuid;not null" json:"created_by"`
	FullUrl   string    `gorm:"column:full_url;type:text;not null" json:"full_url"`
	Code      string    `gorm:"column:code;type:varchar(255);not null" json:"code"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (PostShortUrl) TableName() string {
	return "post_short_urls"
}

func (PostShortUrl PostShortUrl) GetShortedURL() string {
	baseURL := "https://shorturl.com/"
	return baseURL + PostShortUrl.Code
}
