package input

import "os"

type File struct {
	Location string
}

func (f *File) Fetch(limit int) ([]byte, error) {
	return os.ReadFile(f.Location)
}
