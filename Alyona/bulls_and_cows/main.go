package main

import (
	"awesomeProject/Alyona/bulls_and_cows/cmd"
	"awesomeProject/Alyona/bulls_and_cows/ui"
	"log"
)

func main() {
	ui.HelloMess()

	for {
		answer, err := ui.ReadString()
		if err != nil {
			log.Fatal(err)
		}
		switch answer {
		case "y":
			cmd.Play()
		case "n":
			cmd.Exit()
		default:
			ui.ErrorMess2()
			break
		}
	}
}
