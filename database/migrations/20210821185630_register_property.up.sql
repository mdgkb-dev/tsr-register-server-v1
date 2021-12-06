create table register_property
(
    id uuid default uuid_generate_v4() not null
        primary key,
    name varchar not null,
    register_property_order int,
    value_type_id uuid
        references value_type,
    with_other boolean default false,
    register_group_id uuid references register_group on delete cascade,
    tag varchar
);

