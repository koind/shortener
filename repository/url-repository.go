package repository

type UrlRepository interface {
	Add(longUrl, shortUrl string) bool
	FindByShortUrl(url string) string
	FindByLongUrl(url string) string
	Remove(url string) bool
	GetError() error
}
