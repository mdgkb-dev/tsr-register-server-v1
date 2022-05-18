create table drugs_diagnosis
(
    id uuid default uuid_generate_v4() not null
            primary key,
    mkb_diagnosis_id uuid
            references mkb_diagnosis,
    mkb_sub_diagnosis_id uuid
            references mkb_sub_diagnosis,
    drug_id uuid
            references drugs
            on delete cascade
);
