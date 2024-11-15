package cmd

import (
	"awesomeProject/Alyona/bulls_and_cows/domain"
	"awesomeProject/Alyona/bulls_and_cows/ui"
	"log"
)

func Play() {
	ui.RulesMess()

	var (
		count1 int
		count2 int
	)

	x1, x2, x3, x4 := domain.Random()

	for {
		ans, err := ui.ReadInt()
		if err != nil {
			log.Fatal(err)
		}

		if ans > 9999 {
			ui.ErrorMess1()
			continue
		}

		a1 := ans / 1000
		a2 := (ans / 100) % 10
		a3 := (ans % 100) / 10
		a4 := ans % 10

		if a1 >= 10 || a1 < 1 {
			ui.ErrorMess1()
			continue
		}

		sliceA := []int{a1, a2, a3, a4}
		sliceX := []int{x1, x2, x3, x4}

		if a1 == x1 && a2 == x2 && a3 == x3 && a4 == x4 {
			ui.WinMess()
			break
		}

		count1 = 0
		count2 = 0

		if a1 == x1 {
			count1 += 1
		}

		if a2 == x2 {
			count1 += 1
		}

		if a3 == x3 {
			count1 += 1
		}

		if a4 == x4 {
			count1 += 1
		}

		for _, v := range sliceA {
			for _, v2 := range sliceX {
				if v == v2 {
					count2 += 1
				}
			}
		}
		count2 = count2 - count1
		ui.BullsCowsMess(count1, count2)
		continue
	}
}
