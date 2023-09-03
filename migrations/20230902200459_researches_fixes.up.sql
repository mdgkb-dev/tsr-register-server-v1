insert into researches_domains (id, research_id, domain_id)
values  ('e02267e4-cffd-42bc-a441-e75030d73884', '585262b4-369e-40ed-839d-b8567129b9fb', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('a784da20-f159-4701-b930-a8b750fcbdc4', '20d28aac-2f46-4bfb-9d2f-99c1e7d0da5e', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('d00a1794-995e-4684-998a-762dd9ede629', 'bc982ef6-1f15-4905-80e8-b6d6bb30e37b', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('9b29b370-f575-4f57-acbd-0cab84f3a58e', '1ea978e6-cd5f-45e3-a1ed-86906dc5037a', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('5cdcdfe4-7197-4eac-bf1d-23dee872a78f', 'e9f2300f-afb7-43e0-93b9-eb110edfa682', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('7a8a9b2e-5d2b-44b6-a9e9-4382694a844f', '7cdff90d-6f9a-4de8-aa47-e01d35d3d850', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('bcf7b407-c5f1-4a31-b82c-207fd173c308', 'af81a162-90fa-40f2-801c-f8eea0bfa947', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('db3343f9-3664-417e-a3ec-4a3148041351', 'c5d482c7-ae56-4b50-a895-f717bf731f01', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('96bd7378-4fd2-4695-8ef4-595f39ec97a6', 'c5d482c7-ae56-4b50-a895-f717bf731f02', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('6c853df9-dfbf-4527-a3c1-28d68ac5430b', '6be7f5c3-434f-40be-99f2-027c41e78853', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('3d1c22c7-c4b4-402c-ba23-1a590749b8ec', 'ca4f2407-ec10-410d-b47a-267d74cbc7cb', '8669a507-5da0-4603-99b6-3e79e41f3f35'),
        ('b4afcf98-94e0-4639-bc91-0e4dd5e99233', 'a105c091-09e0-4ea3-9eff-41de5c3a3b8c', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('215ded4a-61a2-44a2-b40c-2a8f57289bc3', '1d58d174-c053-4c81-b358-b55e13289417', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('7c41d558-ae96-459a-9e2e-70af8ceda1a1', '1d58d174-c053-4c81-b358-b55e13289418', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('67e15b1a-625d-46a7-9676-6d228acef6b1', '5a193bd3-d8c1-473b-ae80-0c1270e76c84', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949'),
        ('95e9a51f-ec7c-4d7a-8ab0-70be360441d7', '986e4b00-0519-4c81-9447-36dd7858465c', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');



insert into researches (id, name, with_dates, with_scores)
values  ('dc000421-4a9b-40f4-a6ad-a18b08e5cd6e', 'Вопросы к диагнозу СМА', null, false);

insert into mkb_researches (id, research_id, domain_id)
values  ('9751832c-2d44-4d45-b115-b0216d6390ed', 'dc000421-4a9b-40f4-a6ad-a18b08e5cd6e', '8669a507-5da0-4603-99b6-3e79e41f3f35');

insert into researches_domains (id, research_id, domain_id)
values  ('089be2ab-b34b-41f4-98e5-d59a241cf794', 'dc000421-4a9b-40f4-a6ad-a18b08e5cd6e', '8669a507-5da0-4603-99b6-3e79e41f3f35');

insert into questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ('bf7209d4-dced-4c5e-9645-40492136d009', 'Основной/сопуствующий диагноз', 0, 'fc00cc5a-f7a5-4974-ad57-9432656d5e0e', false, false, 'dc000421-4a9b-40f4-a6ad-a18b08e5cd6e', null, null, null, null, null, null, null, null, null, null),
        ('276b618a-d92b-4702-8886-3a2cb3fe1df8', 'Дата постановки', 1, 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, false, 'dc000421-4a9b-40f4-a6ad-a18b08e5cd6e', null, null, null, null, null, null, null, null, null, null),
        ('041f3c96-c795-4724-b86f-3a3ac4f1636b', 'Где выставили', 2, '9f61f302-6821-40b9-94bc-78dedf955a11', false, false, 'dc000421-4a9b-40f4-a6ad-a18b08e5cd6e', null, null, null, null, null, null, null, null, null, null);

insert into answer_variants (id, name, item_order, question_id, score, show_more_questions)
values  ('427543d6-17c7-4cc7-96e6-c09f79263411', 'Основной', 0, 'bf7209d4-dced-4c5e-9645-40492136d009', null, false),
        ('622987c3-98a3-48ff-894f-61d2a383fe37', 'Сопутствующий', 1, 'bf7209d4-dced-4c5e-9645-40492136d009', null, false);

UPDATE questions SET name = 'Вероятность диагноза' WHERE id = 'fc171666-349d-4c79-86fc-cb6a2bfd2baf';


insert into questions (id, name, item_order, value_type_id, with_other, with_dates, research_id, tag, short_name, col_width, age_compare, is_files_storage, code, calculate_scores, parent_id, comment, domain_id)
values  ('6ebabec7-b5f8-40a7-a054-010964441fd7', 'Документ', 6, '50dfc0a4-b260-4c63-b16f-4119e152037f', false, false, '6be7f5c3-434f-40be-99f2-027c41e78853', null, null, null, null, null, null, null, null, null, null);

UPDATE questions SET name = 'Патогенетическая терапия', item_order = 3 WHERE id = 'ec3cfea0-29d7-4308-9f2a-fa3cee71a7ff';
UPDATE questions SET item_order = 2 WHERE id = 'bd35e4b0-8801-4b33-9102-ae0e4aa933bb';

UPDATE questions SET name = 'Дата введения/применения лекарственного препарата' WHERE id = 'e13c50a5-731f-4a2c-a80a-6e7be4bcc624';

insert into document_types (id, name, required, code)
values  ('5b00fbed-f92d-4524-8fdc-76890fef8210', 'Полис ОМС', false, 'polis_oms');

insert into document_type_fields (id, name, "order", type, document_type_id, value_type_id, required, code)
values  ('93fde2a5-bcda-4337-85aa-c5476feeae79', 'Номер', 0, 'string', '5b00fbed-f92d-4524-8fdc-76890fef8210', '9f61f302-6821-40b9-94bc-78dedf955a11', false, 'number'),
        ('2393f212-02ac-4ec0-a32f-6638e361a6b6', 'Страховая компания', 1, 'string', '5b00fbed-f92d-4524-8fdc-76890fef8210', '9f61f302-6821-40b9-94bc-78dedf955a11', false, 'company');

insert into public.document_types (id, name, required, code)
values  ('5b00fbed-f92d-4524-8fdc-76890fef8212', 'Документ регистрации по месту проживания', false, 'registration_temp_life'),
        ('5b00fbed-f92d-4524-8fdc-76890fef8211', 'Документ регистрации по месту жительства', false, 'registration_life');


insert into document_type_fields (id, name, "order", type, document_type_id, value_type_id, required, code)
values
        ('93fde2a5-bcda-4337-85aa-c5476feeae80', 'Номер', 0, 'string', '5b00fbed-f92d-4524-8fdc-76890fef8211', '9f61f302-6821-40b9-94bc-78dedf955a11', false, 'number'),
        ('2393f212-02ac-4ec0-a32f-6638e361a6b7', 'Дата выдачи', 1, 'date', '5b00fbed-f92d-4524-8fdc-76890fef8211', 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, 'date'),
        ('93fde2a5-bcda-4337-85aa-c5476feeae81', 'Номер', 0, 'string', '5b00fbed-f92d-4524-8fdc-76890fef8212', '9f61f302-6821-40b9-94bc-78dedf955a11', false, 'number'),
        ('2393f212-02ac-4ec0-a32f-6638e361a6b8', 'Дата выдачи', 1, 'date', '5b00fbed-f92d-4524-8fdc-76890fef8212', 'efdd456c-091b-49d9-ac32-d0d345f88e64', false, 'date');


UPDATE questions SET item_order = 0 WHERE id = 'e832640f-e824-4d21-b32c-da97763056cc';
UPDATE questions SET item_order = 1 WHERE id = '7c99e32f-bf5a-403d-bc83-77fe577a3835';