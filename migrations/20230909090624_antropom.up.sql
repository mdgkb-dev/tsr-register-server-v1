insert into researches (id, name, with_dates, with_scores, is_laboratory, item_order)
values  ('dbad001c-6aff-4e5d-8d7f-0a5841d1b4f8', 'Окружность головы', true, false, true, 0),
        ('4e216374-eaa8-4ac6-a9f3-6d6bcb862938', 'Окружность груди', true, false, true, 0);

insert into public.researches_domains (id, research_id, domain_id)
values  ('30a74101-821c-4534-a8c7-976c767c60de', 'dbad001c-6aff-4e5d-8d7f-0a5841d1b4f8', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('bea205fd-7c07-48ba-98a0-15221a453853', '4e216374-eaa8-4ac6-a9f3-6d6bcb862938', '8669a507-5da0-4603-99b6-3e79e41f3f35');

insert into public.questions (id, name, item_order, value_type_id, with_other, research_id, short_name, age_compare, code, calculate_scores, parent_id, comment)
values  ('4751326c-fd41-4ffe-8fd5-c83db463ce62', 'Дата измерения', 0, 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, 'dbad001c-6aff-4e5d-8d7f-0a5841d1b4f8', null, null, null, null, null, null),
        ('0802f61d-460e-4cde-a59e-a866ed0bfc96', 'Окружность', 1, '47affcc5-5d32-4b1f-bf07-33382ed06cda', false, 'dbad001c-6aff-4e5d-8d7f-0a5841d1b4f8', null, null, null, null, null, null),
        ('451b56c5-52fa-4a29-b3fe-faf0d9d863cb', 'Дата измерения', 0, 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, '4e216374-eaa8-4ac6-a9f3-6d6bcb862938', null, null, null, null, null, null),
        ('b72a436c-f1aa-448a-9286-96cb8b116a2a', 'Окружность', 1, '47affcc5-5d32-4b1f-bf07-33382ed06cda', false, '4e216374-eaa8-4ac6-a9f3-6d6bcb862938', null, null, null, null, null, null);

UPDATE public.questions SET code = 'circle' WHERE id = 'b72a436c-f1aa-448a-9286-96cb8b116a2a';
UPDATE public.questions SET code = 'circle' WHERE id = '0802f61d-460e-4cde-a59e-a866ed0bfc96';

INSERT INTO public.formulas (id, name, formula, research_id, age_relation, sex_relation, color, xlsx) VALUES (DEFAULT, 'Окружность', 'circle', 'dbad001c-6aff-4e5d-8d7f-0a5841d1b4f8', false, false, '#1c5920', false);
INSERT INTO public.formulas (id, name, formula, research_id, age_relation, sex_relation, color, xlsx) VALUES (DEFAULT, 'Окружность', 'circle', '4e216374-eaa8-4ac6-a9f3-6d6bcb862938', false, false, '#1c5920', false);

alter table document_type_fields add column item_order numeric default 0;