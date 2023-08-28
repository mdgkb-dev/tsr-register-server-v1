DELETE FROM public.answer_variants WHERE id = '360cb808-2c0c-4cf3-8bde-3aec74485391';
DELETE FROM public.answer_variants WHERE id = 'f8cb9a13-1eeb-4f3b-9374-73b9d490f36d';

UPDATE public.answer_variants SET name = 'Да, были' WHERE id = '1c8d210f-4420-4c96-baff-239ff9f09f8e';



delete from questions where parent_id = '336ed8c9-9a42-4249-8bd4-469f05b7904b';

insert into public.questions (name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ( 'Переносите ли причинный продукт/лекарство/ фактор/др. в настоящее время', 2, '9f61f302-6821-40b9-94bc-78dedf955a11', false, false, null, null, null, null, null, null, null, null, '336ed8c9-9a42-4249-8bd4-469f05b7904b', null, null),
        ( 'Причины (продукт/лекарство/ фактор/др.) эпизодов', 1, '9f61f302-6821-40b9-94bc-78dedf955a11', false, false, null, null, null, null, null, null, null, null, '336ed8c9-9a42-4249-8bd4-469f05b7904b', null, null),
        ( 'Указать количество эпизодов', 0, '47affcc5-5d32-4b1f-bf07-33382ed06cda', false, false, null, null, null, null, null, null, null, null, '336ed8c9-9a42-4249-8bd4-469f05b7904b', null, null);


UPDATE questions SET item_order = 1 WHERE id = 'ddfdf3d0-cb10-4a17-a916-579b12a62dc3';




-- Только одно исследование с эпиздоами
update research_results set patient_research_id = pr1.id
from patients_researches pr
         join patients_researches pr1 on pr.patient_id = pr1.patient_id and pr1.research_id ='e9f2300f-afb7-43e0-93b9-eb110edfa686'
where pr.research_id = 'e9f2300f-afb7-43e0-93b9-eb110edfa688';

delete from patients_researches where patients_researches.research_id = 'e9f2300f-afb7-43e0-93b9-eb110edfa688';
delete
from anamneses_researches
where research_id = 'e9f2300f-afb7-43e0-93b9-eb110edfa688';

delete from selected_answer_variants s
where s.question_id in (select id from  questions where research_id = 'e9f2300f-afb7-43e0-93b9-eb110edfa688');

delete from questions where research_id = 'e9f2300f-afb7-43e0-93b9-eb110edfa688';


delete from researches where id = 'e9f2300f-afb7-43e0-93b9-eb110edfa688';

insert into public.questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ('84a53f45-2cfe-4adf-b127-17009e6a1e0b', 'Сканы', 0, '50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, 'a105c091-09e0-4ea3-9eff-41de5c3a3b8c', null, null, null, null, null, null, null, null, null, null),
        ('a0f0f35b-dcf5-4370-b9fa-95a5b3b9b4b4', 'Сканы', 0, '50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '1d58d174-c053-4c81-b358-b55e13289418', null, null, null, null, null, null, null, null, null, null),
        ('abb73192-2bcc-4e5a-95ee-9014b9ffa3a0', 'Сканы', 0, '50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '1d58d174-c053-4c81-b358-b55e13289417', null, null, null, null, null, null, null, null, null, null),
        ('d17ee17b-34de-4fce-9318-2a02641384e5', 'Сканы', 0, '50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '5a193bd3-d8c1-473b-ae80-0c1270e76c84', null, null, null, null, null, null, null, null, null, null);