package cmd

import (
	"awesomeProject/Alyona/bulls_and_cows/ui"
	"os"
)

func Exit() {
	ui.GoodByeMess()
	os.Exit(0)
}
