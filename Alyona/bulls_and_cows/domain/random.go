package domain

import (
	"math/rand"
)

func Random() (int, int, int, int) {
	var (
		x1 int
		x2 int
		x3 int
		x4 int
	)

	for {
		x1 = rand.Intn(10)
		if x1 != 0 {
			break
		}
	}

	for {
		x2 = rand.Intn(10)
		if x2 != x1 {
			break
		}
	}

	for {
		x3 = rand.Intn(10)
		if x3 != x1 && x3 != x2 {
			break
		}
	}

	for {
		x4 = rand.Intn(10)
		if x4 != x1 && x4 != x2 && x4 != x3 {
			break
		}
	}

	return x1, x2, x3, x4
}
