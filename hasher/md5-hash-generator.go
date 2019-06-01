package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/koind/shortener/repository"
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

func (m *Md5HashGenerator) Generate(url string) (string, error) {
	if url == "" {
		return "", errors.New(repository.EmptyUrlError)
	}

	m.hasher.Write([]byte(url))

	return hex.EncodeToString(m.hasher.Sum(nil)), nil
}
