create table patient_diagnosis
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_14ad82a727c31d1f6e14d16749f"
        primary key,
    "primary" boolean,
    mkb_diagnosis_id uuid
        constraint "FK_77c97b0af19a1c580b580c8c76c"
        references mkb_diagnosis,
    mkb_sub_diagnosis_id uuid
        constraint "FK_aa2052f2cea3b0c3a4825fc0a1b"
        references mkb_sub_diagnosis,
    patient_id uuid
        constraint "FK_2d6181c875a32b2ec25356689e5"
        references patient
        on delete cascade
);

create index "IDX_77c97b0af19a1c580b580c8c76"
    on patient_diagnosis (mkb_diagnosis_id);

create index "IDX_aa2052f2cea3b0c3a4825fc0a1"
    on patient_diagnosis (mkb_sub_diagnosis_id);

create index "IDX_2d6181c875a32b2ec25356689e"
    on patient_diagnosis (patient_id);

