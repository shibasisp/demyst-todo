package input

type Input interface {
	Fetch(int) ([]byte, error)
}
