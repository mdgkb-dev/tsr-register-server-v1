create table edv
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_b81e86bf44d2ba7c29143211c80"
        primary key,
    disability_id uuid not null
        constraint "FK_fce3fe296a18bbb9bac5be1b8f9"
        references disability
        on delete cascade,
    parameter1 boolean not null,
    parameter2 boolean not null,
    parameter3 boolean not null,
    period_id uuid,
        file_info_id uuid
        constraint file_info_to_edv_id_fk
        references file_infos (id)
        on update cascade on delete cascade
);

alter table edv owner to mdgkb;

create index "IDX_fce3fe296a18bbb9bac5be1b8f"
    on edv (disability_id);


alter table edv
    add constraint edv_period_id_fk
        foreign key (period_id) references period
            on update cascade on delete cascade;

