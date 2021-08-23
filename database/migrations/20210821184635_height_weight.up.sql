create table height_weight (
  id uuid default uuid_generate_v4() not null primary key,
  height integer not null,
  weight integer not null,
  date date not null,
  patient_id uuid not null references patient on delete cascade
);
