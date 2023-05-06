drop table patient_histories;
create table patient_histories
(
    id          uuid      default uuid_generate_v4() not null
        constraint patient_histories_pk
            primary key,
    created_at  timestamp default CURRENT_TIMESTAMP  not null,
    patient_id  uuid,
    user_id     uuid,
    object_copy json,
    action_type varchar
);