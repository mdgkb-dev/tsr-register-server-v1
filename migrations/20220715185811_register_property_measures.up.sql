create table register_property_measures
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    register_property_measure_order int
      register_property_id uuid references register_property() on delete cascade
);

alter table register_property_to_patient
    add column register_property_measure_id uuid REFERENCES register_property_measures(id) ON UPDATE CASCADE ON DELETE CASCADE;


