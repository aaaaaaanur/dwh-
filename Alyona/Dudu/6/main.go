package main

import "fmt"

func main() {
	a := 10
	b := a
	a = 20

	fmt.Println(a, b)

	c := []int{1, 2, 3}
	d := c
	c[0] = 10
	fmt.Println(c, d)

	i := map[int]int{1: 10, 2: 20, 3: 30}
	g := i
	i[2] = 0

	fmt.Println(i, g)
}
