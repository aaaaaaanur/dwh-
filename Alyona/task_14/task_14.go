package task_14

import (
	"fmt"
)

func Sum() {
	var number int
	fmt.Scan(&number)

	sum := 0

	for {
		sum += number % 10
		number = number / 10

		if number == 0 {
			break
		}
	}
	fmt.Println(sum)
}
