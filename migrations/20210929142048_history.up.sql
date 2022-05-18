create table histories (
  id uuid default uuid_generate_v4() not null constraint history_pkey primary key,
  request_type varchar,
  request_date timestamp default current_timestamp not null
);

create unique index history_id_uindex on histories (id);
