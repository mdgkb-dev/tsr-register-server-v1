create table register_property_variants
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    register_property_id uuid references register_property(id) on delete cascade
);

alter table register_property_to_patient
    add column register_property_variant_id uuid REFERENCES register_property_variants(id) ON UPDATE CASCADE ON DELETE CASCADE;

alter table register_property
    add age_compare boolean;



alter table register_property_to_patient
    alter column value_number type numeric using value_number::numeric;

