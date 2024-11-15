package cmd

import (
	"awesomeProject/Alyona/rock_paper_scissors/domain"
	"awesomeProject/Alyona/rock_paper_scissors/ui"
	"errors"
	"log"
)

var (
	Count1 int
	Count2 int
)

func Play() {
	ui.RulesMessage()

	for {
		a := domain.Random()

		str, err := ui.ReadString()
		if err != nil {
			log.Fatal(err)
		}

		if err = Validate(str); err != nil {
			ui.ErrorMessage()
			return
		}

		if a == str {
			ui.AgainMessage()
			continue
		}

		if a == "r" && str == "s" || a == "s" && str == "p" || a == "p" && str == "r" {
			Count1 += 1
			ui.YouLostMessage(a, Count1, Count2)
			return
		}
		if a == "s" && str == "r" || a == "p" && str == "s" || a == "r" && str == "p" {
			Count2 += 1
			ui.YouWonMessage(a, Count1, Count2)
			return
		}

	}
}

func Validate(a string) error {
	switch a {
	case "r", "s", "p":
		return nil
	default:
		return errors.New("invalid command")
	}
}
