package repository

import "errors"

const (
	EmptyUrlError            = "url must not be empty"
	EmptyLongOrShortUrlError = "longUrl or shortUrl must not be empty"
)

type UrlMemoryRepository struct {
	database map[string]string
	err      error
}

func NewUrlMemoryRepository() *UrlMemoryRepository {
	return &UrlMemoryRepository{
		database: make(map[string]string),
	}
}

func (m *UrlMemoryRepository) Add(longUrl, shortUrl string) bool {
	if m.err != nil {
		return false
	}

	if longUrl == "" || shortUrl == "" {
		m.err = errors.New(EmptyLongOrShortUrlError)
		return false
	}

	m.database[longUrl] = shortUrl

	return true
}

func (m *UrlMemoryRepository) FindByShortUrl(url string) string {
	if m.err != nil {
		return ""
	}

	if url == "" {
		m.err = errors.New(EmptyUrlError)
		return ""
	}

	for longUrl, shortUrl := range m.database {
		if shortUrl == url {
			return longUrl
		}
	}

	return ""
}

func (m *UrlMemoryRepository) FindByLongUrl(url string) string {
	if m.err != nil {
		return ""
	}

	if url == "" {
		m.err = errors.New(EmptyUrlError)
		return ""
	}

	shortUrl, has := m.database[url]
	if !has {
		return ""
	}

	return shortUrl
}

func (m *UrlMemoryRepository) Remove(longUrl string) bool {
	if m.err != nil {
		return false
	}

	if longUrl == "" {
		m.err = errors.New(EmptyUrlError)
		return false
	}

	if url := m.FindByLongUrl(longUrl); url == "" {
		return false
	}

	delete(m.database, longUrl)

	return true
}

func (m *UrlMemoryRepository) GetError() error {
	return m.err
}
