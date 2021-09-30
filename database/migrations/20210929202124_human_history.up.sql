create table human_histories (
  human_history_id uuid default uuid_generate_v4() not null constraint human_history_pkey primary key,
  history_id uuid constraint human_history_history_id_fk references histories on delete cascade,
  id uuid,
  name varchar,
  surname varchar,
  patronymic varchar,
  is_male boolean,
  date_birth date,
  address_registration varchar,
  address_residential varchar,
  contact_id uuid constraint human_contact_id_fk references contact on update cascade on delete cascade,
  photo_id uuid constraint human_files_id_fk references file_infos on update cascade on delete cascade
);

create unique index human_history_id_uindex on human_histories (id);
