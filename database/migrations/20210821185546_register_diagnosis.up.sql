create table register_diagnosis
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_e4146e0330822ee6bc676e81fab"
        primary key,
    mkb_diagnosis_id uuid
        constraint "FK_a3649e77a04511ca712018cf65f"
        references mkb_diagnosis,
    mkb_sub_diagnosis_id uuid
        constraint "FK_a83764b0590c2525c00f894ec29"
        references mkb_sub_diagnosis,
    register_id uuid
        constraint "FK_960bb84f4428048129d2406096e"
        references register
        on delete cascade
);

alter table register_diagnosis owner to mdgkb;

create index "IDX_a3649e77a04511ca712018cf65"
    on register_diagnosis (mkb_diagnosis_id);

create index "IDX_a83764b0590c2525c00f894ec2"
    on register_diagnosis (mkb_sub_diagnosis_id);

create index "IDX_960bb84f4428048129d2406096"
    on register_diagnosis (register_id);

