// DO NOT EDIT! THIS IS GENERATED CODE!
package model

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

type Company struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *Store) GetCompany() ([]*Company, error) {
	var items []*Company
	query := `
		select id, name, created_at, updated_at
		from company;
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
		var item Company

		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return items, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (s *Store) GetCompanyById(id int64) (*Company, error) {
	var item Company
	query := `
		select id, name, created_at, updated_at
		from company
		where id = ?;
	`

	err := s.db.QueryRow(query, id).Scan(
		&item.Id,
		&item.Name,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return &item, err
	}

	return &item, nil
}

func (s *Store) InsertCompany(item *Company) (*Company, error) {
	query := `
		insert into company(name, created_at, updated_at)
		values (?, datetime(), datetime());
	`

	result, err := s.db.Exec(
		query,
		&item.Name,
	)

	if err != nil {
		return item, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}

	return s.GetCompanyById(id)
}

func (s *Store) UpdateCompany(item *Company) (*Company, error) {
	query := `
		update company
		set name = ?, updated_at = datetime()
	`

	_, err := s.db.Exec(
		query,
		&item.Name,
	)

	if err != nil {
		return item, err
	}

	return s.GetCompanyById(item.Id)
}

func (s *Store) DeleteCompany(id int) error {
	query := `
		delete from company
		where id = ?
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

type Department struct {
	Id        int64  `json:"id"`
	CompanyId int64  `json:"company_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *Store) GetDepartment() ([]*Department, error) {
	var items []*Department
	query := `
		select id, company_id, name, created_at, updated_at
		from department;
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
		var item Department

		err = rows.Scan(
			&item.Id,
			&item.CompanyId,
			&item.Name,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return items, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (s *Store) GetDepartmentById(id int64) (*Department, error) {
	var item Department
	query := `
		select id, company_id, name, created_at, updated_at
		from department
		where id = ?;
	`

	err := s.db.QueryRow(query, id).Scan(
		&item.Id,
		&item.CompanyId,
		&item.Name,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return &item, err
	}

	return &item, nil
}

func (s *Store) InsertDepartment(item *Department) (*Department, error) {
	query := `
		insert into department(company_id, name, created_at, updated_at)
		values (?, ?, datetime(), datetime());
	`

	result, err := s.db.Exec(
		query,
		&item.CompanyId,
		&item.Name,
	)

	if err != nil {
		return item, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}

	return s.GetDepartmentById(id)
}

func (s *Store) UpdateDepartment(item *Department) (*Department, error) {
	query := `
		update department
		set company_id = ?, name = ?, updated_at = datetime()
	`

	_, err := s.db.Exec(
		query,
		&item.CompanyId,
		&item.Name,
	)

	if err != nil {
		return item, err
	}

	return s.GetDepartmentById(item.Id)
}

func (s *Store) DeleteDepartment(id int) error {
	query := `
		delete from department
		where id = ?
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

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

func (s *Store) GetEmployee() ([]*Employee, error) {
	var items []*Employee
	query := `
		select id, department_id, name, email, created_at, updated_at
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
			&item.DepartmentId,
			&item.Name,
			&item.Email,
			&item.CreatedAt,
			&item.UpdatedAt,
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
		select id, department_id, name, email, created_at, updated_at
		from employee
		where id = ?;
	`

	err := s.db.QueryRow(query, id).Scan(
		&item.Id,
		&item.DepartmentId,
		&item.Name,
		&item.Email,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return &item, err
	}

	return &item, nil
}

func (s *Store) InsertEmployee(item *Employee) (*Employee, error) {
	query := `
		insert into employee(department_id, name, email, created_at, updated_at)
		values (?, ?, ?, datetime(), datetime());
	`

	result, err := s.db.Exec(
		query,
		&item.DepartmentId,
		&item.Name,
		&item.Email,
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
		set department_id = ?, name = ?, email = ?, updated_at = datetime()
	`

	_, err := s.db.Exec(
		query,
		&item.DepartmentId,
		&item.Name,
		&item.Email,
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

type EmployeeSalary struct {
	Id         int64   `json:"id"`
	EmployeeId int64   `json:"employee_id"`
	Amount     float64 `json:"amount"`
	Currency   *string `json:"currency"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

func (s *Store) GetEmployeeSalary() ([]*EmployeeSalary, error) {
	var items []*EmployeeSalary
	query := `
		select id, employee_id, amount, currency, created_at, updated_at
		from employee_salary;
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
			return items, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (s *Store) GetEmployeeSalaryById(id int64) (*EmployeeSalary, error) {
	var item EmployeeSalary
	query := `
		select id, employee_id, amount, currency, created_at, updated_at
		from employee_salary
		where id = ?;
	`

	err := s.db.QueryRow(query, id).Scan(
		&item.Id,
		&item.EmployeeId,
		&item.Amount,
		&item.Currency,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return &item, err
	}

	return &item, nil
}

func (s *Store) InsertEmployeeSalary(item *EmployeeSalary) (*EmployeeSalary, error) {
	query := `
		insert into employee_salary(employee_id, amount, currency, created_at, updated_at)
		values (?, ?, ?, datetime(), datetime());
	`

	result, err := s.db.Exec(
		query,
		&item.EmployeeId,
		&item.Amount,
		&item.Currency,
	)

	if err != nil {
		return item, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}

	return s.GetEmployeeSalaryById(id)
}

func (s *Store) UpdateEmployeeSalary(item *EmployeeSalary) (*EmployeeSalary, error) {
	query := `
		update employee_salary
		set employee_id = ?, amount = ?, currency = ?, updated_at = datetime()
	`

	_, err := s.db.Exec(
		query,
		&item.EmployeeId,
		&item.Amount,
		&item.Currency,
	)

	if err != nil {
		return item, err
	}

	return s.GetEmployeeSalaryById(item.Id)
}

func (s *Store) DeleteEmployeeSalary(id int) error {
	query := `
		delete from employee_salary
		where id = ?
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
