package structs

import (
	. "fmt"
)

type human struct {
	firstName string
	lastName  string
	age       uint
}

func TaskStruct1() {
	jane := human{"Jane", "Oliver", 47}
	rob := human{"Rob", "Evans", 12}
	karl := human{"Karl", "Johnson", 65}
	ellie := human{"Ellie", "Brown", 24}
	emma := human{"Emma", "Smith", 34}

	slice := []human{jane, rob, karl, ellie, emma}
	Human(slice)
	TwentySixty(slice)
}
func Human(x []human) {
	for _, val := range x {
		Printf("%s %s: %d years old\n", val.firstName, val.lastName, val.age)
	}
}

func TwentySixty(y []human) {
	Println("20-60 years old:")
	for _, val := range y {
		if val.age >= 20 && val.age <= 60 {
			Printf("%s %s\n", val.firstName, val.lastName)
		}
	}
}
