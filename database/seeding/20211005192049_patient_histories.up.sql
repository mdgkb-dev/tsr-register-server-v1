INSERT INTO human_histories (id, name, surname, patronymic, is_male, date_birth, address_registration, address_residential, contact_id, photo_id)
SELECT h.* FROM patient p JOIN human h ON h.id = p.human_id;

alter table patient_histories drop constraint patient_history_history_id_fk;

INSERT INTO patient_histories (history_id, human_history_id, id, human_id, created_at, created_by_id, updated_at, updated_by_id, deleted_at)
SELECT uuid_generate_v4() , hh.human_history_id, p.*
FROM patient p join human_histories hh on hh.id = p.human_id;

INSERT INTO histories (id, request_type, request_date)
SELECT patient_histories.history_id, 'Создание', '2021-10-01' FROM patient_histories;

alter table patient_histories
    add constraint patient_histories_histories_id_fk
        foreign key (history_id) references histories;


