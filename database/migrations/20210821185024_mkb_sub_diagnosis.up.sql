create table mkb_sub_diagnosis
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_8fd0b5ba8a7826d8fec3e386c56"
        primary key,
    name varchar,
    sub_code integer,
    comment varchar,
    leaf boolean default true not null,
    relevant boolean default true not null,
    mkb_diagnosis_id uuid
        constraint "FK_6540cfe031daa6ca8369e31fb88"
        references mkb_diagnosis
        on delete cascade
);
