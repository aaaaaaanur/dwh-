package repository

import (
	"awesomeProject/Alyona/db/domain"
	"database/sql"
)

type AuthorRepository struct {
	DB *sql.DB
}

func (r AuthorRepository) CountAuthors() (int, error) {
	rows, err := r.DB.Query("SELECT COUNT(*) FROM authors")
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

func (r AuthorRepository) FindAll(limit, offset int) ([]domain.Author, error) {
	rows, err := r.DB.Query("SELECT * FROM authors limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []domain.Author

	for rows.Next() {
		var author domain.Author
		err := rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r AuthorRepository) CreateAuthor(author domain.Author) error {
	_, err := r.DB.Exec("INSERT INTO authors(name) VALUES ($1)", author.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r AuthorRepository) DeleteAuthor(id uint) error {
	_, err := r.DB.Exec("DELETE FROM authors WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r AuthorRepository) UpdateAuthor(id uint, newName string) error {
	_, err := r.DB.Exec("UPDATE authors SET name = $1 WHERE id = $2", newName, id)
	if err != nil {
		return err
	}
	return nil
}
