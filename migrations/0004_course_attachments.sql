alter table courses
    add column
        (cover_url varchar(256) default 'https://hyderabadangels.in/wp-content/uploads/2019/11/dummy-logo.png',
        attachment_urls json)
;