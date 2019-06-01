package repository

type UrlRepository interface {
	Add(longUrl, shortUrl string) (bool, error)
	FindByShortUrl(url string) (string, error)
	Remove(url string) (bool, error)
}
