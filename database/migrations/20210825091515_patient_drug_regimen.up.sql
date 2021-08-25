create table patient_drug_regimens
(
    id uuid not null
        constraint patient_drug_regimen_pk
        primary key,
    patient_id uuid
        constraint patient_drug_regimen_patients_id_fk
        references patient,
    drug_regimen_id uuid
        constraint patient_drug_regimen_drug_regimen_id_fk
        references drug_regimens
);

create unique index patient_drug_regimen_id_uindex
    on patient_drug_regimens (id);

