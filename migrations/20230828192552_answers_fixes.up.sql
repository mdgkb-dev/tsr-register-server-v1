delete from questions q
where q.parent_id ='d8c2be17-bace-45c1-8eda-d76f16b27a51' and id != 'd22758e7-2ba4-487d-8e5f-420b4add0e01';

update answers set question_id = 'd22758e7-2ba4-487d-8e5f-420b4add0e01'
where id =
      ('536ffe61-89bc-4714-954a-742d321f8447') or id =
                                                  ('3f4a6289-b8f1-44b8-a12b-452f0a8f0438')or id =
                                                                                             ('9ec00233-1e63-490c-8dc4-7cac53a749a8')or id =
                                                                                                                                        ('de0f189b-d990-451d-891a-01824f567d24')or id =
                                                                                                                                                                                   ('83982a23-5356-4b5a-a65d-e098d9574d8a')or id =
                                                                                                                                                                                                                              ('6609ac2f-ab0d-4ad4-b9b7-4197e82f99c8')or id =
                                                                                                                                                                                                                                                                         ('52496603-0715-4131-a295-46d170393e0b')or id =
                                                                                                                                                                                                                                                                                                                    ('4336f066-a0db-427c-8769-0a857468301a');



update
    answer_variants set show_more_questions = true where question_id = 'd8c2be17-bace-45c1-8eda-d76f16b27a51';


UPDATE questions SET value_type_id = 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e' WHERE id = '9df2a216-ded6-45b5-84c1-d928cf86fd59';
UPDATE questions SET value_type_id = 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e' WHERE id = 'e0409d51-f8cf-492b-b4b4-b8fd59e0f183';

UPDATE answer_variants SET item_order = 1 WHERE id = 'c096c263-7680-424c-b457-b98ada7e26a7';
UPDATE answer_variants SET item_order = 2 WHERE id = '21984435-4e29-4253-bcbb-87cc2f40c53c';


UPDATE public.researches SET name = 'Лабораторная диагностика аллергочип' WHERE id = '5a193bd3-d8c1-473b-ae80-0c1270e76c84';
insert into public.researches (id, name, with_dates, with_scores) values  ('986e4b00-0519-4c81-9447-36dd7858465c', 'Генетическое исследование', true, false);

insert into public.questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ('30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', 'Время забора материала относительно острого периода анафилаксии', 3, 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', false, false, '986e4b00-0519-4c81-9447-36dd7858465c', '', 'Время забора материала', null, null, null, null, null, null, null, null),
        ('5b04a4d9-6de6-4988-af6f-d6e344d44aa7', 'Дата забора крови', 2, 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, false, '986e4b00-0519-4c81-9447-36dd7858465c', '', 'Дата забора крови', null, null, null, null, null, null, null, null),
        ('78b18f4e-033b-42c9-9c0e-3de10c50d235', 'Дата внесения в реестр', 1, 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, false, '986e4b00-0519-4c81-9447-36dd7858465c', '', 'Дата внесения в реестр', null, null, null, null, null, null, null, null),
    ('dc58af7b-0ac7-41ef-8f39-e6edc709c86e', 'Сканы', 0, '50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '986e4b00-0519-4c81-9447-36dd7858465c', null, null, null, null, null, null, null, null, null, null);

insert into public.answer_variants (id, name, item_order, question_id, score, show_more_questions)
values  ('40afe67e-0970-4f41-ac94-de395793e367', '30 минут', 0, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('2864ea56-e762-4700-b72a-f68ddc8c862c', '1 час', 1, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('8851ceb3-f1a8-47f7-90d7-8357895f963f', '2 часа', 2, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('7401cbd5-529a-41c4-a8d6-acc455c1ce20', '3 часа', 3, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('baa2a201-647b-4058-a451-550c57fadba5', '4-6 часов', 4, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('bdcc96d0-bbc6-4f2c-8389-fade9d1e4e96', '6-8 часов', 5, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('be748ef2-7011-4d0b-9d25-1ad439177f4a', 'Более 8-ми часов', 6, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false),
        ('0d0b04ea-c31e-4317-8b84-aea1ee7089eb', 'Вне острого периода анафилаксии', 7, '30bfc740-e8fb-49d5-bbe8-d362b7a14ae6', null, false);

insert into public.researches_pools_researches (id, researches_pool_id, research_id, item_order)
values  ('e236049d-238a-439d-a2a0-d7d22fcae2e2', 'a0680b34-0d9b-4df2-9288-fbc421fd3ee5', '986e4b00-0519-4c81-9447-36dd7858465c', 9);