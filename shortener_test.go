package shortener

import (
	"github.com/koind/shortener/hasher"
	"github.com/koind/shortener/repository"
	"testing"
)

var shortener *LinkShortener
var longUrl, shortUrl = "http://site.com/some-long-link", "http://site.com/72110a"

func init() {
	shortener = NewShortener(repository.NewUrlMemoryRepository(), hasher.NewMd5HashGenerator())
}

func TestLinkShortener_Shorten(t *testing.T) {
	url := shortener.Shorten(longUrl)

	if url != shortUrl {
		t.Errorf("links does not match %s - %s", shortUrl, url)
	}
}

func TestLinkShortener_Resolve(t *testing.T) {
	url := shortener.Resolve(shortUrl)

	if url != longUrl {
		t.Errorf("links does not match %s - %s", longUrl, url)
	}

	url = shortener.Resolve(shortUrl + "-fails")

	if url != "" {
		t.Error("url must be empty")
	}
}

func TestLinkShortener_GetError(t *testing.T) {
	shortener.Resolve("")

	if err := shortener.GetError(); err.Error() != repository.EmptyUrlError {
		t.Errorf("Errors does not match %s - %s", repository.EmptyUrlError, err)
	}
}
