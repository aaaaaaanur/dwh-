package main

import (
	"awesomeProject/Alyona/db/domain"
	"awesomeProject/Alyona/db/repository"
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:secret@localhost:5432/library?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := repository.BookRepository{DB: db}
	readerRepo := repository.ReaderRepository{DB: db}
	authorRepo := repository.AuthorRepository{DB: db}
	readerbookRepo := repository.ReadersBooksRepository{DB: db}

	for {
		fmt.Println("Введите команду:")
		fmt.Print("-> ")
		text, err := readString()
		if err != nil {
			log.Fatal(err)
		}

		switch text {
		case "quit", "exit", "q":
			log.Println("Good bye, Motherfucker!")
			os.Exit(0)
		case "books":
			count, err := bookRepo.CountBooks()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("В нашей библиотеке %d книг. Введите limit и offset через запятую без пробела:\n", count)
			fmt.Print("->")
			limitOffset, err := readString()
			if err != nil {
				log.Fatal(err)
			}

			parts := strings.Split(limitOffset, ",")
			limit, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			if limit == 0 {
				fmt.Print("Ошибка: limit должен быть больше 0. Начните сначала\n")
				continue
			}
			offset, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			books, err := bookRepo.FindAll(limit, offset)
			if err != nil {
				log.Fatal(err)
			}
			if len(books) == 0 && count > 0 {
				fmt.Printf("Ошибка: offset должен быть меньше %d. Начните сначала\n", count)
				continue
			}
			for _, book := range books {
				fmt.Printf("%d, %s, %s\n", book.ID, book.AuthorID, book.Title)
			}
		case "authors":
			count, err := authorRepo.CountAuthors()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("В нашей библиотеке %d авторов. Введите limit и offset через запятую без пробела:\n", count)
			fmt.Print("->")

			limitOffset, err := readString()
			if err != nil {
				log.Fatal(err)
			}

			parts := strings.Split(limitOffset, ",")
			limit, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			if limit == 0 {
				fmt.Print("Ошибка: limit должен быть больше 0. Начните сначала\n")
				continue
			}

			offset, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}

			authors, err := authorRepo.FindAll(limit, offset)
			if err != nil {
				log.Fatal(err)
			}
			if len(authors) == 0 && count > 0 {
				fmt.Printf("Ошибка: offset должен быть меньше %d. Начните сначала\n", count)
				continue
			}

			for _, author := range authors {
				fmt.Printf("%d, %s, %s\n", author.ID, author.Name)
			}
		case "readers":
			count, err := readerRepo.CountReaders()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("В нашей библиотеке %d читателей. Введите limit и offset через запятую без пробела:\n", count)
			fmt.Print("->")

			limitOffset, err := readString()
			if err != nil {
				log.Fatal(err)
			}

			parts := strings.Split(limitOffset, ",")
			limit, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			if limit == 0 {
				fmt.Print("Ошибка: limit должен быть больше 0. Начните сначала\n")
				continue
			}
			offset, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			readers, err := readerRepo.FindAll(limit, offset)
			if err != nil {
				log.Fatal(err)
			}
			if len(readers) == 0 && count > 0 {
				fmt.Printf("Ошибка: offset должен быть меньше %d. Начните сначала\n", count)
				continue
			}
			for _, reader := range readers {
				fmt.Printf("%d, %s, %s\n", reader.ID, reader.FullName, reader.RegisteredAt)
			}
		case "debts":
			readersbooks, err := readerbookRepo.FindDebts()
			if err != nil {
				log.Fatal(err)
			}
			for _, readerbook := range readersbooks {
				fmt.Printf("%d, %s, %s, %s, %v\n", readerbook.ID, readerbook.TitleBook, readerbook.FullNameReader, readerbook.TakenAt, readerbook.ReturnedAt)
			}
		case "returned":
			readersbooks, err := readerbookRepo.FindReturned()
			if err != nil {
				log.Fatal(err)
			}
			for _, readerbook := range readersbooks {
				fmt.Printf("%d, %s, %s, %s, %s\n", readerbook.ID, readerbook.TitleBook, readerbook.FullNameReader, readerbook.TakenAt, *readerbook.ReturnedAt)
			}
		case "taken":
			readersbooks, err := readerbookRepo.FindTaken()
			if err != nil {
				log.Fatal(err)
			}
			for _, readerbook := range readersbooks {
				var returnedAt string
				if readerbook.ReturnedAt != nil {
					returnedAt = *readerbook.ReturnedAt
				} else {
					returnedAt = "<nil>"
				}
				fmt.Printf("%d, %s, %s, %s, %s\n", readerbook.ID, readerbook.TitleBook, readerbook.FullNameReader, readerbook.TakenAt, returnedAt)
			}
		case "add book":

			//books.Create(bookRepo)

			//book := domain.Book{}
			//
			//fmt.Println("Введите title:")
			//fmt.Print("-> ")
			//
			//book.Title, err = readString()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//
			//fmt.Println("Введите author's id:")
			//fmt.Print("-> ")
			//
			//book.AuthorID, err = readUint()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//
			//fmt.Println("Введите copies:")
			//fmt.Print("-> ")
			//
			//book.Copies, err = readInt()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//
			//fmt.Println("Введите borrowed:")
			//fmt.Print("-> ")

			//book.Borrowed, err = readInt()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//
			//err = bookRepo.CreateBook(book)
			//if err != nil {
			//	log.Fatal(err)
			//}

		case "add author":
			author := domain.Author{}

			fmt.Println("Введите author's name:")
			fmt.Print("-> ")

			author.Name, err = readString()
			if err != nil {
				log.Fatal(err)
			}

			err = authorRepo.CreateAuthor(author)
			if err != nil {
				log.Fatal(err)
			}

		case "add reader":
			read := domain.Reader{}

			fmt.Println("Введите full name:")
			fmt.Print("-> ")

			read.FullName, err = readString()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите date of registration (YYYY-MM-DD):")
			fmt.Print("-> ")

			read.RegisteredAt, err = readString()
			if err != nil {
				log.Fatal(err)
			}

			err = readerRepo.CreateReader(read)
			if err != nil {
				log.Fatal(err)
			}
		case "delete book":
			fmt.Println("Введите id книги, которую хотите удалить:")
			fmt.Print("-> ")

			id, err := readInt()
			if err != nil {
				log.Fatal(err)
			}
			err = bookRepo.DeleteBook(uint(id))
			if err != nil {
				log.Fatal(err)
			}

		case "delete author":
			fmt.Println("Введите id автора, которого хотите удалить:")
			fmt.Print("-> ")
			id, err := readInt()
			if err != nil {
				log.Fatal(err)
			}
			err = authorRepo.DeleteAuthor(uint(id))
			if err != nil {
				log.Fatal(err)
			}
		case "delete reader":
			fmt.Println("Введите id читателя, которого хотите удалить:")
			fmt.Print("-> ")

			id, err := readInt()
			if err != nil {
				log.Fatal(err)
			}
			err = readerRepo.DeleteReader(uint(id))
			if err != nil {
				log.Fatal(err)
			}
		case "update book":
			fmt.Println("Ведите id книги, которую хотите изменить:")
			fmt.Print("-> ")

			id, err := readInt()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите new title:")
			fmt.Print("-> ")

			newTitle, err := readString()
			if err != nil {
				log.Fatal(err)
			}

			err = bookRepo.UpdateBook(uint(id), newTitle)
			if err != nil {
				log.Fatal(err)
			}

		case "update author":
			fmt.Println("Введите id автора, которого хотите изменить:")
			fmt.Print("-> ")

			id, err := readInt()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите new author's name:")
			fmt.Print("-> ")

			newName, err := readString()
			if err != nil {
				log.Fatal(err)
			}
			err = authorRepo.UpdateAuthor(uint(id), newName)
			if err != nil {
				log.Fatal(err)
			}

		case "update reader":
			fmt.Println("Ведите id читателя, которого хотите изменить:")
			fmt.Print("-> ")

			id, err := readInt()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите new name:")
			fmt.Print("-> ")

			newName, err := readString()
			if err != nil {
				log.Fatal(err)
			}

			err = readerRepo.UpdateReader(uint(id), newName)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(str, "\n"), nil
}

func readInt() (int, error) {
	str, err := readString()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}

func readUint() (uint, error) {
	str, err := readString()
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
