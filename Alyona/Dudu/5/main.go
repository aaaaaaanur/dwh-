package main

import "fmt"

func Map(x map[string]int) {
	x["b"] = 3
}

func main() {
	v := map[string]int{"a": 1, "b": 2, "c": 3}
	Map(v)
	fmt.Println(v)
}
