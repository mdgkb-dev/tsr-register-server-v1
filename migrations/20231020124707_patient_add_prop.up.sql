ALTER table patients add column is_moscow bool default false;
ALTER table patients add column agreed bool default false;




INSERT into questions (id, name, item_order, value_type_id, research_id) 
values ('3df34118-cf2f-4a09-bca0-0428d93e0401', 'Описание реакций родителями', 4, '50dfc0a4-b260-4c63-b16f-4119e152037f', '8f2f58fd-38e2-4644-b8e7-e05794e838a8');


drop  view  patients_view;
create or replace view patients_view as
SELECT
    p.*,
    h.is_male,
    h.date_birth,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name
FROM patients p
         JOIN humans h on p.human_id = h.id
         join users creator on p.created_by_id = creator.id;
