package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func newYear(c *Car) {
	c.Year = 2020
}

func main() {
	car := Car{Brand: "Lexus", Year: 2015}
	fmt.Println(car)

	newYear(&car)
	fmt.Println(car)
}
