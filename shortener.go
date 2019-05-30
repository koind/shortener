package shortener

import (
	"errors"
	"github.com/koind/shortener/repository"
)

const EmptyUrlError = "url must not be empty"

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinkShortener struct {
	repo repository.UrlRepository
	err  error
}

func NewShortener(repository repository.UrlRepository) *LinkShortener {
	return &LinkShortener{
		repo: repository,
	}
}

func (l *LinkShortener) Shorten(url string) string {
	if l.err != nil {
		return ""
	}

	if url == "" {
		l.err = errors.New(EmptyUrlError)
		return ""
	}

	return ""
}

func (l *LinkShortener) Resolve(url string) string {
	if l.err != nil {
		return ""
	}

	if url == "" {
		l.err = errors.New(EmptyUrlError)
		return ""
	}

	return l.repo.FindByShortUrl(url)
}

func (l *LinkShortener) GetError() error {
	return l.err
}
