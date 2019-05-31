package shortener

import (
	"errors"
	"github.com/koind/shortener/hash"
	"github.com/koind/shortener/repository"
	"net/url"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinkShortener struct {
	repo repository.UrlRepository
	hash hash.HashGenerator
	err  error
}

func NewShortener(repository repository.UrlRepository, hash hash.HashGenerator) *LinkShortener {
	return &LinkShortener{
		repo: repository,
		hash: hash,
	}
}

func (l *LinkShortener) Shorten(longUrl string) string {
	if l.err != nil {
		return ""
	}

	if longUrl == "" {
		l.err = errors.New(repository.EmptyUrlError)
		return ""
	}

	u, err := url.Parse(longUrl)
	if err != nil {
		l.err = err
		return ""
	}

	urlPath := u.RequestURI()
	urlHash := l.hash.Generate(urlPath)
	if err := l.hash.GetError(); err != nil {
		l.err = err
		return ""
	}

	shortUrl, err := u.Parse(urlHash)
	if err != nil {
		l.err = err
		return ""
	}

	l.repo.Add(longUrl, shortUrl.String())

	return shortUrl.String()
}

func (l *LinkShortener) Resolve(shortUrl string) string {
	if l.err != nil {
		return ""
	}

	if shortUrl == "" {
		l.err = errors.New(repository.EmptyUrlError)
		return ""
	}

	return l.repo.FindByShortUrl(shortUrl)
}

func (l *LinkShortener) GetError() error {
	return l.err
}
