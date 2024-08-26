package task_10

var bookID uint
var readerID uint

type Book struct {
	ID        uint
	Name      string
	Author    string
	Year      int
	Available bool
}

type Library struct {
	books   []Book
	readers []Reader
}

func NewBook(name, author string, year int) Book {
	bookID += 1
	B := Book{bookID, name, author, year, true}
	return B
}

func (l *Library) AddBook(b Book) {
	l.books = append(l.books, b)
}

func (l *Library) UpdateBook(b Book) {
	for i, book := range l.books {
		if book.ID == b.ID {
			l.books[i] = b
		}
	}
}

func (l *Library) DeleteBook(id uint) {
	for i, book := range l.books {
		if book.ID == id && (i+1) < len(l.books) {
			l.books = append(l.books[:i], l.books[i+1])
		} else if book.ID == id && (i+1) == len(l.books) {
			l.books = append(l.books[:i])
		}
	}
}

func (l *Library) FindByYear(year int) []Book {
	var b []Book
	for _, value := range l.books {
		if value.Year == year {
			b = append(b, value)
		}
	}
	return b
}

func (l *Library) FindByAuthor(author string) []Book {
	var a []Book
	for _, value := range l.books {
		if value.Author == author {
			a = append(a, value)
		}
	}
	return a
}

func (l *Library) RentBook(bookID uint) {
	for i, book := range l.books {
		if book.ID == bookID {
			l.books[i].Available = false
		}
	}
}

func (l *Library) ReturnBook(bookID uint) {
	for i, book := range l.books {
		if book.ID == bookID {
			l.books[i].Available = true
		}
	}
}

type Reader struct {
	ID        uint
	FirstName string
	LastName  string
	Age       uint
	Address   string
}

func NewReader(first_name, last_name string, age uint, address string) Reader {
	readerID += 1
	R := Reader{readerID, first_name, last_name, age, address}
	return R
}

func (l *Library) AddReader(r Reader) {
	l.readers = append(l.readers, r)
}

func (l *Library) UpdateReader(r Reader) {
	for i, reader := range l.readers {
		if reader.ID == r.ID {
			l.readers[i] = r
		}
	}
}

func (l *Library) DeleteReader(id uint) {
	for i, reader := range l.readers {
		if reader.ID == id && (i+1) < len(l.readers) {
			l.readers = append(l.readers[:i], l.readers[i+1])
		} else if reader.ID == id && (i+1) == len(l.readers) {
			l.readers = append(l.readers[:i])
		}
	}
}
