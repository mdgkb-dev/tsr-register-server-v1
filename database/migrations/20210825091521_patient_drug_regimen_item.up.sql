create table patient_drug_regimen_items
(
    id uuid not null
        constraint patient_drug_regimen_item_pk
        primary key,
    patient_drug_regimen_id uuid
        constraint patient_drug_regimen_item_patient_drug_regimen_id_fk
        references patient_drug_regimens,
    date date,
    getting_date date
);

create unique index patient_drug_regimen_item_id_uindex
    on patient_drug_regimen_items (id);

