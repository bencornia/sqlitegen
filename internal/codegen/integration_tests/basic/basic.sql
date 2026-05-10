drop table if exists basic;
create table basic(
    id integer primary key not null,
    created_at text not null,
    updated_at text not null
) strict;
