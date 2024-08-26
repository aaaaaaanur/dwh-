package examples

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // struct как встраиваемое поле
	Skills    // срез из строк как встраиваемое поле
	int       // встроенный тип как встраиваемое поле
	specialty string
}

func ExampleThree() {
	// Инициализируем студента Джейн
	jane := Student{Human: Human{"Джейн", 35, 100}, specialty: "Биология"}
	// доступ к полям
	fmt.Println("Ее имя: ", jane.name)
	fmt.Println("Ее возраст: ", jane.age)
	fmt.Println("Ее масса: ", jane.weight)
	fmt.Println("Ее специализация: ", jane.specialty)
	// изменяем поле навыков
	jane.Skills = []string{"анатомия"}
	fmt.Println("Ее навыки: ", jane.Skills)
	fmt.Println("Она овладела еще двумя навыками: ")
	jane.Skills = append(jane.Skills, "физика", "golang")
	fmt.Println("Теперь ее навыки: ", jane.Skills)
	// изменяем встраиваемое поле
	jane.int = 3
	fmt.Println("Ее любимое число: ", jane.int)
}
