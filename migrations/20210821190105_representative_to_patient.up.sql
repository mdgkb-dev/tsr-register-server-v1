create table representative_to_patient
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_a0a5ee0b97383efda9364728589"
        primary key,
    representative_type_id uuid not null
        constraint "FK_3e3f0b4a9e1a050b6fb525158ec"
        references representative_types
        on update cascade on delete cascade,
    patient_id uuid not null
        constraint "FK_e6a11936ba24e187c82910aea36"
        references patients
        on update cascade on delete cascade,
    representative_id uuid not null
        constraint "FK_8ae4fe72ff71cdee5b07ff9f4ff"
        references representative
        on update cascade on delete cascade
);

create index "IDX_3e3f0b4a9e1a050b6fb525158e"
    on representative_to_patient (representative_type_id);

create index "IDX_e6a11936ba24e187c82910aea3"
    on representative_to_patient (patient_id);

create index "IDX_8ae4fe72ff71cdee5b07ff9f4f"
    on representative_to_patient (representative_id);
