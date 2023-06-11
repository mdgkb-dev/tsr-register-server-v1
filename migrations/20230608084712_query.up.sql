create table research_query_group_questions
(
    id                uuid    default uuid_generate_v4() not null primary key,
    research_query_group_id uuid references research_query_groups on delete cascade,
    question_id       uuid references questions on delete cascade,
    item_order        integer default 0,
    every_radio_set   boolean,
    count_percentage  boolean,
    aggregate_type    varchar
);

-- data small fix
insert into answers (id, question_id,  research_result_id)
select answer_id, question_id, research_result_id from selected_answer_variants sav
join answer_variants av on sav.answer_variant_id = av.id
join questions q on av.question_id = q.id
on conflict do nothing ;
----
delete from research_query_groups where id is not null;
--
insert into public.research_query_groups (id, research_query_id, research_id, item_order, aggregate_type, count_sum, count_percentage, name)
values  ('ffae64ef-6398-4e17-b15f-191a595d9bb3', '9e3b05b1-3007-4cc1-9d65-58d3211e7153', 'e9f2300f-afb7-43e0-93b9-eb110edfa683', 1, 'none', true, true, 'Наличие аллергических заболеваний по нозологиям'),
        ('42a17103-8a74-4e7e-b2db-ac4f1bd3a79a', '9e3b05b1-3007-4cc1-9d65-58d3211e7153', 'e9f2300f-afb7-43e0-93b9-eb110edfa683', 0, 'existing', true, true, 'Наличие аллергических заболеваний в целом'),
        ('5c2021e0-da55-4bee-a75a-9504ac8390cf', '9e3b05b1-3007-4cc1-9d65-58d3211e7153', 'e9f2300f-afb7-43e0-93b9-eb110edfa685', 2, 'existing', true, true, 'Наличие аллергических заболеваний у ближайших родственников в целом'),
        ('d68dab9f-cda8-4f4c-ad57-f11e03fc6073', '9e3b05b1-3007-4cc1-9d65-58d3211e7153', 'e9f2300f-afb7-43e0-93b9-eb110edfa685', 3, 'none', true, true, 'Наличие аллергических заболеваний у ближайших родственников  по нозологиям'),
        ('a28bd637-e5cb-4fd6-9ee1-c66265a7b8d8', '9e3b05b1-3007-4cc1-9d65-58d3211e7153', 'e9f2300f-afb7-43e0-93b9-eb110edfa686', 4, 'none', true, true, 'Наличие других заболеваний у пациента	'),
        ('f1afb502-50ab-410e-9f39-7bbbcf938593', '9e3b05b1-3007-4cc1-9d65-58d3211e7153', 'e9f2300f-afb7-43e0-93b9-eb110edfa688', 5, 'none', true, true, 'Характеристика первого эпизода реакции');
--
insert into public.research_query_group_questions (id, research_query_group_id, question_id, item_order, every_radio_set, count_percentage, aggregate_type)
values  ('d8ae6890-3f40-473e-a7fc-48e5cfceb29e', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', 'fbf4d365-1f26-4a91-811e-87544021ca3c', 0, null, true, 'none'),
        ('4c240870-7535-4f03-b262-59dfb79662f4', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', 'eb44039a-1e64-41e6-b230-40d15f4a32ee', 5, null, true, 'none'),
        ('695301e6-295e-47bc-a225-7126a1078880', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '4497288a-b554-43af-9892-607c456197c1', 9, null, true, 'none'),
        ('0001ee5c-ad62-43b0-81e2-7eb102d96d08', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '16086229-4bc2-47cd-989e-cd5f6409299d', 8, null, true, 'none'),
        ('4df0ffab-b6a3-4955-a4fa-99640777cd02', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '569f6a06-5eee-4db6-b59b-9f4f5d6c6d17', 3, null, true, 'none'),
        ('2a8590dc-56cc-4552-afd2-5b2b57268b16', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', 'd44eec36-5443-442b-9e77-cb1a963366f0', 1, null, true, 'none'),
        ('ebe6e5db-48d1-46ab-a6c2-3bf9e053babb', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '14833a0d-ca0f-4ecb-94ff-035fc9d6b2ca', 4, null, true, 'none'),
        ('d9d1965a-bbb4-4c7c-91e0-cc52c51a4fb7', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '0d445112-4cbc-4bbb-af92-e90c8e07c946', 7, null, true, 'none'),
        ('93b58223-bae4-4e44-bac2-f43683995d4d', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '32aa7062-8ebe-4774-9208-5111537abeca', 2, null, true, 'none'),
        ('492dcda2-283f-40dd-bb8a-dface895c21a', 'ffae64ef-6398-4e17-b15f-191a595d9bb3', '336ed8c9-9a42-4249-8bd4-469f05b7904b', 6, null, true, 'none'),
        ('71fefda2-a273-4c4f-924b-cd81bab3a953', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a10', 0, null, true, 'none'),
        ('3c6fce64-05af-4584-9417-c47c0d9afde9', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a11', 1, null, true, 'none'),
        ('6ece9273-1ff8-4801-8359-155116c4835f', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a12', 2, null, true, 'none'),
        ('ec6bddce-1820-4259-b443-f250c770cb9a', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a13', 3, null, true, 'none'),
        ('57227a3f-d5f2-4aa1-8cb7-a85c2426bd9b', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'b46c277a-3866-4b6f-aee1-7f367d5a7a9b', 4, null, true, 'none'),
        ('1e987a16-ea39-4479-b77d-932e7ca8c684', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a15', 5, null, true, 'none'),
        ('419d5303-d63d-47e6-8039-435613075a79', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a16', 6, null, true, 'none'),
        ('c60960dc-66b5-48c9-bf9e-d000bcfd641d', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a17', 7, null, true, 'none'),
        ('b117e2e1-5007-496e-b2e0-e415d65584bb', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a18', 8, null, true, 'none'),
        ('58d0ec5a-f7b6-4f2a-9290-99abf53c71ab', 'd68dab9f-cda8-4f4c-ad57-f11e03fc6073', 'd8c2be17-bace-45c1-8eda-d76f16b27a19', 9, null, true, 'none'),
        ('41e36597-97c1-4647-a546-2651be2335cc', 'a28bd637-e5cb-4fd6-9ee1-c66265a7b8d8', 'd8c2be17-bace-45c1-8eda-d76f16b27a20', 0, null, true, 'none'),
        ('dd96cb91-cbff-4085-98ce-eb6b8fc5a6a2', 'a28bd637-e5cb-4fd6-9ee1-c66265a7b8d8', 'ddfdf3d0-cb10-4a17-a916-579b12a62dc3', 1, null, true, 'none'),
        ('172e54a0-d5b3-4fb6-a63f-ffaeaf73a120', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a46', 0, true, true, 'none'),
        ('f09d0cd0-48bb-4f58-9405-f22033406706', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a26', 1, true, true, 'none'),
        ('ca5fbac6-51f3-4936-a0ff-3013d4475861', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a31', 6, true, true, 'none'),
        ('a32e349b-d9f1-466a-8a6e-cd5325df4bf5', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a32', 7, true, true, 'none'),
        ('501be31c-6e00-4173-89d3-d1f7cb142215', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a33', 8, true, true, 'none'),
        ('433f211c-898d-466d-a45b-c321a71374de', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a34', 9, true, true, 'none'),
        ('ebd6dc72-1086-4707-9892-403ee5660cbe', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd9031d5d-70b1-4893-b7aa-bf034b0c6db4', 10, true, true, 'none'),
        ('b4979115-6a24-4269-8ee9-a32d7c7ffee2', 'f1afb502-50ab-410e-9f39-7bbbcf938593', '2cce0bf8-dbd5-4de9-bf00-d077bde92ce6', 11, true, true, 'none'),
        ('d3925b2e-3baf-409a-ad86-5f248440f2c3', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'c1b5bcac-996e-4546-bafc-973388e4efa7', 12, true, true, 'none'),
        ('73e73f4a-f29d-4496-a485-5d3a9141ae47', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'b3db397e-73e0-4507-ad2b-e21baee49679', 13, true, true, 'none'),
        ('f3f30dd1-2328-4a96-af84-015465ffb95d', 'f1afb502-50ab-410e-9f39-7bbbcf938593', '583d55ff-e081-4923-afef-270324ebd114', 14, true, true, 'none'),
        ('950053ec-b0f2-4038-8e53-10d48c2b48bb', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'a9b4860a-49b5-4560-bcff-1b6581565315', 15, true, true, 'none'),
--         ('06ef5441-dfc5-414a-b060-0191c294c63a', 'f1afb502-50ab-410e-9f39-7bbbcf938593', '44d6b77d-b916-46d2-85bd-af5c9cbd7d76', 16, true, true, 'none'),
        ('019082a2-378c-4b39-afa2-727d3b7fd118', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'ffcc0e13-7853-4a5e-b5dd-0d8b988b74f5', 17, true, true, 'none'),
        ('9b1f4e29-2957-4f33-9be5-ed18db402720', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a35', 18, true, true, 'none'),
        ('826fa2ab-3d19-47f3-ae83-ab88668a0491', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a36', 19, true, true, 'none'),
        ('11f4c0f8-75bd-409d-aa89-c505ab8d6908', 'f1afb502-50ab-410e-9f39-7bbbcf938593', '3cd87551-1248-4776-a6ca-95fe7b5f6a25', 20, true, true, 'none'),
        ('4f358887-d8a0-43f6-8fa0-abc73af840e9', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a42', 21, true, true, 'none'),
        ('6b8b4ce0-b385-4375-bb6a-39a10438e544', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a38', 22, true, true, 'none'),
        ('9c94523e-5ef2-4357-bd46-8c8c8b14981c', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'dd46e764-14fc-40d2-808f-11405cf781aa', 23, true, true, 'none'),
        ('e26bb420-e70d-46bf-857b-f8362e8cd1e0', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a37', 24, true, true, 'none'),
        ('ddea70fd-e779-4f86-9ad7-07b6f97b45ab', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'fc38a49a-a74a-48bc-b236-10cbd70eea52', 25, true, true, 'none'),
        ('a5d3814e-e54c-42af-9eb0-782c84d39262', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a40', 26, true, true, 'none'),
        ('af16d952-b326-46a5-acd6-4a74de760d0c', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a41', 27, true, true, 'none'),
        ('d524cd85-98c9-4ab9-91a9-87f256212c78', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd013ec4a-c995-4eb9-9562-5147071bf823', 28, true, true, 'none'),
        ('628d47d9-0d96-4f7a-acc0-8506b0d8a94f', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'ff9255d6-9c13-44a6-9b9b-55136392e73e', 29, true, true, 'none'),
        ('50c6ac9a-290d-4ac9-bafb-43d87b8d5164', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a43', 30, true, true, 'none'),
        ('b4c063b2-62a1-4e39-848b-c4ac08fdbb52', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a44', 31, true, true, 'none'),
        ('b48de3f6-da0e-43ec-b79e-3b085556324f', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a45', 32, true, true, 'none'),
        ('580e7288-0f32-4e85-b64e-50eb560944ac', 'f1afb502-50ab-410e-9f39-7bbbcf938593', '52fb3017-99f8-4a56-aea2-1268561b6fac', 33, true, true, 'none'),
        ('77132ed6-d6d2-48bd-ab8c-ed840af9087c', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a29', 4, false, true, 'none'),
        ('615332bb-c210-433c-bb31-8d76167873c2', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a30', 5, false, true, 'none'),
        ('3f0fe80a-4297-4387-a1d8-8887dba94334', 'f1afb502-50ab-410e-9f39-7bbbcf938593', '271ac938-852d-42b7-8798-09f829e79e87', 3, false, true, 'none'),
        ('49649385-c8bd-4b43-a24e-3117c68737a2', 'f1afb502-50ab-410e-9f39-7bbbcf938593', 'd8c2be17-bace-45c1-8eda-d76f16b27a28', 2, false, true, 'none');