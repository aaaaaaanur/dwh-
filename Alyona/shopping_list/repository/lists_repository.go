package repository

import (
	"awesomeProject/Alyona/shopping_list/domain"
	"database/sql"
)

type ListRepository struct {
	DB *sql.DB
}

func (r ListRepository) CreateList(list domain.List) error {
	_, err := r.DB.Exec("INSERT INTO lists(name) VALUES ($1)", list.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r ListRepository) FindAll() ([]domain.List, error) {
	var lists []domain.List

	rows, err := r.DB.Query("SELECT * FROM lists")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var list domain.List

		err := rows.Scan(&list.ID, &list.Name)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r ListRepository) DeleteList(id uint) error {
	_, err := r.DB.Exec("DELETE FROM goods WHERE list_id = $1", id)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec("DELETE FROM lists WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
