# SQLite Code Generation

## Purpose

- generate a `models` package in go from a sqlite file

## Restrictions

- Each table must have an id, updated_at, and created_at fields
- Otherwise, they will be ignored

## Usage

```bash
cd examples/basic
go generate
go run main.go
```
