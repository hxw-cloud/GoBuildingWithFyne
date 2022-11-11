package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Connect *sql.DB
}

func NewSqLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Connect: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	create table if not exists holdings(
		id integer primary key autoincrement,
		amount real not null,
		purchase_date integer not null,
		purchase_price integer not null);
	`

	_, err := repo.Connect.Exec(query)
	return err
}

func (repo *SQLiteRepository) InsertHolding(holdings Holdings) (*Holdings, error) {
	stmt := "insert into holdings (amount , purchase_date , purchase_price) values (?,?,?)"

	res, err := repo.Connect.Exec(stmt, holdings.Amount, holdings.PurchaseDate.Unix(), holdings.PurchasePrice)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	holdings.ID = id
	return &holdings, nil
}

func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := "select id,amount,purchase_date,purchase_price from holdings order by purchase_date"
	rows, err := repo.Connect.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var all []Holdings
	for rows.Next() {
		var h Holdings
		var unixTime int64
		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h)
	}
	return all, nil
}

func (repo *SQLiteRepository) GetHoldingByID(id int64) (*Holdings, error) {
	row := repo.Connect.QueryRow("select id ,amount,purchase_date,purchase_price from holdings where id = ?", id)

	var h Holdings
	var unixTime int64
	err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)
	if err != nil {
		return nil, err
	}
	h.PurchaseDate = time.Unix(unixTime, 0)
	return &h, nil
}

func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("没有 0 id")
	}
	stmt := "Update holdings set amount = ? ,purchase_date = ? ,purchase_price = ? where id = ?"
	res, err := repo.Connect.Exec(stmt, updated.Amount, updated.PurchaseDate, updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errUpdateFailed
	}
	return nil
}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	res, err := repo.Connect.Exec("delete from holdings where id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errUpdateFailed
	}
	return nil
}
