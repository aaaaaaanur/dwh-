package task_5

import (
	"fmt"
)

func TaskFive() {
	a := []int{17, 33, 9, 2, 99, 0, 23, 14}
	b := []int{-10, 0, 15, 8000, 33, 99}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Общие значения: ", Ed(a, b))

}

func Ed(z, z1 []int) []int {
	var k []int
	for _, v := range z {
		for _, v1 := range z1 {
			if v == v1 {
				k = append(k, v)
			}
		}
	}
	return k
}
