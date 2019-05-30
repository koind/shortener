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
	storage map[string]string
	err     error
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

	return ""
}

func (l *LinkShortener) GetError() error {
	return l.err
}

func (l *LinkShortener) hasLink(url string) bool {
	if l.err != nil {
		return false
	}

	_, has := l.storage[url]

	return has
}

func NewShortener() *LinkShortener {
	return new(LinkShortener)
}
