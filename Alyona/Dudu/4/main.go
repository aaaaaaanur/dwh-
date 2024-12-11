package main

import "fmt"

func slice(x []int) {
	x[1] = 3
}

func main() {
	a := []int{1, 2, 3}

	slice(a)
	fmt.Println(a)
}
