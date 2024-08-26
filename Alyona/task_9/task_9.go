package task_9

import (
	"fmt"
)

func TaskNine() {
	a := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{26, 27, 28, 29, 30, 31, 32, 33, 34, 35}
	MoySladkiy(a, b)
}

func MoySladkiy(x, y []int) {
	var p []int
	var k []int
	for _, val2 := range y {
		if (val2 % 2) == 0 {
			p = append(p, val2)
		} else {
			k = append(k, val2)
		}
	}
	for i1, _ := range x {
		if (i1 % 2) == 0 {
			fmt.Println(x[i1])
			fmt.Println(p)
		} else {
			fmt.Println(x[i1])
			fmt.Println(k)
		}
	}
}
