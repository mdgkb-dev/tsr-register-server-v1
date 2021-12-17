create table registers_users (
  id uuid default uuid_generate_v4() not null primary key,
  register_id uuid references register on delete cascade,
  user_id uuid not null references users on update cascade
);
