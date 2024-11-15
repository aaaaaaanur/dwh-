package ui

import "fmt"

func HelloMess() {
	fmt.Println("Хочешь сыграть в игру \"Быки и коровы\"? (Если да, ответь \"y\", а если нет - \"n\")")
	fmt.Print("->")
}

func RulesMess() {
	fmt.Println("Хорошо, правила таковы:\nЯ задумываю четыре различные цифры из 0,1,2,...9. Ты делаешь ходы, чтобы узнать эти цифры и их порядок.\nКаждый ход состоит из четырёх цифр (0 не может стоять на первом месте).\nВ ответ я показываю число отгаданных цифр, стоящих на своих местах (число быков) и число отгаданных цифр, стоящих не на своих местах (число коров).") //\nЕсли сможешь угадать моё число за 10 ходов, ты победил!")
	fmt.Println("Пример: я задумал 8341.\nТы сделал ход 3401.\nМой ответ: 2 быка и 1 корова.\n(Т.к. ты отгадал 2 цифры (3 и 4), стоящие не на своём месте, и 1 цифру (1), стоящую на своём месте.)")
	fmt.Println("Итак, я загадал число. Твой ответ:")
	fmt.Print("->")
}

func GoodByeMess() {
	fmt.Println("Возвращайся ещё!")
}

func ErrorMess1() {
	fmt.Println("Ошибка. Ход - четырёхзначное число. Начни сначала.")
	fmt.Println("Я загадал новое число. Твой ответ:")
	fmt.Print("->")
}

func ErrorMess2() {
	fmt.Println("Ошибка ввода данных. Начни сначала.")
}
func BullsCowsMess(b, c int) {
	if b == 0 {
		fmt.Print("0 быков, ")
	}
	if b == 1 {
		fmt.Print("1 бык, ")
	}
	if b == 2 || b == 3 || b == 4 {
		fmt.Printf("%d быка, ", b)
	}

	if c == 0 {
		fmt.Print("0 коров\n")
		fmt.Print("->")
	}
	if c == 1 {
		fmt.Print("1 корова\n")
		fmt.Print("->")
	}
	if c == 2 || c == 3 || c == 4 {
		fmt.Printf("%d коровы\n", c)
		fmt.Print("->")
	}
}

func WinMess() {
	fmt.Println("Урааа, ты отгадал число!\nХочешь сыграть ещё раз? (Если да, ответь \"y\", а если нет - \"n\")")
	fmt.Print("->")

	ui.ReadString()
}