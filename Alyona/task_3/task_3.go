package task_3

import (
	"fmt"
)

func TaskThree() {
	fmt.Println(Sum(-9, 10))
}
func Sum(a, b int) int {
	sum := 0
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}
