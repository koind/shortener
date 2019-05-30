package shortener

type UrlRepository interface {
	FindByShortUrl(url string) string
	FindByLongUrl(url string) string
	GetError() error
}

type UrlMemoryRepository struct {
	database map[string]string
	err      error
}

func NewUrlMemoryRepository() *UrlMemoryRepository {
	return &UrlMemoryRepository{
		database: make(map[string]string),
	}
}

func (m *UrlMemoryRepository) FindByShortUrl(url string) string {
	if m.err != nil {
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

	for longUrl, shortUrl := range m.database {
		if longUrl == url {
			return shortUrl
		}
	}

	return ""
}

func (m *UrlMemoryRepository) GetError() error {
	return m.err
}
