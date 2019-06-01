package repository

import "errors"

const (
	EmptyUrlError            = "url must not be empty"
	EmptyFullOrShortUrlError = "fullUrl or shortUrl must not be empty"
	NoSuchUrl                = "no such url"
)

type UrlMemoryRepository struct {
	database map[string]string
}

func NewUrlMemoryRepository() *UrlMemoryRepository {
	return &UrlMemoryRepository{
		database: make(map[string]string),
	}
}

func (m *UrlMemoryRepository) Add(fullUrl, shortUrl string) (bool, error) {
	if fullUrl == "" || shortUrl == "" {
		return false, errors.New(EmptyFullOrShortUrlError)
	}

	m.database[shortUrl] = fullUrl

	return true, nil
}

func (m *UrlMemoryRepository) FindByShortUrl(url string) (string, error) {
	if url == "" {
		return "", errors.New(EmptyUrlError)
	}

	fullUrl, has := m.database[url]
	if !has {
		return "", errors.New(NoSuchUrl)
	}

	return fullUrl, nil
}

func (m *UrlMemoryRepository) Remove(shortUrl string) (bool, error) {
	if shortUrl == "" {
		return false, errors.New(EmptyUrlError)
	}

	_, err := m.FindByShortUrl(shortUrl)
	if err != nil {
		return false, err
	}

	delete(m.database, shortUrl)

	return true, nil
}
