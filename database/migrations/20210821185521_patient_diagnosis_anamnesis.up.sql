create table patient_diagnosis_anamnesis
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_21532b5fe975918c36ba3e8b052"
        primary key,
    value varchar not null,
    date date not null,
    patient_diagnosis_id uuid
        constraint "FK_acf9b8a991f6feae36841ccc51d"
        references patient_diagnosis
);

alter table patient_diagnosis_anamnesis owner to mdgkb;

create index "IDX_acf9b8a991f6feae36841ccc51"
    on patient_diagnosis_anamnesis (patient_diagnosis_id);

