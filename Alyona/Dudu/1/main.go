package main

import "fmt"

func modifyValue(x *int) {
	*x = 100
}
func main() {
	var a int

	a = 20
	fmt.Println(a)

	modifyValue(&a)
	fmt.Println(a)
}
