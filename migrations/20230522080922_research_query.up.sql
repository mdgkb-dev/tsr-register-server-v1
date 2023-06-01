alter table register_queries rename column register_id to researches_pool_id;
alter table register_queries drop constraint register_queries_register_id_fkey;
alter table register_queries rename to research_queries;

alter table register_query_group_properties rename to research_query_questions;
alter table research_query_questions rename column register_query_group_id to research_query_group_id;
alter table research_query_questions     rename column register_property_id to question_id;

alter table register_query_groups rename column register_query_id to research_query_id;
alter table register_query_groups rename to research_query_groups;
alter table research_query_groups rename column register_group_id to research_id;
alter table research_query_groups add name varchar;

alter table register_query_to_register_property rename to register_query_questions;

-- исправление вопросов
update questions
set item_order = item_order - 1
where id  in (select questions.id from
    researches_pools rp
        join researches_pools_researches rpr on rp.id = rpr.researches_pool_id
        join researches on rpr.research_id = researches.id
        join questions on researches.id = questions.research_id
              where rp.id = 'a0680b34-0d9b-4df2-9288-fbc421fd3ee4');

update questions
set item_order = 2
where id = '179ab020-0f90-4e49-849d-62f3c580a721';

insert into questions (name, item_order, value_type_id,  research_id, short_name)
values ('Экзонов 7-8 гена SMN1', 1, '47affcc5-5d32-4b1f-bf07-33382ed06cda', '6be7f5c3-434f-40be-99f2-027c41e78853', 'Экзонов');

UPDATE public.researches SET name = 'Молекулярно-генетическое подтверждение' WHERE id = '6be7f5c3-434f-40be-99f2-027c41e78853';


insert into questions (name, item_order, value_type_id,  research_id, short_name)
values ('Дата исследования', 3 , 'efdd456c-091b-49d9-ac32-d0d345f88e64', '6be7f5c3-434f-40be-99f2-027c41e78853', 'Дата');

insert into questions (name, item_order, value_type_id,  research_id, short_name)
values ('Место исследования', 4 , '9f61f302-6821-40b9-94bc-78dedf955a11', '6be7f5c3-434f-40be-99f2-027c41e78853', 'Место исследования');


UPDATE questions SET item_order = 3 WHERE id = 'bd35e4b0-8801-4b33-9102-ae0e4aa933bb';

UPDATE public.answer_variants SET item_order = 0 WHERE id = '3d55f36c-40f9-414b-b905-ddafa87ad40c';
UPDATE public.answer_variants SET item_order = 2 WHERE id = '826a57bc-de4e-4e38-a4fb-9513827ec637';
UPDATE public.answer_variants SET item_order = 3 WHERE id = 'a475cae2-bba6-4471-9bd3-0982ad36bfc1';
UPDATE public.answer_variants SET item_order = 1 WHERE id = '6a67752d-361a-41bc-92a6-5903341db65d';



insert into questions (id, name, item_order, value_type_id, research_id)
values ('5d3b9a8a-9750-4caf-a0bf-7bbc0047400a','Если проводилась замена лекарственного препарата, то что явилось причиной', 5, '6fd80c4c-ece3-479c-94b6-005ccebcfe73', '7cdff90d-6f9a-4de8-aa47-e01d35d3d850');

INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('6a5c57b6-9c77-49b2-b5e6-6315f800c039', 'Наличине побочных явлений, зарегистрированных в Фармаконадзоре', 0, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400a', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('2abbbc2e-15bf-4f7d-84a2-aa03c487830a', 'Неэффективность проводимой терапии', 2, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400a', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('957af688-901e-4b5a-b1da-40c87f72e32a', 'Отказ пациента', 3, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400a', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('9041d73b-644f-4f94-b211-823942da67c4', 'Риск проводимой терапии превышает её пользу', 4, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400a', null, false);

insert into questions (id, name, item_order, value_type_id, research_id)
values ('5d3b9a8a-9750-4caf-a0bf-7bbc0047400f','Если проводилась замена лекарственного препарата, то что на какой препарат заменили', 6, '6fd80c4c-ece3-479c-94b6-005ccebcfe73', '7cdff90d-6f9a-4de8-aa47-e01d35d3d850');


INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('6a5c57b6-9c77-49b2-b5e6-6315f800c03a', 'Спинраза (МНН нусинерсен)', 0, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400f', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('2abbbc2e-15bf-4f7d-84a2-aa03c487830d', 'Эврисди (МНН рисдиплам)', 2, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400f', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('957af688-901e-4b5a-b1da-40c87f72e32f', 'Не получал', 3, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400f', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('2839eccd-5b99-4824-8884-1ad2a99c9bca', 'Золгенсма (МНН онасемноген абепарвовек)', 1, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400f', null, false);
INSERT INTO public.answer_variants (id, name, item_order, question_id, score, show_more_questions) VALUES ('9041d73b-644f-4f94-b211-823942da67c5', 'Другое', 4, '5d3b9a8a-9750-4caf-a0bf-7bbc0047400f', null, false);