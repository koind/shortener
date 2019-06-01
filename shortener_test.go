package shortener

import (
	"github.com/koind/shortener/hasher"
	"github.com/koind/shortener/repository"
	"testing"
)

var shortener *LinkShortener
var fullUrl, shortUrl = "http://site.com/some-long-link", "http://site.com/72110a"

func init() {
	shortener = NewShortener(repository.NewUrlMemoryRepository(), hasher.NewMd5HashGenerator())
}

func TestLinkShortener_Shorten(t *testing.T) {
	url, _ := shortener.Shorten(fullUrl)

	if url != shortUrl {
		t.Errorf("links does not match %s - %s", shortUrl, url)
	}
}

func TestLinkShortener_Resolve(t *testing.T) {
	url, _ := shortener.Resolve(shortUrl)

	if url != fullUrl {
		t.Errorf("links does not match %s - %s", fullUrl, url)
	}

	url, _ = shortener.Resolve(shortUrl + "-fails")

	if url != "" {
		t.Error("url must be empty")
	}
}
