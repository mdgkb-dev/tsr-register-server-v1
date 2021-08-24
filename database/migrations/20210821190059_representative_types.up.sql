create table representative_types (
  id uuid default uuid_generate_v4() not null constraint "PK_dc3da5abc0483a92389dda30251" primary key,
  parent_male_type varchar,
  parent_woman_type varchar,
  child_male_type varchar,
  child_woman_type varchar
);

alter table
  representative_types owner to mdgkb;
