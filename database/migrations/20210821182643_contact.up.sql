create table contact
(
    id uuid not null
        constraint contact_pk
        primary key,
    phone varchar,
    email varchar
);

alter table contact owner to mdgkb;

create unique index contact_id_uindex
    on contact (id);

