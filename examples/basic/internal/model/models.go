// DO NOT EDIT! GENERATED CODE!
package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Company struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CompanyStore struct {
	db *sql.DB
}

func NewCompanyStore(db *sql.DB) *CompanyStore {
	return &CompanyStore{db: db}
}

func (s *CompanyStore) GetById(ctx context.Context, id int64) (*Company, error) {
	query := `
		select	id,
			name,
			created_at,
			updated_at
		from	company
		where	id = ?;
	`

	var item Company
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

func (s *CompanyStore) UpdateById(ctx context.Context, item *Company) error {
	query := `
		update	company
		set	name = ?,
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

func (s *CompanyStore) Insert(ctx context.Context, item *Company) (int64, error) {
	query := `
		insert into company(
			name
		)
		values (?);
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

func (s *CompanyStore) DeleteById(ctx context.Context, id int64) error {
	query := `
		delete from company
		where id = ?;
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStore) GetMany(ctx context.Context, ids []int64) ([]*Company, error) {
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
		from	company
		where	id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	var results []*Company
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var item Company
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

	return results, nil
}

func (s *CompanyStore) UpdateMany(ctx context.Context, ids []int64, item *Company) ([]int64, error) {
	placeholders := make([]string, len(ids))
	idArgs := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		idArgs[i] = id
	}

	query := `
		update	company
		set	name = ?,
		updated_at = datetime()
		where	id in (%s)
		returning id;
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))
	args := append(
		[]any{
			item.Name,
		},
		idArgs...,
	)

	var results []int64
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return results, err
		}

		results = append(results, id)
	}

	return results, nil
}

func (s *CompanyStore) InsertMany(ctx context.Context, items []*Company) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *CompanyStore) DeleteMany(ctx context.Context, ids []int64) error {
	// TOOD: complete body
	return nil
}

type Department struct {
	Id        int64  `json:"id"`
	CompanyId int64  `json:"company_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DepartmentStore struct {
	db *sql.DB
}

func NewDepartmentStore(db *sql.DB) *DepartmentStore {
	return &DepartmentStore{db: db}
}

func (s *DepartmentStore) GetById(ctx context.Context, id int64) (*Department, error) {
	query := `
		select	id,
			company_id,
			name,
			created_at,
			updated_at
		from	department
		where	id = ?;
	`

	var item Department
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.CompanyId,
		&item.Name,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *DepartmentStore) UpdateById(ctx context.Context, item *Department) error {
	query := `
		update	department
		set	company_id = ?,
		name = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		&item.CompanyId,
		&item.Name,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *DepartmentStore) Insert(ctx context.Context, item *Department) (int64, error) {
	query := `
		insert into department(
			company_id,
			name
		)
		values (?, ?);
	`

	result, err := s.db.ExecContext(
		ctx,
		query,
		&item.CompanyId,
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

func (s *DepartmentStore) DeleteById(ctx context.Context, id int64) error {
	query := `
		delete from department
		where id = ?;
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DepartmentStore) GetMany(ctx context.Context, ids []int64) ([]*Department, error) {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		select	id,
			company_id,
			name,
			created_at,
			updated_at
		from	department
		where	id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	var results []*Department
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var item Department
		err = rows.Scan(
			&item.Id,
			&item.CompanyId,
			&item.Name,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, &item)
	}

	return results, nil
}

func (s *DepartmentStore) UpdateMany(ctx context.Context, ids []int64, item *Department) ([]int64, error) {
	placeholders := make([]string, len(ids))
	idArgs := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		idArgs[i] = id
	}

	query := `
		update	department
		set	company_id = ?,
			name = ?,
		updated_at = datetime()
		where	id in (%s)
		returning id;
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))
	args := append(
		[]any{
			item.CompanyId,
			item.Name,
		},
		idArgs...,
	)

	var results []int64
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return results, err
		}

		results = append(results, id)
	}

	return results, nil
}

func (s *DepartmentStore) InsertMany(ctx context.Context, items []*Department) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *DepartmentStore) DeleteMany(ctx context.Context, ids []int64) error {
	// TOOD: complete body
	return nil
}

type Employee struct {
	Id           int64   `json:"id"`
	DepartmentId int64   `json:"department_id"`
	Name         string  `json:"name"`
	Email        *string `json:"email"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
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
			department_id,
			name,
			email,
			created_at,
			updated_at
		from	employee
		where	id = ?;
	`

	var item Employee
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.DepartmentId,
		&item.Name,
		&item.Email,
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
		set	department_id = ?,
		name = ?,
		email = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		&item.DepartmentId,
		&item.Name,
		&item.Email,
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
			department_id,
			name,
			email
		)
		values (?, ?, ?);
	`

	result, err := s.db.ExecContext(
		ctx,
		query,
		&item.DepartmentId,
		&item.Name,
		&item.Email,
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

func (s *EmployeeStore) GetMany(ctx context.Context, ids []int64) ([]*Employee, error) {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		select	id,
			department_id,
			name,
			email,
			created_at,
			updated_at
		from	employee
		where	id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	var results []*Employee
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var item Employee
		err = rows.Scan(
			&item.Id,
			&item.DepartmentId,
			&item.Name,
			&item.Email,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, &item)
	}

	return results, nil
}

func (s *EmployeeStore) UpdateMany(ctx context.Context, ids []int64, item *Employee) ([]int64, error) {
	placeholders := make([]string, len(ids))
	idArgs := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		idArgs[i] = id
	}

	query := `
		update	employee
		set	department_id = ?,
			name = ?,
			email = ?,
		updated_at = datetime()
		where	id in (%s)
		returning id;
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))
	args := append(
		[]any{
			item.DepartmentId,
			item.Name,
			item.Email,
		},
		idArgs...,
	)

	var results []int64
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return results, err
		}

		results = append(results, id)
	}

	return results, nil
}

func (s *EmployeeStore) InsertMany(ctx context.Context, items []*Employee) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *EmployeeStore) DeleteMany(ctx context.Context, ids []int64) error {
	// TOOD: complete body
	return nil
}

type EmployeeSalary struct {
	Id         int64   `json:"id"`
	EmployeeId int64   `json:"employee_id"`
	Amount     float64 `json:"amount"`
	Currency   *string `json:"currency"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type EmployeeSalaryStore struct {
	db *sql.DB
}

func NewEmployeeSalaryStore(db *sql.DB) *EmployeeSalaryStore {
	return &EmployeeSalaryStore{db: db}
}

func (s *EmployeeSalaryStore) GetById(ctx context.Context, id int64) (*EmployeeSalary, error) {
	query := `
		select	id,
			employee_id,
			amount,
			currency,
			created_at,
			updated_at
		from	employee_salary
		where	id = ?;
	`

	var item EmployeeSalary
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.EmployeeId,
		&item.Amount,
		&item.Currency,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *EmployeeSalaryStore) UpdateById(ctx context.Context, item *EmployeeSalary) error {
	query := `
		update	employee_salary
		set	employee_id = ?,
		amount = ?,
		currency = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		&item.EmployeeId,
		&item.Amount,
		&item.Currency,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSalaryStore) Insert(ctx context.Context, item *EmployeeSalary) (int64, error) {
	query := `
		insert into employee_salary(
			employee_id,
			amount,
			currency
		)
		values (?, ?, ?);
	`

	result, err := s.db.ExecContext(
		ctx,
		query,
		&item.EmployeeId,
		&item.Amount,
		&item.Currency,
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

func (s *EmployeeSalaryStore) DeleteById(ctx context.Context, id int64) error {
	query := `
		delete from employee_salary
		where id = ?;
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSalaryStore) GetMany(ctx context.Context, ids []int64) ([]*EmployeeSalary, error) {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		select	id,
			employee_id,
			amount,
			currency,
			created_at,
			updated_at
		from	employee_salary
		where	id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	var results []*EmployeeSalary
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var item EmployeeSalary
		err = rows.Scan(
			&item.Id,
			&item.EmployeeId,
			&item.Amount,
			&item.Currency,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, &item)
	}

	return results, nil
}

func (s *EmployeeSalaryStore) UpdateMany(ctx context.Context, ids []int64, item *EmployeeSalary) ([]int64, error) {
	placeholders := make([]string, len(ids))
	idArgs := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		idArgs[i] = id
	}

	query := `
		update	employee_salary
		set	employee_id = ?,
			amount = ?,
			currency = ?,
		updated_at = datetime()
		where	id in (%s)
		returning id;
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))
	args := append(
		[]any{
			item.EmployeeId,
			item.Amount,
			item.Currency,
		},
		idArgs...,
	)

	var results []int64
	rows, err := s.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return results, err
		}

		results = append(results, id)
	}

	return results, nil
}

func (s *EmployeeSalaryStore) InsertMany(ctx context.Context, items []*EmployeeSalary) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *EmployeeSalaryStore) DeleteMany(ctx context.Context, ids []int64) error {
	// TOOD: complete body
	return nil
}
