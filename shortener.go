package shortener

import (
	"errors"
	"github.com/koind/shortener/hasher"
	"github.com/koind/shortener/repository"
	"github.com/koind/shortener/stringer"
	"net/url"
)

type Shortener interface {
	Shorten(url string) (string, error)
	Resolve(url string) (string, error)
}

type LinkShortener struct {
	repo   repository.UrlRepository
	hasher hasher.HashGenerator
}

func NewShortener(repository repository.UrlRepository, hasher hasher.HashGenerator) *LinkShortener {
	return &LinkShortener{
		repo:   repository,
		hasher: hasher,
	}
}

func (l *LinkShortener) Shorten(fullUrl string) (string, error) {
	if fullUrl == "" {
		return "", errors.New(repository.EmptyUrlError)
	}

	u, err := url.Parse(fullUrl)
	if err != nil {
		return "", err
	}

	urlPath := u.RequestURI()
	urlHash, err := l.hasher.Generate(urlPath)
	if err != nil {
		return "", err
	}

	urlHash = stringer.Substr(urlHash, 0, 6)

	shortUrl, err := u.Parse(urlHash)
	if err != nil {
		return "", err
	}

	_, err = l.repo.Add(fullUrl, shortUrl.String())
	if err != nil {
		return "", err
	}

	return shortUrl.String(), nil
}

func (l *LinkShortener) Resolve(shortUrl string) (string, error) {
	if shortUrl == "" {
		return "", errors.New(repository.EmptyUrlError)
	}

	fullUrl, err := l.repo.FindByShortUrl(shortUrl)
	if err != nil {
		return "", err
	}

	return fullUrl, nil
}
