package hash

type Md5HashGenerator struct {
	err error
}

func NewMd5HashGenerator() *Md5HashGenerator {
	return new(Md5HashGenerator)
}

func (m *Md5HashGenerator) Generate(url string) string {
	return ""
}

func (m *Md5HashGenerator) GetError() error {
	return m.err
}
