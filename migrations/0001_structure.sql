create database if not exists kazusa;

create table if not exists courses (
    id binary(16) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    title varchar(256) not null,
    description varchar(2048) not null,
    price integer not null,
    primary key (id)
);
