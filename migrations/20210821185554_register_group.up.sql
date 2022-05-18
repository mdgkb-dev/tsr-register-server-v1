create table register_group
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar not null,
    register_id uuid references register on delete cascade,
    register_group_order int
);
