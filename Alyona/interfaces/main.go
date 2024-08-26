//package main
//
//import "fmt"
//
//type Printer interface {
//	Print()
//}
//
//type Animal struct {
//	Nickname string
//}
//
//func (a Animal) Print() {
//	fmt.Println(a.Nickname)
//}
//
//func (a Animal) String() string {
//	return a.Nickname + "!!!!"
//}
//
//type MyByte byte
//
//func (m MyByte) Print() {
//	///
//}
//
//type MyBool bool
//
//func (m MyBool) Print() {
//	///
//}
//
//type MyMap map[string]int
//
//func (m MyMap) Print() {
//	//
//}
//
//type Human struct {
//	FirstName string
//	LastName  string
//	Age       int
//}
//
//func (h Human) Print() {
//	fmt.Println(h.FirstName, h.LastName)
//}
//
//var t MyMap = map[string]int{"waddsf": 7585}
//
//func main() {
//
//	animal := Animal{Nickname: "Pupu"}
//
//	man := Human{
//		FirstName: "John",
//		LastName:  "Smith",
//		Age:       30,
//	}
//
//	myByte := MyByte(1)
//
//	f := false
//	myBool := MyBool(f)
//
//	pepe(myBool)
//
//	var pp Printer
//	pp = myBool
//	pepe(pp)
//
//	s := []Printer{t, man, myByte, myBool}
//
//	a := s[1]
//
//}
//
//func pepe(p Printer) {
//
//}
