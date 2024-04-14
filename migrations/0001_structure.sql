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

create table if not exists modules (
    id binary(16) not null,
    course_id binary(16) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    name varchar(256) not null,
    content varchar(8192) not null,
    duration_minutes integer not null,
    foreign key (course_id) references courses (id)
);
