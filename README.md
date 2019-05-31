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
	
	linkShortener.Shorten("http://site.com/some-long-link") // http://site.com/72110a
	linkShortener.Resolve("http://site.com/72110a") // http://site.com/some-long-link
	
	if err := linkShortener.GetError(); err != nil {
		println(err)
	}
}
```

## Available Methods

The following methods are available:

##### koind/shortener

```go
NewShortener(repository repository.UrlRepository, hasher hasher.HashGenerator) *LinkShortener
Shorten(longUrl string) string
Resolve(shortUrl string) string
GetError() error
```

##### koind/shortener/repository

```go
Add(longUrl, shortUrl string) bool
FindByShortUrl(url string) string
FindByLongUrl(url string) string
Remove(url string) bool
GetError() error
```

##### koind/shortener/hasher

```go
Generate(url string) string
GetError() error
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