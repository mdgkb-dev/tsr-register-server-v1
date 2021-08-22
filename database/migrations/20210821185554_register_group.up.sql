create table register_group
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_e75c0bdfe758028ce3d014966bc"
        primary key,
    name varchar not null
);

alter table register_group owner to mdgkb;

