package task_4

import (
	"fmt"
)

func TaskFour() {
	Hun(1, 100)
}
func Hun(x, y uint) {
	for i := x; i <= y; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		} else if i%3 == 0 && i%5 != 0 {
			i := "Fizz"
			fmt.Println(i)
		} else if i%3 != 0 && i%5 == 0 {
			i := "Buzz"
			fmt.Println(i)
		} else {
			i := "FizzBuzz"
			fmt.Println(i)
		}
	}
}
