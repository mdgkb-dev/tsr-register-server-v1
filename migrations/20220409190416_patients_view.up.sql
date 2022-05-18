create or replace view patients_view as
SELECT
    p.*,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM patients p
         JOIN human h on p.human_id = h.id;
