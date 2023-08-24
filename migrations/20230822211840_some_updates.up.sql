-- Вопросы к диагнозу

create table mkb_questions_domains
(
    id          uuid default uuid_generate_v4() not null primary key,
    question_id uuid not null references questions,
    domain_id uuid not null references domains
);

create table mkb_researches
(
    id          uuid default uuid_generate_v4() not null primary key,
    research_id uuid not null references researches,
    domain_id uuid not null references domains
);

create table anamneses_researches
(
    id          uuid default uuid_generate_v4() not null primary key,
    research_id uuid not null references researches,
    domain_id uuid not null references domains
);

insert into anamneses_researches ("research_id", "domain_id")
select id, 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' from
    researches r
where r.id in (
               '8f2f58fd-38e2-4644-b8e7-e05794e838a8',
               'e9f2300f-afb7-43e0-93b9-eb110edfa683',
               'e9f2300f-afb7-43e0-93b9-eb110edfa686',
               'e9f2300f-afb7-43e0-93b9-eb110edfa688',
               'e9f2300f-afb7-43e0-93b9-eb110edfa689'
    );





insert into public.researches (id, name, with_dates, with_scores) values  ('c21eb5de-e180-4faa-b1eb-bfd77dddb470', 'Вопросы к диагнозу Анафилаксия', null, false);
insert into public.mkb_researches (id, research_id, domain_id) values  ('8fe4fec1-046b-4220-9e44-2c5f683d3f25', 'c21eb5de-e180-4faa-b1eb-bfd77dddb470', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');

INSERT INTO public.questions (id, name, item_order, value_type_id, age_compare)
VALUES
    ('fc171666-349d-4c79-86fc-cb6a2bfd2baf', 'Вероятность дианоза', 0, 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', false),
    ('fc171666-349d-4c79-86fc-cb6a2bfd2bad', 'Диагноз поставлен', 1, 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', false);

UPDATE public.questions SET research_id = 'c21eb5de-e180-4faa-b1eb-bfd77dddb470' WHERE id in('fc171666-349d-4c79-86fc-cb6a2bfd2baf', 'fc171666-349d-4c79-86fc-cb6a2bfd2bad');

insert into answer_variants (id, name, item_order, question_id, score, show_more_questions)
values  ('b508f7ac-54b7-411d-bf3c-9594f8e0bf2c', 'Плановая', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bb', 1, false),
        ('022fc517-81af-4c53-9207-89e7e3eb5aad', 'Подтверждён', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bc', 0, false),
        ('b508f7ac-54b7-411d-bf3c-9594f8e0bf2d', 'Не подтверждён', null, 'b54b5ce1-998b-4363-8ddb-e5dddba237bc', 1, false);


insert into public.answer_variants (id, name, item_order, question_id, score, show_more_questions)
values  ('7ca20145-8add-41cf-a9ea-eaccf579fb5f', 'При включении в регистр впервые', 0, 'fc171666-349d-4c79-86fc-cb6a2bfd2bad', null, false),
        ('31e2bb9c-3c6d-4764-bbf2-d97452deef10', 'Поставлен ранее', 1, 'fc171666-349d-4c79-86fc-cb6a2bfd2bad', null, false);

insert into public.mkb_questions_domains (id, question_id, domain_id)
values  ('3b835c1d-6c6e-43d5-b576-a9d8d97bc037', 'fc171666-349d-4c79-86fc-cb6a2bfd2baf', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('93d80e6e-7207-485a-8af4-a73d08e8c08d', 'fc171666-349d-4c79-86fc-cb6a2bfd2bad', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');
-- 52 больница
--         ('c356d559-5ddb-4bf8-bbde-f86758881059', 'fc171666-349d-4c79-86fc-cb6a2bfd2bad', 'c3424b60-da2c-4d49-8fc3-7a57a5a51377'),
--         ('2409cabc-412f-4d5e-b5d6-4e0727bc9f43', 'fc171666-349d-4c79-86fc-cb6a2bfd2bad', 'c3424b60-da2c-4d49-8fc3-7a57a5a51377');


-- Вопросы к анамнезу

create table anamnesis_researches_domains
(
    id          uuid default uuid_generate_v4() not null primary key,
    research_id uuid not null references researches,
    domain_id uuid not null references domains
);

insert into public.anamnesis_researches_domains (id, research_id, domain_id)
values  ('298be1c7-7bf4-4bd4-b5b9-33b6a5326101', '8f2f58fd-38e2-4644-b8e7-e05794e838a8', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('c1e3c26b-4ecc-4941-9b63-fbb6061dbf13', 'e9f2300f-afb7-43e0-93b9-eb110edfa683', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('a5c5ec1b-457b-4f10-8c54-7cb4cc5adb72', 'e9f2300f-afb7-43e0-93b9-eb110edfa685', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('df930bba-fd2f-4b94-b171-d32501709dbd', 'e9f2300f-afb7-43e0-93b9-eb110edfa686', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('215785bc-c2f4-43b6-961c-62a1557f19da', 'e9f2300f-afb7-43e0-93b9-eb110edfa688', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('4ffea145-6cc8-4650-9615-2f54c13b3bd3', 'e9f2300f-afb7-43e0-93b9-eb110edfa689', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');



delete from researches_pools_researches where research_id in
                                              (select id from researches
                                               where name in ('Медицинские документы',
                                                              'Наличие аллергических заболеваний у пациента',
                                                              'Наличие других сопутствующих заболеваний у пациента',
                                                              'Наличие аллергических заболеваний у ближайших родствеников пациента (мать/отец, родные братья/сестры)',
                                                              'Характеристика последующих эпизодов системной реакции',
                                                              'Характеристика первого эпизода системной реакции'
));

create table users_domains
(
    id          uuid default uuid_generate_v4() not null primary key,
    user_id uuid not null references users,
    domain_id uuid not null references domains
);

insert into public.users_domains (id, user_id, domain_id)
values  ('ad78abcf-b13e-4a0e-a5da-91c739412106', 'bc3bf514-834c-416b-a2e6-a349f4b4ecce', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('c2b198fb-2ac5-4d28-83b9-564da1489390', '7ef578b3-86c5-421f-8f6a-2a4de2cbf6dd', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('f4256b20-16ac-411c-a21b-c19d8a878475', '86f5fb02-f65c-431b-a7fd-849369d9c21b', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('4d0b026b-1128-42ed-a799-55e954463fed', '1c05d56c-5b98-4361-b42c-ebd5107cfb5e', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('088d5663-d5ce-4910-84f7-308c1d0cb302', '59f39975-83ab-434d-954c-8b6fd1ffa4a2', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('0aa0931c-a156-402c-a917-8f441c42533c', '66f6fe6b-4cbf-4bfa-96de-cdaadb4febbd', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('f330e6da-6a3f-47bf-ab75-5cd74df333d6', '342221bf-810d-4909-b68a-da939783fba1', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('c22077ec-5f48-4232-9db1-409e0ed19b08', '5d04ac28-70d0-4926-adca-5da55a6cf7e3', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('e8034332-5a97-4b5e-8367-a87ade914349', '5d04ac28-70d0-4926-adca-5da55a6cf7e4', '8669a507-5da0-4603-99b6-3e79e41f3f35');

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

alter table users drop  column domain_id;





create table patients_domains
(
    id          uuid default uuid_generate_v4() not null primary key,
    patient_id uuid not null references patients,
    domain_id uuid not null references domains
);

insert into patients_domains (patient_id, domain_id)
select p.id, ud.domain_id from patients p
                                   join users u on p.created_by_id = u.id
                                   join users_domains ud on u.id = ud.user_id;


-- Выносим пациентов 52 больницы
insert into public.domains (id, name)
values  ('c3424b60-da2c-4d49-8fc3-7a57a5a51377', '52hospital');

UPDATE public.users_domains SET domain_id = 'c3424b60-da2c-4d49-8fc3-7a57a5a51377' WHERE id = '088d5663-d5ce-4910-84f7-308c1d0cb302'

update patients_domains set domain_id = 'c3424b60-da2c-4d49-8fc3-7a57a5a51377'
where patient_id in (
    select p.id from patients p
                         join humans h on p.human_id = h.id
    where h.date_birth < '2005-08-23'
)


-- files

delete
from questions
where id in ('3df34118-cf2f-4a09-bca0-0428d93e0521', 'e053c6bf-9bc9-4e88-9a40-d058f923b6d2');


insert into questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, is_files_storage)
values
    ('3df34118-cf2f-4a09-bca0-0428d93e0525','Выписки', 0 ,'50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '8f2f58fd-38e2-4644-b8e7-e05794e838a8', true),
    ('3df34118-cf2f-4a09-bca0-0428d93e0527','Эпикризы', 1 ,'50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '8f2f58fd-38e2-4644-b8e7-e05794e838a8', true),
    ('3df34118-cf2f-4a09-bca0-0428d93e0528','Заключения', 2 ,'50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '8f2f58fd-38e2-4644-b8e7-e05794e838a8', true),
    ('3df34118-cf2f-4a09-bca0-0428d93e0529','Консультации', 3 ,'50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '8f2f58fd-38e2-4644-b8e7-e05794e838a8', true);

create table answer_files(
                             id uuid default uuid_generate_v4() not null primary key,
                             comment varchar,
                             answer_id uuid not null references answers,
                             file_info_id uuid not null references file_infos
);


alter table anamneses
    add column patient_id uuid references patients;

alter table anamneses
    add column mkb_item_id uuid references mkb_items;

alter table anamneses
drop column patient_diagnosis_id;


update anamneses a set patient_id = pd.patient_id
from patient_diagnosis pd where pd.id = a.patient_diagnosis_id;