alter table patients
    add created_at timestamp default current_timestamp not null;

alter table patients
    add created_by_id uuid;

alter table patients
    add updated_at timestamp default current_timestamp;

alter table patients
    add updated_by_id uuid;
