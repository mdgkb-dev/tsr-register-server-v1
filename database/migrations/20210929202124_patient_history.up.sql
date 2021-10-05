create table patient_histories (
  patient_history_id uuid default uuid_generate_v4() not null constraint patient_history_pkey primary key,
  history_id uuid constraint patient_history_history_id_fk references histories,
  human_history_id uuid constraint patient_history_human_history_id_fk references human_histories,
  id uuid,
  human_id uuid,
  created_at timestamp default current_timestamp not null,
  created_by_id uuid,
  updated_at timestamp default current_timestamp,
  updated_by_id uuid
);

create unique index patient_history_patient_history_id_uindex on patient_histories (patient_history_id);
