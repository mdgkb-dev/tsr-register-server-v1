create table drugs
(
    id uuid not null
        constraint drugs_pk
        primary key,
    name varchar ,
    name_mnn varchar,
    form varchar,
    doze varchar,
    registered boolean,
    date_registration date
);

create unique index drugs_id_uindex
    on drugs (id);

