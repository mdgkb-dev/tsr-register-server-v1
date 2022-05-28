create table mkb_concrete_diagnosis
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    comment varchar,
    leaf boolean default true not null,
    relevant boolean default true not null,
    mkb_sub_diagnosis_id uuid
            references mkb_sub_diagnosis
            on delete cascade
);
