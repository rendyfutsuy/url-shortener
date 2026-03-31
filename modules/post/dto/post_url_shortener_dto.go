package dto

import (
	"github.com/rendyfutsuy/base-go/models"
)

type ReqUrlShortener struct {
	Url string `json:"url" validate:"required" form:"url"`
}

type GetUrlShortener struct {
	FullUrl      string `json:"full_url"`
	ShortenerUrl string `json:"shortener_url"`
}

func ToUrlShortener(m *models.PostShortUrl) GetUrlShortener {
	return GetUrlShortener{
		// FullUrl: m.FullUrl,
		// ShortenerUrl: m.GetShortedURL(),
	}
}
