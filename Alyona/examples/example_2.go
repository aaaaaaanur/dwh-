package examples

import "fmt"

type Person struct {
	name string
	age  int
}

func ExampleTwo() {
	var tom Person
	tom.name, tom.age = "Tom", 18
	bob := Person{age: 25, name: "Bob"}
	paul := Person{"Paul", 43}
	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Из %s и %s %s старше на %d лет\n", tom.name, bob.name, tb_Older.name, tb_diff)
	fmt.Printf("Из %s и %s %s старше на %d лет\n", tom.name, paul.name, tp_Older.name, tp_diff)
	fmt.Printf("Из %s и %s %s старше на %d лет\n", bob.name, paul.name, bp_Older.name, bp_diff)
}

func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}
