package task_7

import "fmt"

func TaskSeven() {
	Ch(7, 13)
	Hc(7, 13)
}

func Ch(a, b int) {
	for i := a; i <= b; i += 2 {
		fmt.Println(i)
	}
}
func Hc(a, b int) {
	for i := b; i >= a; i -= 2 {
		fmt.Println(i)
	}
}
