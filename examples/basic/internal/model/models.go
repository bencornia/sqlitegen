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

func (s *CompanyStore) Insert(ctx context.Context, item *Company) (int64, error) {
	query := `
		insert into company(
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

func (s *CompanyStore) DeleteMany(ctx context.Context, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from company
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStore) GetByIdTx(ctx context.Context, tx *sql.Tx, id int64) (*Company, error) {
	query := `
		select	id,
				name,
				created_at,
				updated_at
		from	company
		where	id = ?;
	`

	var item Company
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

func (s *CompanyStore) UpdateByIdTx(ctx context.Context, tx *sql.Tx, item *Company) error {
	query := `
		update	company
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

func (s *CompanyStore) InsertTx(ctx context.Context, tx *sql.Tx, item *Company) (int64, error) {
	query := `
		insert into company(
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

func (s *CompanyStore) DeleteByIdTx(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `
		delete from company
		where id = ?;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStore) GetManyTx(ctx context.Context, tx *sql.Tx, ids []int64) ([]*Company, error) {
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
	rows, err := tx.QueryContext(ctx, query, args...)
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

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *CompanyStore) DeleteManyTx(ctx context.Context, tx *sql.Tx, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from company
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStore) UpdateManyTx(ctx context.Context, tx *sql.Tx, ids []int64, item *Company) ([]int64, error) {
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("%d", id)
	}

	query := `
		update	company
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

func (s *CompanyStore) InsertMany(ctx context.Context, tx *sql.Tx, items []*Company) ([]int64, error) {
	query := `
		insert into company (
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
		set		company_id = ?,
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
		values (
			?,
			?,
			datetime(),
			datetime()
		);
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

func (s *DepartmentStore) DeleteMany(ctx context.Context, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from department
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *DepartmentStore) GetByIdTx(ctx context.Context, tx *sql.Tx, id int64) (*Department, error) {
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
	err := tx.QueryRowContext(ctx, query, id).Scan(
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

func (s *DepartmentStore) UpdateByIdTx(ctx context.Context, tx *sql.Tx, item *Department) error {
	query := `
		update	department
		set		company_id = ?,
				name = ?,
				updated_at = datetime()
		where	id = ?;
	`

	_, err := tx.ExecContext(
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

func (s *DepartmentStore) InsertTx(ctx context.Context, tx *sql.Tx, item *Department) (int64, error) {
	query := `
		insert into department(
			company_id,
			name
		)
		values (
			?,
			?,
			datetime(),
			datetime()
		);
	`

	result, err := tx.ExecContext(
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

func (s *DepartmentStore) DeleteByIdTx(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `
		delete from department
		where id = ?;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DepartmentStore) GetManyTx(ctx context.Context, tx *sql.Tx, ids []int64) ([]*Department, error) {
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
	rows, err := tx.QueryContext(ctx, query, args...)
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

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *DepartmentStore) DeleteManyTx(ctx context.Context, tx *sql.Tx, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from department
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *DepartmentStore) UpdateManyTx(ctx context.Context, tx *sql.Tx, ids []int64, item *Department) ([]int64, error) {
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("%d", id)
	}

	query := `
		update	department
		set		company_id = ?,
				name = ?,
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

func (s *DepartmentStore) InsertMany(ctx context.Context, tx *sql.Tx, items []*Department) ([]int64, error) {
	query := `
		insert into department (
			company_id,
			name,
			created_at,
			updated_at
		) values (			
			?,
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
			&item.CompanyId,
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
		set		department_id = ?,
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
		values (
			?,
			?,
			?,
			datetime(),
			datetime()
		);
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
				department_id,
				name,
				email,
				created_at,
				updated_at
		from	employee
		where	id = ?;
	`

	var item Employee
	err := tx.QueryRowContext(ctx, query, id).Scan(
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

func (s *EmployeeStore) UpdateByIdTx(ctx context.Context, tx *sql.Tx, item *Employee) error {
	query := `
		update	employee
		set		department_id = ?,
				name = ?,
				email = ?,
				updated_at = datetime()
		where	id = ?;
	`

	_, err := tx.ExecContext(
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

func (s *EmployeeStore) InsertTx(ctx context.Context, tx *sql.Tx, item *Employee) (int64, error) {
	query := `
		insert into employee(
			department_id,
			name,
			email
		)
		values (
			?,
			?,
			?,
			datetime(),
			datetime()
		);
	`

	result, err := tx.ExecContext(
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
	rows, err := tx.QueryContext(ctx, query, args...)
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
		set		department_id = ?,
				name = ?,
				email = ?,
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
			department_id,
			name,
			email,
			created_at,
			updated_at
		) values (			
			?,
			?,
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
			&item.DepartmentId,
			&item.Name,
			&item.Email,
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
		set		employee_id = ?,
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
		values (
			?,
			?,
			?,
			datetime(),
			datetime()
		);
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

func (s *EmployeeSalaryStore) DeleteMany(ctx context.Context, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from employee_salary
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSalaryStore) GetByIdTx(ctx context.Context, tx *sql.Tx, id int64) (*EmployeeSalary, error) {
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
	err := tx.QueryRowContext(ctx, query, id).Scan(
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

func (s *EmployeeSalaryStore) UpdateByIdTx(ctx context.Context, tx *sql.Tx, item *EmployeeSalary) error {
	query := `
		update	employee_salary
		set		employee_id = ?,
				amount = ?,
				currency = ?,
				updated_at = datetime()
		where	id = ?;
	`

	_, err := tx.ExecContext(
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

func (s *EmployeeSalaryStore) InsertTx(ctx context.Context, tx *sql.Tx, item *EmployeeSalary) (int64, error) {
	query := `
		insert into employee_salary(
			employee_id,
			amount,
			currency
		)
		values (
			?,
			?,
			?,
			datetime(),
			datetime()
		);
	`

	result, err := tx.ExecContext(
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

func (s *EmployeeSalaryStore) DeleteByIdTx(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `
		delete from employee_salary
		where id = ?;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSalaryStore) GetManyTx(ctx context.Context, tx *sql.Tx, ids []int64) ([]*EmployeeSalary, error) {
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
	rows, err := tx.QueryContext(ctx, query, args...)
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

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *EmployeeSalaryStore) DeleteManyTx(ctx context.Context, tx *sql.Tx, ids []int64) error {
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		delete from employee_salary
		where id in (%s);
	`

	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSalaryStore) UpdateManyTx(ctx context.Context, tx *sql.Tx, ids []int64, item *EmployeeSalary) ([]int64, error) {
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("%d", id)
	}

	query := `
		update	employee_salary
		set		employee_id = ?,
				amount = ?,
				currency = ?,
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

func (s *EmployeeSalaryStore) InsertMany(ctx context.Context, tx *sql.Tx, items []*EmployeeSalary) ([]int64, error) {
	query := `
		insert into employee_salary (
			employee_id,
			amount,
			currency,
			created_at,
			updated_at
		) values (			
			?,
			?,
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
			&item.EmployeeId,
			&item.Amount,
			&item.Currency,
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
