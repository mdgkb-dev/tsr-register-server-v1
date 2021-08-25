create table drugs
(
    id uuid not null
        constraint drugs_pk
        primary key,
    name varchar not null
);

create unique index drugs_id_uindex
    on drugs (id);

