create table register_property_others (
  id uuid default uuid_generate_v4() not null primary key,
  name varchar not null,
  register_property_id uuid references register_property on delete cascade,
  register_property_radio_id uuid references register_property_radio on delete cascade,
  register_property_set_id uuid references register_property_set on delete cascade
);
