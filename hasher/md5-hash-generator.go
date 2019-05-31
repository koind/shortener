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
	err    error
}

func NewMd5HashGenerator() *Md5HashGenerator {
	return &Md5HashGenerator{
		hasher: md5.New(),
	}
}

func (m *Md5HashGenerator) Generate(url string) string {
	if m.err != nil {
		return ""
	}

	if url == "" {
		m.err = errors.New(repository.EmptyUrlError)
		return ""
	}

	m.hasher.Write([]byte(url))

	return hex.EncodeToString(m.hasher.Sum(nil))
}

func (m *Md5HashGenerator) GetError() error {
	return m.err
}
