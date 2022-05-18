create table register_to_patient
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_8dcc4e1fb603e488f7384f1f26a"
        primary key,
    register_id uuid not null
        constraint "FK_869a2523a06d431f82897e1ce2f"
        references register,
    patient_id uuid not null
        constraint "FK_b192a657a48ce5e370b8bedd93c"
        references patients
);

create index "IDX_869a2523a06d431f82897e1ce2"
    on register_to_patient (register_id);

create index "IDX_b192a657a48ce5e370b8bedd93"
    on register_to_patient (patient_id);
