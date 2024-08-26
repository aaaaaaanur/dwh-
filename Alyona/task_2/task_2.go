package task_2

import (
	"fmt"
	"sort"
)

func TaskTwo() {
	a := []int{17, 33, 9, 2, 99, 0, 23, 14}
	fmt.Print("\n")
	SmallBig(a)
	BigSmall(a)
}
func SmallBig(s []int) {
	sort.Ints(s)
	fmt.Println(s)
}
func BigSmall(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
