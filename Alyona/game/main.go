package main

import (
	"awesomeProject/Alyona/game/cmd"
	"awesomeProject/Alyona/game/ui"
	"log"
)

// 1. Запускаем программу"Сыграем? Я загадаю целое число от 1 до 100, а ты попробуешь угадать, чётное оно или нет. Если хочешь поиграть, ответь "y", если нет - "n"."
// если нет - "До свидания!", выходим из программы, если да - компьютер загадывает число от 1 до 100
// 2. "Отлично! Я загадал своё число. Если думаешь, что оно чётное, ответь "1", если нет - "2", если хочешь выйти из игры - "q"."
// если "q" - выходим из программы; если вводим 1 и загаданное число четное (если вводим 2 и число нечетное) - "ты прав! Моё число - х. попробуешь снова? Ответь "y" (да), чтобы продолжить играть, или "n" (нет), чтобы выйти из игры."
//если n - выходим из программы, если y - цикл повторяется заново.
// 3. если вводим 1 и загаданное число нечетное (либо вводим 2 и загаданное число четное) - "ты проиграл, попробуешь снова? Ответь "y" (да), чтобы продолжить играть, или "n" (нет), чтобы выйти из игры."
//если y - цикл повторяется, если n - выходим из программы

func main() {
	for {
		ui.HelloMessage()

		answer, err := ui.ReadString()
		if err != nil {
			log.Fatal(err)
		}

		switch answer {
		case "n":
			cmd.Exit()
		case "y":
			cmd.Play()
		}
	}
}
