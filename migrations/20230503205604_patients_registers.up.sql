alter table register_to_patient rename to patients_registers;
alter table patients_registers add column user_id uuid references users;

