package repository

import (
	"awesomeProject/Alyona/db/domain"
	"database/sql"
)

type ReadersBooksRepository struct {
	DB *sql.DB
}

func (r ReadersBooksRepository) FindDebts() ([]domain.ReadersBooks, error) {
	rows, err := r.DB.Query("SELECT readers_books.id, book.title, reader.full_name, taken_at, returned_at FROM readers_books INNER JOIN book ON readers_books.book_id = book.id INNER JOIN reader ON readers_books.reader_id = reader.id WHERE returned_at IS NULL")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var readersbooks []domain.ReadersBooks

	for rows.Next() {
		var readerbook domain.ReadersBooks

		var returnedAt sql.NullString
		err := rows.Scan(&readerbook.ID, &readerbook.TitleBook, &readerbook.FullNameReader, &readerbook.TakenAt, &returnedAt)
		if err != nil {
			return nil, err
		}

		if returnedAt.Valid {
			readerbook.ReturnedAt = &returnedAt.String
		}

		readersbooks = append(readersbooks, readerbook)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return readersbooks, nil
}

func (r ReadersBooksRepository) FindReturned() ([]domain.ReadersBooks, error) {
	rows, err := r.DB.Query("SELECT readers_books.id, book.title, reader.full_name, taken_at, returned_at FROM readers_books INNER JOIN book ON readers_books.book_id = book.id INNER JOIN reader ON readers_books.reader_id = reader.id WHERE returned_at NOTNULL")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var readersbooks []domain.ReadersBooks

	for rows.Next() {
		var readerbook domain.ReadersBooks

		var returnedAt sql.NullString

		err := rows.Scan(&readerbook.ID, &readerbook.TitleBook, &readerbook.FullNameReader, &readerbook.TakenAt, &returnedAt)
		if err != nil {
			return nil, err
		}

		if returnedAt.Valid {
			readerbook.ReturnedAt = &returnedAt.String
		}

		readersbooks = append(readersbooks, readerbook)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return readersbooks, nil
}

func (r ReadersBooksRepository) FindTaken() ([]domain.ReadersBooks, error) {
	rows, err := r.DB.Query("SELECT readers_books.id, book.title, reader.full_name, taken_at, returned_at FROM readers_books INNER JOIN book ON readers_books.book_id = book.id INNER JOIN reader ON readers_books.reader_id = reader.id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var readersbooks []domain.ReadersBooks

	for rows.Next() {
		var readerbook domain.ReadersBooks

		var returnedAt sql.NullString

		err := rows.Scan(&readerbook.ID, &readerbook.TitleBook, &readerbook.FullNameReader, &readerbook.TakenAt, &returnedAt)
		if err != nil {
			return nil, err
		}

		if returnedAt.Valid {
			readerbook.ReturnedAt = &returnedAt.String
		}

		readersbooks = append(readersbooks, readerbook)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return readersbooks, nil
}
