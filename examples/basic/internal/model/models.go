// DO NOT EDIT! GENERATED CODE!
package model

import (
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

func (s *CompanyStore) Get(id int64) (*Company, error) {
	query := `
		select	id,
			name,
			created_at,
			updated_at
		from	company
		where	id = ?;
	`

	var item Company
	err := s.db.QueryRow(query, id).Scan(
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

func (s *CompanyStore) Update(item *Company) error {
	query := `
		update	company
		set	name = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.Exec(
		query,
		&item.Name,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStore) Insert(item *Company) (int64, error) {
	query := `
		insert into company(
			name
		)
		values (?);
	`

	result, err := s.db.Exec(
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

func (s *CompanyStore) Delete(id int64) error {
	query := `
		delete from company
		where id = ?;
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStore) GetMany(ids []int64) ([]*Company, error) {
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
	rows, err := s.db.Query(query, args...)
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

func (s *CompanyStore) UpdateMany(items []*Company) error {
	// TOOD: complete body
	return nil
}

func (s *CompanyStore) InsertMany(items []*Company) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *CompanyStore) DeleteMany(ids int64) error {
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

func (s *DepartmentStore) Get(id int64) (*Department, error) {
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
	err := s.db.QueryRow(query, id).Scan(
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

func (s *DepartmentStore) Update(item *Department) error {
	query := `
		update	department
		set	company_id = ?,
			name = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.Exec(
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

func (s *DepartmentStore) Insert(item *Department) (int64, error) {
	query := `
		insert into department(
			company_id,
			name
		)
		values (?, ?);
	`

	result, err := s.db.Exec(
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

func (s *DepartmentStore) Delete(id int64) error {
	query := `
		delete from department
		where id = ?;
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DepartmentStore) GetMany(ids []int64) ([]*Department, error) {
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
	rows, err := s.db.Query(query, args...)
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

func (s *DepartmentStore) UpdateMany(items []*Department) error {
	// TOOD: complete body
	return nil
}

func (s *DepartmentStore) InsertMany(items []*Department) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *DepartmentStore) DeleteMany(ids int64) error {
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

func (s *EmployeeStore) Get(id int64) (*Employee, error) {
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
	err := s.db.QueryRow(query, id).Scan(
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

func (s *EmployeeStore) Update(item *Employee) error {
	query := `
		update	employee
		set	department_id = ?,
			name = ?,
			email = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.Exec(
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

func (s *EmployeeStore) Insert(item *Employee) (int64, error) {
	query := `
		insert into employee(
			department_id,
			name,
			email
		)
		values (?, ?, ?);
	`

	result, err := s.db.Exec(
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

func (s *EmployeeStore) Delete(id int64) error {
	query := `
		delete from employee
		where id = ?;
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeStore) GetMany(ids []int64) ([]*Employee, error) {
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
	rows, err := s.db.Query(query, args...)
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

func (s *EmployeeStore) UpdateMany(items []*Employee) error {
	// TOOD: complete body
	return nil
}

func (s *EmployeeStore) InsertMany(items []*Employee) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *EmployeeStore) DeleteMany(ids int64) error {
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

func (s *EmployeeSalaryStore) Get(id int64) (*EmployeeSalary, error) {
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
	err := s.db.QueryRow(query, id).Scan(
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

func (s *EmployeeSalaryStore) Update(item *EmployeeSalary) error {
	query := `
		update	employee_salary
		set	employee_id = ?,
			amount = ?,
			currency = ?,
			updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.Exec(
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

func (s *EmployeeSalaryStore) Insert(item *EmployeeSalary) (int64, error) {
	query := `
		insert into employee_salary(
			employee_id,
			amount,
			currency
		)
		values (?, ?, ?);
	`

	result, err := s.db.Exec(
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

func (s *EmployeeSalaryStore) Delete(id int64) error {
	query := `
		delete from employee_salary
		where id = ?;
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSalaryStore) GetMany(ids []int64) ([]*EmployeeSalary, error) {
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
	rows, err := s.db.Query(query, args...)
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

func (s *EmployeeSalaryStore) UpdateMany(items []*EmployeeSalary) error {
	// TOOD: complete body
	return nil
}

func (s *EmployeeSalaryStore) InsertMany(items []*EmployeeSalary) ([]int64, error) {
	// TOOD: complete body
	var results []int64
	return results, nil
}

func (s *EmployeeSalaryStore) DeleteMany(ids int64) error {
	// TOOD: complete body
	return nil
}
