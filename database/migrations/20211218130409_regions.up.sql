create table regions
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar not null
);
