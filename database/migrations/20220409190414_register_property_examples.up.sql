create table register_property_examples
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar not null,
    register_property_radio_order int,
    register_property_id uuid references register_property on delete cascade,
        register_property_example_order int
);