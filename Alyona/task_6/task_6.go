package task_6

import (
	"fmt"
	"sort"
)

func TaskSix() {
	a := []int{-87, 24, 90, -2, 568, 17, 24, -113, 90}
	b := []int{-2, 967, 3765, 90, 24, 55, 24, 4, -2}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Общие значения: ", Od(a, b))

}

func isEqualValues() bool {

}

func Od(x, y []int) []int {
	var p []int

	sort.Ints(x)
	sort.Ints(y)

	var xi, yi int

	for {
		if xi >= len(x) {
			break
		}
		if yi >= len(y) {
			break
		}
		//
		if x[xi] == y[yi] {

			continue
		}

	}
}
