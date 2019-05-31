package hash

type HashGenerator interface {
	Generate(url string) string
	GetError() error
}
