package repository

import (
	"awesomeProject/Alyona/db/domain"
	"database/sql"
)

type ReaderRepository struct {
	DB *sql.DB
}

func (r ReaderRepository) CountReaders() (int, error) {
	rows, err := r.DB.Query("SELECT COUNT(*) FROM reader")
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

func (r ReaderRepository) FindAll(limit, offset int) ([]domain.Reader, error) {
	rows, err := r.DB.Query("SELECT * FROM reader limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readers []domain.Reader
	for rows.Next() {
		var reader domain.Reader
		err := rows.Scan(&reader.ID, &reader.FullName, &reader.RegisteredAt)
		if err != nil {
			return nil, err
		}
		readers = append(readers, reader)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return readers, nil
}

func (r ReaderRepository) CreateReader(reader domain.Reader) error {
	_, err := r.DB.Exec("INSERT INTO reader (full_name, registered_at) VALUES ($1, $2)", reader.FullName, reader.RegisteredAt)
	if err != nil {
		return err
	}
	return nil
}

func (r ReaderRepository) DeleteReader(id uint) error {
	_, err := r.DB.Exec("DELETE FROM reader WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r ReaderRepository) UpdateReader(id uint, newName string) error {
	_, err := r.DB.Exec("UPDATE reader SET full_name = $1 WHERE id = $2", newName, id)
	if err != nil {
		return err
	}
	return nil
}
