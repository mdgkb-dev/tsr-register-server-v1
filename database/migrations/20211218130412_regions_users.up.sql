create table regions_users (
  id uuid default uuid_generate_v4() not null primary key,
  region_id uuid references regions on delete cascade,
  user_id uuid not null references users on update cascade
);
