create table register_groups_to_patients
(
    id uuid default uuid_generate_v4() not null primary key,
    register_groups_to_patients_date date,
    patient_id uuid  REFERENCES patients(id) ON UPDATE CASCADE ON DELETE CASCADE,
    register_group_id uuid  REFERENCES register_group(id) ON UPDATE CASCADE ON DELETE CASCADE
);

drop index "IDX_607312bf055d7e060c75b151b6";


alter table register_property_to_patient
    drop column patient_id;

alter table register_property_to_patient
    add column register_group_to_patient_id uuid REFERENCES register_groups_to_patients(id) ON UPDATE CASCADE ON DELETE CASCADE;


--alter table register_property
--    drop column with_dates;

alter table register_group
    add column with_dates boolean;


alter table register_property_set_to_patient
    drop column patient_id;

alter table register_property_set_to_patient
    add column register_group_to_patient_id uuid REFERENCES register_groups_to_patients(id) ON UPDATE CASCADE ON DELETE CASCADE;


alter table register_property_other_to_patient
    drop column patient_id;

alter table register_property_other_to_patient
    add column register_group_to_patient_id uuid REFERENCES register_groups_to_patients(id) ON UPDATE CASCADE ON DELETE CASCADE;

