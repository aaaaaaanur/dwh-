package ui

import (
	"awesomeProject/Alyona/shopping_list/domain"
	"fmt"
)

func MenuMessage() {
	fmt.Println("Выберите пункт:\n1. Показать списки покупок\n2. Показать товары из списка покупок\n3. Создать список\n4. Добавить товар в список\n5. Удалить список\n6. Выйти")
	fmt.Print("->")
}

func ListMessage(lists []domain.List) {
	fmt.Println("Выберите список:")

	for _, list := range lists {
		fmt.Printf("%d, %s\n", list.ID, list.Name)
	}

	fmt.Println("\n0 - вернуться в главное меню\n100 - выйти")

	fmt.Print("->")
}

func CreateListMessage() {
	fmt.Println("Введите имя списка.")
	fmt.Println("\nb - вернуться в главное меню\ne - выйти")
	fmt.Print("->")
}

func ListCreatedMessage() {
	fmt.Print("Список создан!\n\n")
}
func CreateGoodMessage() {
	fmt.Println("Введите имя товара.")
	fmt.Println("\nb - вернуться в главное меню\ne - выйти")
	fmt.Print("->")
}

func GoodsAddedMessage() {
	fmt.Print("Товар добавлен!\n\n")
}

func BackExitMessage() {
	fmt.Println("\nb - вернуться в главное меню\ne - выйти")
	fmt.Print("->")
}

func ListDeletedMessage() {
	fmt.Print("Список удалён!\n\n")
}

func GoodByeMessage() {
	fmt.Println("До свидания!")
}
