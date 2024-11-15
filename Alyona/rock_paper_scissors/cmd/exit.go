package cmd

import (
	"awesomeProject/Alyona/rock_paper_scissors/ui"
	"os"
)

func Exit() {
	ui.GoodByeMessage()
	os.Exit(0)
}
