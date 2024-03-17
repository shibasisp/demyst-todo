package input

import (
	"demyst-todo/types"
)

type File struct {
	Location string
}

func (f *File) Fetch(limit int) ([]types.TODO, error) {
	return nil, nil
}
