# SQLite Code Generation

## Purpose

When I write a web application, my number one bottleneck is writing the data layer. This project mostly solves that problem. It automatically generates go code from your SQLite tables. Of course, I could just use an LLM but where is the fun in that!

## Restrictions

This project takes a highly opinionated view of how SQL should be written. Besides the databse being SQLite[^1] there are several notable requirements:

- Every table must be `strict`[^2]
- Every table must have an `id integer primary key` column
- Every table must have `created_at text` and `updated_at text` columns
- Every column must be `not null`

Tables that do not meet these requirements will not have code generated for them.

Given those requirements the minimal working table has this shape:

```sql
create table some_table_name(
    id integer primary key not null,
    created_at text not null,
    updated_at text not null,
) strict;
```

**Note:** has not been tested on Windows or MacOS.

## Installation

```bash
go install github.com/bencornia/sqlitegen/cmd/sqlitegen@latest
```

## Usage

### CLI

Create a database.

```bash
sqlite3 db.sqlite <<EOF
create table employee(
    id integer primary key not null,
    name text not null,
    created_at text not null,
    updated_at text not null
) strict;
EOF
```

Generate code from database

```bash
sqlitegen db.sqlite
```

And you will see the following

```go
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
                select  id,
                        name,
                        created_at,
                        updated_at
                from    employee
                where   id = ?;
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

// code continues...
```

### go generate

The `go generate`[^3] command is a tool for running programs that generate code like `sqlitegen`. At the top of your go file put the following:

```go
//go:generate go run github.com/bencornia/sqlitegen/cmd/sqlitegen@latest -output <path to generated file> <path to database file>
```

When you run `go build`, it will automatically detect the `//go:generate` command and generate your code prior to compilation.

See the `examples` for usage.

## References

[^1]: [SQLite Home Page](https://sqlite.org)

[^2]: [STRICT tables](https://www.sqlite.org/stricttables.html)

[^3]: [The Go Blog: Generating code](https://go.dev/blog/generate)
