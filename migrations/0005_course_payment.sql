create table if not exists course_payments (
    id binary(16) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    user_id binary(16) not null,
    course_id binary(16) not null,
    primary key (id),
    foreign key (user_id) references users (id),
    foreign key (course_id) references courses (id)
)
