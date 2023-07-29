create table domains
(
    id          uuid      default uuid_generate_v4() not null primary key ,
    name              varchar
);

alter table questions
    add column domain_id uuid REFERENCES domains(id) ON UPDATE CASCADE ON DELETE CASCADE;

alter table answers
    add column patient_id uuid REFERENCES patients(id) ON UPDATE CASCADE ON DELETE CASCADE;


alter table users
    add column domain_id uuid REFERENCES domains(id) ON UPDATE CASCADE ON DELETE CASCADE;



insert into public.domains (id, name)
values  ('8669a507-5da0-4603-99b6-3e79e41f3f35', 'sma'),
        ('b9d7b8a5-d155-4dd5-8040-83c2648f0949', 'anaphylaxis');


UPDATE public.users SET domain_id = 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' WHERE id = '1c05d56c-5b98-4361-b42c-ebd5107cfb5e';
UPDATE public.users SET domain_id = '8669a507-5da0-4603-99b6-3e79e41f3f35' WHERE id = '5d04ac28-70d0-4926-adca-5da55a6cf7e4';
UPDATE public.users SET domain_id = 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' WHERE id = '59f39975-83ab-434d-954c-8b6fd1ffa4a2';
UPDATE public.users SET domain_id = 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' WHERE id = '86f5fb02-f65c-431b-a7fd-849369d9c21b';
UPDATE public.users SET domain_id = 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' WHERE id = '342221bf-810d-4909-b68a-da939783fba1';
UPDATE public.users SET domain_id = '8669a507-5da0-4603-99b6-3e79e41f3f35' WHERE id = '5d04ac28-70d0-4926-adca-5da55a6cf7e3';
UPDATE public.users SET domain_id = 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' WHERE id = 'bc3bf514-834c-416b-a2e6-a349f4b4ecce';

drop  view  patients_view;
create or replace view patients_view as
SELECT
    p.*,
    h.is_male,
    h.date_birth,
    CONCAT_WS(' '::TEXT, h.surname, h.name, h.patronymic) AS full_name,
    creator.domain_id as domain_id
FROM patients p
         JOIN humans h on p.human_id = h.id
         join users creator on p.created_by_id = creator.id;


insert into questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ('b54b5ce1-998b-4363-8ddb-e5dddba237bb', 'госпитализация', 0, 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', false, false, null, null, null, null, null, null, null, null, null, null, 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');


insert into public.answer_variants (id, name, item_order, question_id, score, show_more_questions)
values  ('022fc517-81af-4c53-9207-89e7e3eb5aae', 'Экстренная', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bb', 0, false),
        ('b508f7ac-54b7-411d-bf3c-9594f8e0bf2c', 'Плановая', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bb', 1, false);


insert into questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ('b54b5ce1-998b-4363-8ddb-e5dddba237bc', 'диагноз T78.0', 0, 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', false, false, null, null, null, null, null, null, null, null, null, null, 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');


insert into public.answer_variants (id, name, item_order, question_id, score, show_more_questions)
values  ('022fc517-81af-4c53-9207-89e7e3eb5aad', 'Подтверждён', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bc', 0, false),
        ('b508f7ac-54b7-411d-bf3c-9594f8e0bf2d', 'Не подтверждён', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bc', 1, false);


update patients p1
set created_by_id = '342221bf-810d-4909-b68a-da939783fba1'
from patients p2
         join patients_researches_pools pr on p2.id = pr.patient_id
where pr.researches_pool_id = 'a0680b34-0d9b-4df2-9288-fbc421fd3ee5'
  and p1.id = p2.id;

insert into public.users (id, login, password, region, email, uuid, domain_id)
values  ('7ef578b3-86c5-421f-8f6a-2a4de2cbf6dd', '', '$2a$10$r3AlsAUhJZASy028tu6GcuitOp6vkL9xKszP1FV/RQ2ZdNagKHKW6', null, 'anaphylaxis@gmail.com', 'b7263799-152e-476a-b79a-2406e1a2768f', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');