create table if not exists user_activity (
    id binary(16) not null,
    created_at timestamp default current_timestamp,
    user_id binary(16) not null,
    course_id binary(16) not null,
    module_id binary(16) not null,
    primary key (id),
    foreign key (user_id) references users (id),
    foreign key (course_id) references courses (id),
    foreign key (module_id) references modules (id)
)
