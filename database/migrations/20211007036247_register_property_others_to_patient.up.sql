create table register_property_other_to_patient (
  id uuid default uuid_generate_v4() not null primary key,
  value varchar not null,
  register_property_other_id uuid references register_property_others on delete cascade,
      patient_id uuid not null references patient on update cascade
);
