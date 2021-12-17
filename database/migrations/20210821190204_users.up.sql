create table users
(
    id uuid default uuid_generate_v4() not null primary key,
    login varchar not null unique,
    password varchar not null,
    region varchar
);
