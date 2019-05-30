package repository

type UrlRepository interface {
	FindByShortUrl(url string) string
	FindByLongUrl(url string) string
	GetError() error
}
