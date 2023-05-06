INSERT INTO public.researches (id, name, with_dates) VALUES ('c5d482c7-ae56-4b50-a895-f717bf731f01', 'Шкала CHOP INTEND', true);
INSERT INTO public.researches_pools_researches (id, researches_pool_id, research_id, item_order) VALUES (DEFAULT, 'b3aa8d37-58a8-4655-a613-f238d1cf1120', 'c5d482c7-ae56-4b50-a895-f717bf731f01', 1);
alter table questions add comment varchar;

insert into questions(id, name, comment, item_order, research_id, value_type_id, calculate_scores)
select id, name, description, item_order, 'c5d482c7-ae56-4b50-a895-f717bf731f01', 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', true from chop_scale_questions;

insert into answer_variants (id, name, item_order, question_id, score, show_more_questions)
select id, name, score, chop_scale_question_id, score, false from chop_scale_question_scores;

INSERT INTO public.researches (id, name, with_dates) VALUES ('c5d482c7-ae56-4b50-a895-f717bf731f02', 'Шкала HMFSE', true);
INSERT INTO public.researches_pools_researches (id, researches_pool_id, research_id, item_order) VALUES (DEFAULT, 'b3aa8d37-58a8-4655-a613-f238d1cf1120', 'c5d482c7-ae56-4b50-a895-f717bf731f02', 2);

insert into questions(id, name, comment, item_order, research_id, value_type_id, calculate_scores)
select id, name, description, item_order, 'c5d482c7-ae56-4b50-a895-f717bf731f02', 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', true from hmfse_scale_questions;

insert into answer_variants (id, name, item_order, question_id, score, show_more_questions)
select id, name, score, hmfse_scale_question_id, score, false from hmfse_scale_question_scores;

drop view mkb_flat_view;
drop view mkb_diagnosis_view;
drop view mkb_groups_view;
drop view mkb_sub_diagnosis_view;

create view mkb_items_view(id, full_name) as
SELECT mi.id,
       case
           when (mi.code != '-') then concat(mi.code, ' ' ,mi.range_start, '-', mi.range_end, ' ', mi.name)
           when (mi.code = '-') then concat(mi.range_start, '-', mi.range_end, ' ', mi.name)
           when (mi.range_start = '-' and mi.range_end = '-') then concat(mi.code, ' ', mi.name)
           end
FROM mkb_items mi;
UPDATE public.search_groups SET key = 'mkbItem', search_group_table = 'mkb_items_view' WHERE id = 'ffcdd1b9-6de4-479b-becd-035fb450a45d';

alter table patient_diagnosis add doctor_name varchar;


alter table researches add with_scores bool default false;
update researches set with_scores = true
where name in ('Шкала CHOP INTEND', 'Шкала HMFSE');