create or replace view humans_view(id, name, surname, patronymic, full_name, is_male, date_birth, address_registration, address_residential, contact_id, photo_id) as
SELECT human.id,
       human.name,
       human.surname,
       human.patronymic,
       CONCAT_WS(' '::TEXT, human.surname, human.name, human.patronymic) AS full_name,
       human.is_male,
       human.date_birth,
       human.address_registration,
       human.address_residential,
       human.contact_id,
       human.photo_id
FROM human;
