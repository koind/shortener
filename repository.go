package shortener

type Repository interface {
	FindByShortUrl(url string) string
	FindByLongUrl(url string) string
	GetError() error
}

type MemoryRepository struct {
	database map[string]string
	err      error
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		database: make(map[string]string),
	}
}

func (m *MemoryRepository) FindByShortUrl(url string) string {
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

func (m *MemoryRepository) FindByLongUrl(url string) string {
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

func (m *MemoryRepository) GetError() error {
	return m.err
}
