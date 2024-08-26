package examples

import "fmt"

func ExampleOne() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("Срез = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Нечётные элементы среза: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Чётные элементы среза: ", even)
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
