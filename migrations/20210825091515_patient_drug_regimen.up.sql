create table patient_drug_regimens
(
    id uuid not null
        constraint patient_drug_regimen_pk
        primary key,
    patient_id uuid
        constraint patient_drug_regimen_patients_id_fk
        references patients
        on delete cascade,
    drug_regimen_id uuid
        constraint patient_drug_regimen_drug_regimen_id_fk
        references drug_regimens
        on delete cascade,
    date date not null
);

create unique index patient_drug_regimen_id_uindex
    on patient_drug_regimens (id);

