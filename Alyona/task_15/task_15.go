package task_15

import (
	"fmt"
)

func NumberOfRepet() {
	var x int
	var y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	if x > y {
		return
	}

	var slice []int

	for i := x; i <= y; i++ {
		ii := i
		var a []int
		for {
			a = append(a, ii%10)
			ii = ii / 10
			if ii == 0 {
				break
			}
		}
		for index, value := range a {
			var exit bool
			for idx, val := range a {
				if value == val && index != idx {
					slice = append(slice, i)
					exit = true
					break
				}
			}
			if exit {
				break
			}
		}
	}
	fmt.Println(len(slice))
}
