package repository

import (
	"awesomeProject/Alyona/shopping_list/domain"
	"database/sql"
)

type GoodRepository struct {
	DB *sql.DB
}

func (r GoodRepository) CreateGood(good domain.Good) error {
	_, err := r.DB.Exec("INSERT INTO goods(name, list_id) VALUES ($1, $2)", good.Name, good.ListID)
	if err != nil {
		return err
	}
	return nil
}

func (r GoodRepository) FindAll() ([]domain.Good, error) {
	var goods []domain.Good

	rows, err := r.DB.Query("SELECT * FROM goods")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var good domain.Good

		err := rows.Scan(&good.ID, &good.ListID, &good.Name)
		if err != nil {
			return nil, err
		}
		goods = append(goods, good)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return goods, nil
}

func (r GoodRepository) DeleteGood(id uint) error {
	_, err := r.DB.Exec("DELETE FROM goods WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r GoodRepository) FindList(id uint) ([]domain.Good, error) {
	var goods []domain.Good

	rows, err := r.DB.Query("SELECT * FROM goods WHERE list_id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var good domain.Good

		err := rows.Scan(&good.ID, &good.ListID, &good.Name)
		if err != nil {
			return nil, err
		}
		goods = append(goods, good)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return goods, nil
}
