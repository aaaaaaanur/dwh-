package cmd

import (
	"awesomeProject/Alyona/game/domain"
	"awesomeProject/Alyona/game/ui"
	"log"
	"math/rand"
)

func Play() {
	ui.RulesMessage()

	x := rand.Intn(101)

	str, err := ui.ReadString()
	if err != nil {
		log.Fatal(err)
	}

	var guess bool

	switch str {
	case "q":
		Exit()
	case "1":
		guess = true
	case "2":
		guess = false
	default:
		ui.ErrorMessage()
		return
	}

	if domain.Random() == guess {
		ui.YouWonMessage(x)
	} else {
		ui.YouLostMessage(x)
	}

	Continue()
}

func Continue() {
	str, err := ui.ReadString()
	if err != nil {
		log.Fatal(err)
	}

	switch str {
	case "n":
		Exit()
	case "y":
		return
	default:
		ui.ErrorMessage()
	}
}
