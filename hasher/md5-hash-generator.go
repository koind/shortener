package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/koind/shortener/repository"
	"github.com/koind/shortener/stringer"
	"hash"
)

type Md5HashGenerator struct {
	hasher hash.Hash
}

func NewMd5HashGenerator() *Md5HashGenerator {
	return &Md5HashGenerator{
		hasher: md5.New(),
	}
}

func (m *Md5HashGenerator) Generate(text string) (string, error) {
	if text == "" {
		return "", errors.New(repository.EmptyUrlError)
	}

	m.hasher.Write([]byte(text))
	urlHash := hex.EncodeToString(m.hasher.Sum(nil))
	urlHash = stringer.Substr(urlHash, 0, 6)

	return urlHash, nil
}
