package repository

import (
	"awesomeProject/Alyona/db/domain"
	"database/sql"
)

type BookRepository struct {
	DB *sql.DB
}

func (r BookRepository) CountBooks() (int, error) {
	rows, err := r.DB.Query("SELECT COUNT(*) FROM books")
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var count int

	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (r BookRepository) FindAll(limit, offset int) ([]domain.Book, error) {
	rows, err := r.DB.Query("SELECT * FROM books limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.AuthorID, &book.Title, &book.Copies, &book.Borrowed)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r BookRepository) CreateBook(book domain.Book) error {
	_, err := r.DB.Exec("INSERT INTO books (author_id, title, copies, borrowed) VALUES ($1, $2, $3, $4)", book.AuthorID, book.Title, book.Copies, book.Borrowed)
	if err != nil {
		return err
	}
	return nil
}

func (r BookRepository) DeleteBook(id uint) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r BookRepository) UpdateBook(id uint, newTitle string) error {
	_, err := r.DB.Exec("UPDATE books SET title = $1 WHERE id = $2", newTitle, id)
	if err != nil {
		return err
	}
	return nil
}
