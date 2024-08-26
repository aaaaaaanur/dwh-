package my_tasks

import (
	"fmt"
	"testing"
)

type C struct {
	F string
}

func TestPassByValue(t *testing.T) {

	a := C{F: "ololo"}
	fmt.Printf("a = %p\n", &a)
	X(a)
	fmt.Printf("a = %p\n", &a)

}

func X(in C) {
	fmt.Printf("a = %p\n", &in)
	in.F = "new value"
}
