-- create table foo(id integer primary key not null, updated_at text not null, created_at text not null) strict;
create table foo(
	id integer primary key not null,
	updated_at text not null,
	created_at text not null
) strict;
with flags(
	is_strict_table,
	has_not_null_columns,
	has_valid_pk,
	has_valid_created_at,
	has_valid_updated_at
) as (
	select
		(
			select strict
			from pragma_table_list('foo')
		) as is_strict_table,
		(
			select sum("notnull" = 0) = 0
			from pragma_table_info('foo')
		) as has_not_null_columns,
		(
			select count(*) = 1
			from pragma_table_info('foo')
			where name = 'id'
				and type = 'INTEGER'
				and pk = 1
		) as has_valid_pk,
		(
			select count(*) = 1
			from pragma_table_info('foo')
			where name = 'created_at'
				and type = 'TEXT'						
		) as has_valid_created_at,
		(
			select count(*) = 1
			from pragma_table_info('foo')
			where name = 'updated_at'
				and type = 'TEXT'
		) as has_valid_updated_at
)
select is_strict_table
	and has_not_null_columns
	and has_valid_pk
	and has_valid_created_at
	and has_valid_updated_at
from flags;

