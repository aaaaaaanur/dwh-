package main

import "fmt"

func main() {
	var x *int

	fmt.Println(x)

	y := 20

	x = &y

	if x != nil {
		fmt.Println(*x)
	}
	if x == nil {
		fmt.Println(x)
	}
}
