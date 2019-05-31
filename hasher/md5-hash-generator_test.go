package hasher

import (
	"github.com/koind/shortener/repository"
	"testing"
)

var hasher *Md5HashGenerator
var str, hashStr = "hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3"

func init() {
	hasher = NewMd5HashGenerator()
}

func TestMd5HashGenerator_Generate(t *testing.T) {
	hashText := hasher.Generate(str)

	if hashText != hashStr {
		t.Errorf("hash strings does not match %s - %s", hashStr, hashText)
	}
}

func TestMd5HashGenerator_GetError(t *testing.T) {
	hasher.Generate("")

	if err := hasher.GetError(); err.Error() != repository.EmptyUrlError {
		t.Errorf("Errors does not match %s - %s", repository.EmptyUrlError, err)
	}
}
