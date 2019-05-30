package shortener

import (
	"errors"
)

const EmptyUrlError = "url must not be empty"

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinkShortener struct {
	repo Repository
	err  error
}

func NewShortener(repository Repository) *LinkShortener {
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
