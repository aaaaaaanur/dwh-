package examples

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //срез, состоящий из элементов типа Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"белый", "черный", "синий", "красный", "желтый"}
	return strings[c]
}

func ExampleFive() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("В наборе имеются %d коробок\n", len(boxes))
	fmt.Println("Объем первой из них равен ", boxes[0].Volume(), "cm³")
	fmt.Println("Цвет последней - ", boxes[len(boxes)-1].color.String())
	fmt.Println("Самая большая из них имеет цвет: ", boxes.BiggestsColor().String())

	fmt.Println("Давайте покрасим их все в черный цвет")
	boxes.PaintItBlack()
	fmt.Println("Цвет второй коробки - ", boxes[1].color.String())

	fmt.Println("Очевидно, что цвет самой большой коробки теперь ", boxes.BiggestsColor().String())
}
