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

## Installation

```bash
go install github.com/bencornia/sqlitegen/cmd/sqlitegen@latest
```

## References

- [^1]: [SQLite Home Page](https://sqlite.org)
- [^2]: [STRICT tables](https://www.sqlite.org/stricttables.html)
