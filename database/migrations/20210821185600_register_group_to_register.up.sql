create table register_group_to_register
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_dc8ce69c9367b63ba567dff4c9d"
        primary key,
    register_id uuid not null
        constraint "FK_c948733e47b257d8e14fa1145c3"
        references register
        on update cascade on delete cascade,
    "order" integer,
    register_group_id uuid not null
        constraint "FK_bda72dac4ea175db2a9f4e3c93a"
        references register_group
        on update cascade on delete cascade
);

alter table register_group_to_register owner to mdgkb;

create index "IDX_c948733e47b257d8e14fa1145c"
    on register_group_to_register (register_id);

create index "IDX_bda72dac4ea175db2a9f4e3c93"
    on register_group_to_register (register_group_id);

