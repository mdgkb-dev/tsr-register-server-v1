create table register_property_to_patient
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_762042febd449b4a25722e68f4d"
        primary key,
    value_string varchar,
    value_other varchar,
    value_number integer,
    value_date date,
    register_property_radio_id uuid
        constraint "FK_7e029f3d1c51de39d4a6fa79825"
        references register_property_radio,
    register_property_id uuid not null
        constraint "FK_94ad51b950a805160948c8540ef"
        references register_property,
    patient_id uuid not null
        constraint "FK_607312bf055d7e060c75b151b6f"
        references patient
        on update cascade
);

create index "IDX_7e029f3d1c51de39d4a6fa7982"
    on register_property_to_patient (register_property_radio_id);

create index "IDX_94ad51b950a805160948c8540e"
    on register_property_to_patient (register_property_id);

create index "IDX_607312bf055d7e060c75b151b6"
    on register_property_to_patient (patient_id);

