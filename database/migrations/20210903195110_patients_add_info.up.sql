alter table patient
    add created_at timestamp default current_timestamp not null;

alter table patient
    add created_by_id uuid;

alter table patient
    add updated_at timestamp default current_timestamp;

alter table patient
    add updated_by_id uuid;
