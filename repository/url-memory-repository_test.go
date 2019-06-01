package repository

import (
	"testing"
)

var urlMemoryRepository *UrlMemoryRepository
var fullUrl, shortUrl = "http://site.ru/test-simple-page", "http://site.ru/ldj32"

func init() {
	urlMemoryRepository = NewUrlMemoryRepository()
}

func TestUrlMemoryRepository_Add(t *testing.T) {
	urlMemoryRepository.Add(fullUrl, shortUrl)

	if url, _ := urlMemoryRepository.FindByShortUrl(shortUrl); url != fullUrl {
		t.Errorf("fullUrl does not match %s - %s", fullUrl, url)
	}
}

func TestUrlMemoryRepository_FindByShortUrl(t *testing.T) {
	if url, _ := urlMemoryRepository.FindByShortUrl(shortUrl); url != fullUrl {
		t.Errorf("fullUrl does not match %s - %s", fullUrl, url)
	}

	if _, err := urlMemoryRepository.FindByShortUrl(shortUrl + "fails"); err.Error() != NoSuchUrl {
		t.Errorf("Errors does not match %s - %s", NoSuchUrl, err)
	}
}

func TestUrlMemoryRepository_Remove(t *testing.T) {
	urlMemoryRepository.Remove(shortUrl)

	if _, err := urlMemoryRepository.FindByShortUrl(shortUrl); err.Error() != NoSuchUrl {
		t.Errorf("Errors does not match %s - %s", NoSuchUrl, err)
	}
}
