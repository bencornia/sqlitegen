// DO NOT EDIT! GENERATED CODE!
package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Employee struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type EmployeeStore struct {
	db *sql.DB
}

func NewEmployeeStore(db *sql.DB) *EmployeeStore {
	return &EmployeeStore{db: db}
}

func (s *EmployeeStore) GetById(ctx context.Context, id int64) (*Employee, error) {
	query := `
		select	id,
				name,
				created_at,
				updated_at
		from	employee
		where	id = ?;
	`

	var item Employee
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.Name,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *EmployeeStore) UpdateById(ctx context.Context, item *Employee) error {
	query := `
		update	employee
		set		name = ?,
				updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		&item.Name,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) Insert(ctx context.Context, item *Employee) (int64, error) {
	query := `
		insert into employee(
			name
		)
		values (
			?,
			datetime(),
			datetime()
		);
	`

	result, err := s.db.ExecContext(
		ctx,
		query,
		&item.Name,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *EmployeeStore) DeleteById(ctx context.Context, id int64) error {
	query := `
		delete from employee
		where id = ?;
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) DeleteMany(ctx context.Context, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from employee
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) GetByIdTx(ctx context.Context, tx *sql.Tx, id int64) (*Employee, error) {
	query := `
		select	id,
				name,
				created_at,
				updated_at
		from	employee
		where	id = ?;
	`

	var item Employee
	err := tx.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.Name,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *EmployeeStore) UpdateByIdTx(ctx context.Context, tx *sql.Tx, item *Employee) error {
	query := `
		update	employee
		set		name = ?,
				updated_at = datetime()
		where	id = ?;
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		&item.Name,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) InsertTx(ctx context.Context, tx *sql.Tx, item *Employee) (int64, error) {
	query := `
		insert into employee(
			name
		)
		values (
			?,
			datetime(),
			datetime()
		);
	`

	result, err := tx.ExecContext(
		ctx,
		query,
		&item.Name,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *EmployeeStore) DeleteByIdTx(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `
		delete from employee
		where id = ?;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) GetManyTx(ctx context.Context, tx *sql.Tx, ids []int64) ([]*Employee, error) {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		select	id,
				name,
				created_at,
				updated_at
		from	employee
		where	id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	var results []*Employee
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var item Employee
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, &item)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *EmployeeStore) DeleteManyTx(ctx context.Context, tx *sql.Tx, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from employee
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) UpdateManyTx(ctx context.Context, tx *sql.Tx, ids []int64, item *Employee) ([]int64, error) {
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("%d", id)
	}

	query := `
		update	employee
		set		name = ?,
				updated_at = datetime()
		where id in (%s)
		returning id;
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []int64
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		results = append(results, id)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *EmployeeStore) InsertMany(ctx context.Context, tx *sql.Tx, items []*Employee) ([]int64, error) {
	query := `
		insert into employee (
			name,
			created_at,
			updated_at
		) values (			
			?,
			datetime(),
			datetime()
		);
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []int64
	for _, item := range items {
		result, err := stmt.ExecContext(
			ctx,
			&item.Name,
		)
		if err != nil {
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		results = append(results, id)
	}

	return results, nil
}
