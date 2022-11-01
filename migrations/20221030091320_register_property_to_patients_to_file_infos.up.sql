create table register_properties_to_patients_to_file_infos
(
    id uuid default uuid_generate_v4() not null primary key,
      register_property_to_patient_id uuid references register_property_to_patient(id) on delete cascade,
      file_info_id uuid references file_infos(id) on delete cascade
);


alter table register_property
    add is_files_storage bool;
