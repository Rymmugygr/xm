create table if not exists companies
(
    id          uuid default uuidv7() not null
        constraint companies_pk
            primary key,
    name        text                  not null
        constraint companies_name_uniq
            unique,
    description text                  not null,
    employees   integer               not null,
    registered  boolean               not null,
    type        text                  not null,
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);
