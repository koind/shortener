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

func NewShortener() *LinkShortener {
	return &LinkShortener{
		storage: make(map[string]string),
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

	return l.findByShortUrl(url)
}

func (l *LinkShortener) GetError() error {
	return l.err
}

func (l *LinkShortener) findByShortUrl(url string) string {
	if l.err != nil {
		return ""
	}

	for longUrl, shortUrl := range l.storage {
		if shortUrl == url {
			return longUrl
		}
	}

	return ""
}
