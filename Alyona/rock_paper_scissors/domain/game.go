package domain

import "math/rand"

func Random() string {
	x := rand.Intn(3)
	if x == 0 {
		return "r"
	}
	if x == 1 {
		return "s"
	}
	return "p"
}
