package input

import "demyst-todo/types"

type Input interface {
	Fetch(limit int) ([]types.TODO, error)
}
