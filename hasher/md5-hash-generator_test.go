package hasher

import (
	"testing"
)

var hasher *Md5HashGenerator
var str, hashStr = "hello world", "5eb63b"

func init() {
	hasher = NewMd5HashGenerator()
}

func TestMd5HashGenerator_Generate(t *testing.T) {
	hashText, _ := hasher.Generate(str)

	if hashText != hashStr {
		t.Errorf("hash strings does not match %s - %s", hashStr, hashText)
	}
}
