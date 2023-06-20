create table drug_application_files
(
    id   uuid default uuid_generate_v4() not null primary key,
    drug_application_id uuid REFERENCES drug_applications(id) ON UPDATE CASCADE ON DELETE CASCADE,
    file_info_id uuid REFERENCES file_infos(id) ON UPDATE CASCADE ON DELETE CASCADE,
    comment varchar,
    name varchar
);


alter table fund_contracts add column
drug_application_id uuid REFERENCES drug_applications(id) ON UPDATE CASCADE ON DELETE CASCADE;

alter table drug_arrives drop column  drug_id;
alter table fund_contracts drop constraint fund_contracts_number_key;

alter table drug_decreases drop constraint drug_decreases_patient_id_fkey;
alter table drug_decreases alter column patient_id drop not null;

