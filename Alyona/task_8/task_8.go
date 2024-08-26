package task_8

import "fmt"

func TaskEight() {
	a := []int{-87, 24, 90, -2, 568, 17, 24, -113, 90, -113}
	b := []int{-2, 967, 3765, 90, 24, 55, 24, 4, -2}
	fmt.Println(Differ(a, b))
	fmt.Println(Differ1(a, b))

}

func Differ(x, y []int) []int {
	p := diff(x, y)
	p2 := diff(y, x)
	p = append(p, p2...)

	return p

}

func Differ1(x, y []int) []int {
	k := diff1(x, y)
	k2 := diff1(y, x)
	k = append(k, k2...)

	return k
}

func diff(x, y []int) []int {
	var p []int

	for _, xv := range x {
		var isEqual bool
		for _, yv := range y {
			if xv == yv {
				isEqual = true
			}
		}
		if isEqual == false {
			p = append(p, xv)
		}
	}

	return p
}

func diff1(x, y []int) []int {
	var k []int

	for _, xv := range x {
		if xv < 1 {
			continue
		}
		var isEqual bool
		for _, yv := range y {
			if yv < 1 {
				continue
			} else if xv == yv {
				isEqual = true
			}
		}
		if isEqual == false {
			k = append(k, xv)
		}
	}

	return k
}
