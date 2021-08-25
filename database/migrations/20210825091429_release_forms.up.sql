create table release_forms
(
    id uuid not null
        constraint release_forms_pk
        primary key,
    name varchar
);

create unique index release_forms_id_uindex
    on release_forms (id);

