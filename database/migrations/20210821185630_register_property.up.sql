create table register_property
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_4f116464d566295172a583cb8e9"
        primary key,
    name varchar not null,
    value_type_id uuid
        constraint "FK_f021d1969d6682de1f265c62a6c"
        references value_type,
    with_other boolean default false
);

alter table register_property owner to mdgkb;

create index "IDX_f021d1969d6682de1f265c62a6"
    on register_property (value_type_id);

