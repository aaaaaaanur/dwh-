package domain

type Book struct {
	ID       uint
	AuthorID uint
	Title    string
	//Author   string
	Copies   int
	Borrowed int
}
