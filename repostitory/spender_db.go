package repository

import "database/sql"

type spenderRepository struct {
	db *sql.DB
}

func NewSpenderRepository(db *sql.DB) SpenderRepository {
	return spenderRepository{db: db}
}

func (r spenderRepository) Create(s Spender) (*Spender, error) {
	query := `INSERT INTO spender (name,email) VALUES ($1,$2)`
	result, err := r.db.Exec(query, s.Name, s.Email)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	s.ID = int(id)

	return &s, nil
}

func (r spenderRepository) GetAll() ([]Spender, error) {
	query := `SELECT id, name, email FROM spender`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ss := []Spender{}
	for rows.Next() {
		s := Spender{}
		rows.Scan(&s.ID, &s.Name, &s.Email)
		ss = append(ss, s)
	}

	return ss, nil
}
