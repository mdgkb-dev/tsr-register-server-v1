create table head_circumference (
  id uuid default uuid_generate_v4() not null primary key,
  value integer not null,
  date date not null,
  patient_id uuid not null references patients on delete cascade,
  deleted_at timestamptz default NULL
);
