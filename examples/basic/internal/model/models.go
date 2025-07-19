package model

import "database/sql"

// DO NOT EDIT! THIS IS GENERATED CODE!

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

type Employee struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func (s *Store) GetEmployee() ([]*Employee, error) {
	var items []*Employee
	query := `
		select id, first_name, last_name, age, updated_at, created_at
		from employee;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return items, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var item Employee

		err = rows.Scan(
			&item.Id,
			&item.FirstName,
			&item.LastName,
			&item.Age,
			&item.UpdatedAt,
			&item.CreatedAt,
		)

		if err != nil {
			return items, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (s *Store) GetEmployeeById(id int64) (*Employee, error) {
	var item Employee
	query := `
		select id, first_name, last_name, age, updated_at, created_at
		from employee
		where id = ?;
	`

	err := s.db.QueryRow(query, id).Scan(
		&item.Id,
		&item.FirstName,
		&item.LastName,
		&item.Age,
		&item.UpdatedAt,
		&item.CreatedAt,
	)

	if err != nil {
		return &item, err
	}

	return &item, nil
}

func (s *Store) InsertEmployee(item *Employee) (*Employee, error) {
	query := `
		insert into employee(first_name, last_name, age, updated_at, created_at)
		values (?, ?, ?, datetime(), datetime());
	`

	result, err := s.db.Exec(
		query,
		&item.FirstName,
		&item.LastName,
		&item.Age,
	)

	if err != nil {
		return item, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}

	return s.GetEmployeeById(id)
}

func (s *Store) UpdateEmployee(item *Employee) (*Employee, error) {
	query := `
		update employee
		set first_name = ?, last_name = ?, age = ?, updated_at = datetime()
	`

	_, err := s.db.Exec(
		query,
		&item.FirstName,
		&item.LastName,
		&item.Age,
	)

	if err != nil {
		return item, err
	}

	return s.GetEmployeeById(item.Id)
}

func (s *Store) DeleteEmployee(id int) error {
	query := `
		delete from employee
		where id = ?
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
