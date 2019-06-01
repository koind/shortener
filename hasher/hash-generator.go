package hasher

type HashGenerator interface {
	Generate(url string) (string, error)
}
