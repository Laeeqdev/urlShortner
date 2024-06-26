package models

type ShortUrlRequest struct {
	LongUrl string `json:"long_url" validate:"required,url"`
}

type LongUrlRequest struct {
	ShortUrl string `json:"short_url" validate:"required,url"`
}

type UrlResponse struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type DomainCount struct {
	Domain string
	Count  int
}
