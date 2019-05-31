package shortener

import (
	"errors"
	"github.com/koind/shortener/hasher"
	"github.com/koind/shortener/repository"
	"github.com/koind/shortener/stringer"
	"net/url"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinkShortener struct {
	repo   repository.UrlRepository
	hasher hasher.HashGenerator
	err    error
}

func NewShortener(repository repository.UrlRepository, hasher hasher.HashGenerator) *LinkShortener {
	return &LinkShortener{
		repo:   repository,
		hasher: hasher,
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
	urlHash := l.hasher.Generate(urlPath)
	if err := l.hasher.GetError(); err != nil {
		l.err = err
		return ""
	}

	urlHash = stringer.Substr(urlHash, 0, 6)

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
