package cmd

import (
	"awesomeProject/Alyona/game/ui"
	"os"
)

func Exit() {
	ui.GoodByeMessage()
	os.Exit(0)
}
