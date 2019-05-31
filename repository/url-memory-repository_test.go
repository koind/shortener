package repository

import (
	"testing"
)

var urlMemoryRepository *UrlMemoryRepository
var longUrl, shortUrl = "http://site.ru/test-simple-page", "http://site.ru/ldj32"

func init() {
	urlMemoryRepository = NewUrlMemoryRepository()
}

func TestUrlMemoryRepository_Add(t *testing.T) {
	urlMemoryRepository.Add(longUrl, shortUrl)

	if url := urlMemoryRepository.FindByLongUrl(longUrl); url != shortUrl {
		t.Errorf("shortUrl does not match %s - %s", shortUrl, url)
	}
}

func TestUrlMemoryRepository_FindByLongUrl(t *testing.T) {
	if url := urlMemoryRepository.FindByLongUrl(longUrl); url != shortUrl {
		t.Errorf("shortUrl does not match %s - %s", shortUrl, url)
	}
}

func TestUrlMemoryRepository_FindByShortUrl(t *testing.T) {
	if url := urlMemoryRepository.FindByShortUrl(shortUrl); url != longUrl {
		t.Errorf("longUrl does not match %s - %s", longUrl, url)
	}
}

func TestUrlMemoryRepository_Remove(t *testing.T) {
	urlMemoryRepository.Remove(longUrl)

	if url := urlMemoryRepository.FindByLongUrl(longUrl); url != "" {
		t.Error("Urls must be removed")
	}
}

func TestUrlMemoryRepository_GetError(t *testing.T) {
	urlMemoryRepository.Add("", "")

	if err := urlMemoryRepository.GetError(); err.Error() != EmptyLongOrShortUrlError {
		t.Errorf("Errors does not match %s - %s", EmptyLongOrShortUrlError, err)
	}
}
