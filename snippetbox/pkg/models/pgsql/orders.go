package pgsql

import (
	"database/sql"
	"woonbeaj/snippetbox/pkg/models"
)

type OrderModel struct {
	DB *sql.DB
}

func (m *OrderModel) Insert(surname, name string, amount uint) (int, error) {
	var id int
	m.DB.QueryRow("insert into cf_orders (surname, name, amount, order_date) values ($1, $2, $3, current_timestamp at time zone ('utc')) returning  id", surname, name, amount).Scan(&id)
	if id == 0 {
		return 0, nil
	}
	return id, nil
}

func (m *OrderModel) Get(id int) (*models.Order, error) {
	s := &models.Order{}
	err := m.DB.QueryRow("select * from cf_orders WHERE id = $1", id).Scan(
		&s.ID, &s.Surname, &s.Name, &s.Amount, &s.Order_date)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *OrderModel) Latest() ([]*models.Order, error) {
	return nil, nil
}