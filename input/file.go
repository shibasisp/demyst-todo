package input

import (
	"demyst-todo/types"
)

type File struct {
	Location string
}

// We can implement more input methods by implementing like this. This is not implemented as it is out of scope of the problem statement.
func (f *File) Fetch(limit int, pattern string) ([]types.TODO, error) {
	return nil, nil
}
