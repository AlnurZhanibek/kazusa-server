alter table course_payments
    add column order_id binary(16),
    add column confirmed bool default 0;