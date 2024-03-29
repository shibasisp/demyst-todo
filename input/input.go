package input

import "demyst-todo/types"

type Input interface {
	Fetch(limit int, pattern string) ([]types.TODO, error)
}

// We can extend this to add more formulas whenever needed without interfering with the inputs
var FormulaMap = map[string]func(int) int{
	"even": func(id int) int { return 2 * id },
	"odd":  func(id int) int { return 2*id - 1 },
	"all":  func(id int) int { return id },
	"":     func(id int) int { return id },
}
