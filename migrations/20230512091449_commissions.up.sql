create table commissions_templates
(
    id   uuid default uuid_generate_v4() not null primary key,
    name varchar,
    drug_regimen_id  uuid  references drug_regimens,
    drug_id  uuid  references drugs,
    volume  varchar
);

create table commissions
(
    id   uuid default uuid_generate_v4() not null primary key,
    item_date date,
    start_date date,
    end_date date,
    patient_id uuid  references patients,
    drug_regimen_id  uuid  references drug_regimens,
    drug_id  uuid  references drugs,
    volume  varchar
);

alter table commissions add number serial;
create unique index commissions_number_uindex on commissions (number);



create table doctors
(
    id   uuid default uuid_generate_v4() not null primary key,
    name varchar,
    position varchar
);

create table commissions_doctors
(
    id   uuid default uuid_generate_v4() not null primary key,
    commission_id uuid not null references commissions,
    doctor_id uuid not null references doctors,
    doctor_role varchar,
    item_order int default 0
);

create table commissions_doctors_templates
(
    id   uuid default uuid_generate_v4() not null primary key,
    commission_template_id uuid not null references commissions_templates,
    doctor_id uuid not null references doctors,
    doctor_role varchar,
    item_order int default 0
);

create table drug_applications
(
    id   uuid default uuid_generate_v4() not null primary key,
    item_date date,
    commission_id uuid not null references commissions
);

create table fund_contracts
(
    id   uuid default uuid_generate_v4() not null primary key
);

create table buy_contracts
(
    id   uuid default uuid_generate_v4() not null primary key,
    fund_contract_id uuid not null references fund_contracts
);


insert into public.commissions_templates (id, name, drug_regimen_id, drug_id, volume)
values  ('d49cafcc-3736-44d3-960b-285d4cb1c9d0', 'Заболевание СПИНАЛЬНАЯ МЫШЕЧНАЯ АТРОФИЯ', null, null, null),
        ('cbcbec00-8c0b-4823-8f9e-00d8d898fe69', 'Заболевание МУКОВИСЦИДОЗ', null, null, null);


insert into public.doctors (id, name, position)
values  ('97bc5bb2-f419-45fd-b638-713ab5d83202', 'Л.В.Дымнова', null),
        ('95f09014-59e7-4b76-833c-253d662a9fb8', 'Т.Н.Кекеева', 'Зав. медико-генетическим отделением'),
        ('77123f85-c0a4-4be5-9de9-5b4ca7fecc19', 'Е.Е.Якушина', 'Врач-педиатр'),
        ('e050ce80-ed93-4b8d-8016-5000567051b8', 'А.Б.Малахов', 'Главный внештатный детский специалист пульмонолог ДЗМ, д.м.н., проф.'),
        ('8626b62d-f150-4819-845e-b72db78122a3', 'В.В.Горев', 'Главный врач '),
        ('67a01de4-d47e-46a5-bcdc-a62e18a8a1a1', 'А.В.Власова', 'Зав. отделения клинической фармакологии, к.м.н.'),
        ('0fd6506d-7a93-4075-bac7-6baa2c8521bf', 'И.П.Витковская', 'Зам. главного врача по организационно-методической работе, к.м.н'),
        ('f252f3f2-543a-44c4-ba64-ba935eaa7d52', 'А.Е.Анджель', 'Зам.главного врача по медицинской части'),
        ('26fdc52f-db0f-44d3-88c5-b645300b07a8', 'О.Ю.Брунова', 'Зав.отделением реанимации и интенсивной терапии'),
        ('acb4ed10-2d5b-4588-b2f0-9b73629b0bef', 'Ю.В.Спасская', 'Ведущий юрисконсульт ГБУЗ «Морозовская ДГКБ ДЗМ»'),
        ('b64d56d4-574a-4555-b6a7-267e4d312bb6', 'О.И.Симонова', 'Руководитель службы муковисцидоза ГБУЗ «Морозовская ДГКБ ДЗМ», д.м.н., проф.'),
        ('f476887b-30cf-438d-b08d-d6c9c0834e5b', 'О.В.Высоколова', 'Врач-педиатр'),
        ('81a3eb6b-9362-480f-a783-c865e2b26349', 'М.А.Мухина', 'Врач-педиатр'),
        ('4a3f8c02-a2b3-4423-bd2a-e7869d9eb4e9', 'Н.С.Демикова', 'Главный внештатный детский специалист по медицинской генетике ДЗМ, д.м.н., проф'),
        ('3974effe-46d2-49c7-861f-3825aa6f978f', 'Ю.Е.Мартыненко', 'Врач-невролог'),
        ('739bf072-aecc-41f4-a47d-25954d4cbc4c', 'Н.А.Краснощекова', 'Зав. отделением наследственных нарушений обмена веществ');



insert into public.commissions_doctors_templates (id, commission_template_id, doctor_id, doctor_role, item_order)
values  ('44fc0442-f056-4744-b4d4-2def83e7162e', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '8626b62d-f150-4819-845e-b72db78122a3', 'Председатель', 0),
        ('891ad973-0da7-4889-aaa8-dd1fe91d36a5', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '0fd6506d-7a93-4075-bac7-6baa2c8521bf', 'Зам.председателя', 1),
        ('0138a013-6ae0-4456-ba2c-47559bd7d90d', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '97bc5bb2-f419-45fd-b638-713ab5d83202', 'Секретарь', 2),
        ('f6a4879c-7048-41ee-8f7b-fdd3cf46cc68', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '67a01de4-d47e-46a5-bcdc-a62e18a8a1a1', 'Член консилиума', 3),
        ('a133248b-20d2-4747-b98d-95fd72acfe6b', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '739bf072-aecc-41f4-a47d-25954d4cbc4c', 'Член консилиума', 4),
        ('56956ec9-3c66-415d-a8dd-f9bbf4f455d1', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '95f09014-59e7-4b76-833c-253d662a9fb8', 'Член консилиума', 5),
        ('8067fdb6-9e08-46b2-802a-78d01e3eed67', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '3974effe-46d2-49c7-861f-3825aa6f978f', 'Член консилиума', 6),
        ('b72f34e7-9f43-4f3b-9efc-b19b3903d216', 'd49cafcc-3736-44d3-960b-285d4cb1c9d0', '4a3f8c02-a2b3-4423-bd2a-e7869d9eb4e9', 'Пришлашённый специалист', 7);