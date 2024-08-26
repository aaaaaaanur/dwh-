package task_10

import (
	"reflect"
	"testing"
)

func TestNewBook(t *testing.T) {

	b := Book{
		ID:        1,
		Name:      "Name",
		Author:    "Author",
		Year:      2024,
		Available: true,
	}

	book := NewBook(b.Name, b.Author, b.Year)

	if b.Name != book.Name {
		t.Error("name is incorrect")
		t.FailNow()
	}
	if b.Author != book.Author {
		t.Error("name is incorrect")
		t.FailNow()
	}
	if b.Year != book.Year {
		t.Error("name is incorrect")
		t.FailNow()
	}
}

func TestNewReader(t *testing.T) {

	r := Reader{
		ID:        1,
		FirstName: "FirstName",
		LastName:  "LastName",
		Age:       20,
		Address:   "Address",
	}

	reader := NewReader(r.FirstName, r.LastName, r.Age, r.Address)

	if reflect.DeepEqual(r, reader) == false {
		t.Errorf("have = %v, need = %v", reader, r)
		t.FailNow()
	}
}

func TestLibrary_AddBook(t *testing.T) {
	l := Library{
		books:   []Book{},
		readers: []Reader{},
	}

	b1 := NewBook("Бла-бла", "Иванов И. И.", 1980)
	b2 := NewBook("Война и мир", "Толстой", 2000)

	l.AddBook(b1)
	l.AddBook(b2)

	if len(l.books) != 2 {
		t.Error("ошибка с добавлением книги")
		t.FailNow()
	}
}

func TestLibrary_DeleteBook(t *testing.T) {
	l := Library{
		books:   []Book{},
		readers: []Reader{},
	}

	b1 := NewBook("Бла-бла", "Иванов И. И.", 1980)
	b2 := NewBook("Война и мир", "Толстой", 2000)

	l.AddBook(b1)
	l.AddBook(b2)
	l.DeleteBook(b2.ID)

	if len(l.books) != 1 {
		t.Error("ошибка с удалением книги")
		t.FailNow()
	}
}

func TestLibrary_AddReader(t *testing.T) {
	l := Library{
		books:   []Book{},
		readers: []Reader{},
	}

	r1 := NewReader("Алёна", "Нурдавлетова", 23, "Комсомольская, 15")
	r2 := NewReader("Сергей", "Чеботарь", 38, "Бехтерева, 16")

	l.AddReader(r1)
	l.AddReader(r2)

	if len(l.readers) != 2 {
		t.Error("ошибка с добавлением читателя")
		t.FailNow()
	}
}

func TestLibrary_DeleteReader(t *testing.T) {
	l := Library{
		books:   []Book{},
		readers: []Reader{},
	}

	r1 := NewReader("Алёна", "Нурдавлетова", 23, "Комсомольская, 15")
	r2 := NewReader("Сергей", "Чеботарь", 38, "Бехтерева, 16")

	l.AddReader(r1)
	l.AddReader(r2)
	l.DeleteReader(r2.ID)

	if len(l.readers) != 1 {
		t.Error("ошибка с удалением читателя")
		t.FailNow()
	}
}

func TestLibrary_UpdateBook(t *testing.T) {
	l := Library{
		books:   []Book{},
		readers: []Reader{},
	}

	b1 := NewBook("Бла-бла", "Иванов И. И.", 1980)
	b2 := NewBook("Война и мир", "Толстой", 2000)

	l.AddBook(b1)
	l.AddBook(b2)
	b1.Author = "Петров П. П."
	l.UpdateBook(b1)

	for _, val := range l.books {
		if val.ID == b1.ID {
			if reflect.DeepEqual(val, b1) == false {
				t.Error("ошибка обновления книги")
			}
		}
	}
}

func TestLibrary_UpdateReader(t *testing.T) {
	l := Library{
		books:   []Book{},
		readers: []Reader{},
	}

	r1 := NewReader("Алена", "Нурдавлетова", 23, "Комсомольская, 15")
	r2 := NewReader("Сергей", "Чеботарь", 38, "Бехтерева, 16")

	l.AddReader(r1)
	l.AddReader(r2)
	r1.FirstName = "Алёна"
	l.UpdateReader(r1)

	for _, val := range l.readers {
		if val.ID == r1.ID {
			if reflect.DeepEqual(val, r1) == false {
				t.Error("ошибка обновления читателя")
			}
		}
	}
}
