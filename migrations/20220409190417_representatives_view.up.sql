create or replace view representatives_view as
SELECT
    r.*,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM representative r
         JOIN human h on r.human_id = h.id;
