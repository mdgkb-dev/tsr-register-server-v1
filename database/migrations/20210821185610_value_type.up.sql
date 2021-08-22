create table value_type
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_38e761371cca3cc3051fb60674a"
        primary key,
    name varchar,
    value_relation value_type_value_relation_enum not null
);

alter table value_type owner to mdgkb;

