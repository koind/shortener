# Shortener
Small library for shortening links.

## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/shortener
 ```

## Usage

Package usage example.

```go
package main

import (
	"github.com/koind/shortener"
	"github.com/koind/shortener/hasher"
	"github.com/koind/shortener/repository"
)

func main() {
	linkShortener := shortener.NewShortener(repository.NewUrlMemoryRepository(), hasher.NewMd5HashGenerator())
	
	fullUrl, err := linkShortener.Shorten("http://site.com/some-long-link")
	if err != nil {
		println(err)
	}
	
	println(fullUrl) // http://site.com/72110a
	
	shortUrl, err := linkShortener.Resolve("http://site.com/72110a")
	if err != nil {
		println(err)
	}
	println(shortUrl) // http://site.com/some-long-link
}
```

## Available Methods

The following methods are available:

##### koind/shortener

```go
NewShortener(repository repository.UrlRepository, hasher hasher.HashGenerator) *LinkShortener
Shorten(url string) (string, error)
Resolve(url string) (string, error)
```

##### koind/shortener/repository

```go
Add(longUrl, shortUrl string) (bool, error)
FindByShortUrl(url string) (string, error)
Remove(url string) (bool, error)
```

##### koind/shortener/hasher

```go
Generate(url string) (string, error)
```

##### koind/shortener/stringer

```go
Substr(input string, start int, length int) string
```

## Tests

Run the following command from you terminal:

```
go test -v .
```