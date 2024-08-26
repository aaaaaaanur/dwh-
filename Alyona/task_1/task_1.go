package task_1

import (
	"fmt"
)

func Task() {
	a := []int{17, 33, 9, 2, 99, 0, 23, 14}
	fmt.Print("Чётные чила: ")
	printChet(a)
	fmt.Print("\nНечётные чила: ")
	printNeChet(a)
}

func printChet(s []int) {
	for _, v := range s {
		if v%2 == 0 {
			fmt.Print(v, " ")
		}
	}
}

func printNeChet(s []int) {
	for _, v := range s {
		if v%2 != 0 {
			fmt.Println(v, " ")
		}
	}
}
