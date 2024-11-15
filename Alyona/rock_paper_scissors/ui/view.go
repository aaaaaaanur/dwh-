package ui

import "fmt"

func HelloMessage() {
	fmt.Println("Сыграем в игру \"камень-ножницы-бумага\"? (Если да, ответь \"y\", если нет - \"n\")")
	fmt.Print("->")
}

func GoodByeMessage() {
	fmt.Println("До свидания!")
}

func RulesMessage() {
	fmt.Println("Правила просты: камень бьёт ножницы, бумага бьёт камень, а ножницы - бумагу.\nЕсли хочешь выбрать камень, ответь \"r\", если ножницы - \"s\", а если бумагу - \"p\".\nЯ сделал свой выбор, теперь твоя очередь.")
	fmt.Print("->")
}

func ErrorMessage() {
	fmt.Println("Ошибка ввода данных. Начните сначала.")
}

func YouLostMessage(x string, compWin, userWin int) {
	var w string
	if x == "r" {
		w = "был камень"
	}
	if x == "s" {
		w = "были ножницы"
	}
	if x == "p" {
		w = "была бумага"
	}
	fmt.Printf("Ты проиграл.У меня %s.\nКоличество побед компьютера: %d\nКоличество побед пользователя: %d\nХочешь сыграть ещё? (Ответь \"y\" или \"n\")\n", w, compWin, userWin)
	fmt.Print("->")
}

func YouWonMessage(x string, compWin, userWin int) {
	var w string
	if x == "r" {
		w = "был камень"
	}
	if x == "s" {
		w = "были ножницы"
	}
	if x == "p" {
		w = "была бумага"
	}
	fmt.Printf("Ты выиграл. У меня %s.\nКоличество побед компьютера: %d\nКоличество побед пользователя: %d\nХочешь сыграть ещё? (Ответь \"y\" или \"n\")\n", w, compWin, userWin)
	fmt.Print("->")
}

func AgainMessage() {
	fmt.Println("Ничья. Попробуй ещё раз.")
	fmt.Print("->")
}
