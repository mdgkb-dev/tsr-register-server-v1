insert into anamneses_researches (id, research_id, domain_id)
values ('1ab79cc8-a23b-44a5-97ff-e24169b59a4a', 'e9f2300f-afb7-43e0-93b9-eb110edfa685', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');

alter table researches add column item_order int default 0;

UPDATE public.researches SET name = 'Характеристика эпизодов системной реакции' WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa689';
UPDATE public.researches SET item_order = 4 WHERE id = '8f2f58fd-38e2-4644-b8e7-e05794e838a8';
UPDATE public.researches SET item_order = 1 WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa683';
UPDATE public.researches SET item_order = 2 WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa686';
UPDATE public.researches SET item_order = 3 WHERE id = 'e9f2300f-afb7-43e0-93b9-eb110edfa685';
