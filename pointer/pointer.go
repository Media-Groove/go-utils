package pointer

import "golang.org/x/exp/constraints"

type primitive interface {
	constraints.Ordered | ~bool
}

func P[T primitive](v T) *T {
	return &v
}
