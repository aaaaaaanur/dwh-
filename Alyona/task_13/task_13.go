package task_13

import (
	"fmt"
)

func SameNumbers() {
	var number int

	fmt.Scan(&number)

	a := number % 10
	b := (number / 10) % 10
	c := (number / 100) % 10

	if a == b || a == c || b == c {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
