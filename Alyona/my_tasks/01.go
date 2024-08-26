package my_tasks

import "fmt"

func MyTask01() {
	slice := []string{"B", "F", "D", "A", "G", "C", "E"}
	Cycle(slice)
}

func Cycle(x []string) {
	i := 0
	for {
		fmt.Println(x[i])
		i++
		if i >= len(x) {
			break
		}
	}
}
