create table file_infos
(
    id uuid default uuid_generate_v4() not null
        constraint file_infos_pkey
        primary key,
    original_name varchar,
    file_system_path varchar
);
