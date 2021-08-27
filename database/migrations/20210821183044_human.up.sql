create table human
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_fa2d597665c4d7604049d5f7792"
        primary key,
    name varchar,
    surname varchar,
    patronymic varchar,
    is_male boolean,
    date_birth date,
    address_registration varchar,
    address_residential varchar,
    contact_id uuid
        constraint human_contact_id_fk
        references contact
        on update cascade on delete cascade,
    photo_id uuid
        constraint human_files_id_fk
        references file_infos
        on update cascade on delete cascade
);

create index "IDX_258ba83caba75fd1eaeab015c5"
    on human (name);

create index "IDX_00bc4bd4741185902403939f01"
    on human (surname);

