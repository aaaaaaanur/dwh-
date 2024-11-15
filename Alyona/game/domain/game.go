package domain

import "math/rand"

func Random() bool {
	x := rand.Intn(101)
	if x%2 == 0 {
		return true
	}
	return false
}
