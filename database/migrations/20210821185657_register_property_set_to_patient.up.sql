create table register_property_set_to_patient
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_65167a07c63bec490ef37417f11"
        primary key,
    register_property_set_id uuid not null
        constraint "FK_8281318758557dfc2a1fd67f090"
        references register_property_set,
    patient_id uuid not null
        constraint "FK_312cd0a936e61b0870b9bd8428a"
        references patient
        on delete cascade
);

create index "IDX_8281318758557dfc2a1fd67f09"
    on register_property_set_to_patient (register_property_set_id);

create index "IDX_312cd0a936e61b0870b9bd8428"
    on register_property_set_to_patient (patient_id);
