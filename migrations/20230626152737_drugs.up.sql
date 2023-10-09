alter table drugs rename column name_mnn to name_inn;
alter table drugs drop column form;
alter table drugs drop column doze;
alter table drugs drop column registered;
alter table drugs drop column date_registration;

delete  from drugs where id is not null;



create table drug_forms
(
    id          uuid      default uuid_generate_v4() not null primary key ,
    name              varchar,
    report_name          varchar,
    date_registration date,
    drug_id uuid REFERENCES drugs(id) ON UPDATE CASCADE ON DELETE CASCADE
);

create table drug_dozes
(
    id          uuid      default uuid_generate_v4() not null primary key ,
    name              varchar,
    quantity     numeric     ,
    drug_form_id uuid REFERENCES drug_forms(id) ON UPDATE CASCADE ON DELETE CASCADE
);



insert into drugs (id, name, name_inn)
values  ('f82435fe-cb3a-44aa-aeb3-376e6649b7d4', 'Майозайм', 'Алглюкозидаза альфа'),
        ('f80372fb-e735-4fbf-b19f-d66991efd93c', 'Стрензик', 'Асфотаза альфа'),
        ('e8891845-01a1-4463-8fb7-47bb39fbc486', 'Трансларна', 'Аталурен'),
        ('b9f18a00-d3b7-4a60-b60a-26f17c5bb6e7', 'Цистадан', 'Бетаина ангидрид'),
        ('25b0d921-4603-4c1f-9add-252a91a939a7', 'Крисвита', 'Буросумаб'),
        ('60bdd153-fea0-45fe-896b-1ebad3e1e091', 'Ламзид', 'Велманаза альфа'),
        ('f1b9378f-d0fb-4466-a6d8-06aa85c4ebaf', '', 'Вилтоларсен'),
        ('0b6831f0-d8c7-4647-9555-dec251f78f91', '', 'Возоритид'),
        ('064a36c3-9482-4ea7-b983-cc024478c531', '', 'Воретиген непарвовек'),
        ('632d328e-9a91-4cb1-ab83-9fd2887e8294', 'Равикти', 'Глицерола фенилбутират'),
        ('4c03bcde-00f5-4632-b7a4-8c7ac8822200', '', 'Голодирсен'),
        ('db1317d6-b4fd-4a2f-9ff1-11d08ffee168', '', 'Динутуксимаб бета'),
        ('eb7e75f4-d139-453d-b746-9e3ba1addd7c', 'Оркамби', 'Ивакафтор+Лумакафтор'),
        ('f6175ab0-0aba-44c5-9228-a508ab00372e', 'Привиджен', 'Иммуноглобулин человека нормальный '),
        ('65f95d31-9fd0-4313-9be9-2cbacddb2844', 'Хайцентра', 'Иммуноглобулин человека нормальный (для подкожного введения)'),
        ('06b1ff22-e2d7-4325-b6e7-c4c3c2da21a6', 'Иларис', 'Канакинумаб'),
        ('e33152a8-0a3c-40ee-8b19-6ebeeeef2404', 'Карбаглу, Уцедан', 'Карглумовая кислота'),
        ('d156f64e-31ce-40ff-800b-61dabbf67061', 'Такзайро', 'Ланаделумаб'),
        ('860ef387-38fd-46ff-bae7-a41578305fe0', 'Витракви', 'Ларотректиниб'),
        ('380af9b5-e857-467a-89d2-37a95d224806', 'Лоджакста', 'Ломитапид'),
        ('f84163b7-51dd-4d56-b23d-c4c61664ab9d', 'Окслумо', 'Лумасиран'),
        ('5ed14ee4-4742-4c13-94c2-4a33f520bba1', '', 'Мараликсибат'),
        ('9919995d-9a71-4742-93f5-62dfe5958ccc', '', 'медицинского изделия стимулятор диафрагмального (френического) нерва стимулятор Mark IV с принадлежностями'),
        ('99e61acb-fb16-451e-a68d-2b2144a61412', 'Цистагон', 'Меркаптамина битартрат'),
        ('79f1f9cc-e4a4-4451-b16f-53571c949d4c', 'Майалепта', 'Метрелептин'),
        ('dd6a86f9-332a-49f4-be8b-3a88f9b11b07', 'Коргард, Солгол, Надолол', 'Надолол'),
        ('7ac5f502-39f8-4fe2-ad99-cf19a34f4855', 'Спинраза', 'Нусинерсен'),
        ('3bbeb86d-71b1-49b0-951c-31d990e7f52b', 'Золгенсма', 'Онасемноген абепарвовек'),
        ('3b8dc6ec-760c-460e-93e5-d29c5e8159db', null, 'Привиджен'),
        ('b3a31bac-05eb-46a5-85f2-f8327b570d3c', 'Цистагон', 'Пэгвалиаза'),
        ('e1197846-a3b6-4f21-8a78-df5747a69ca2', 'Эврисди', 'Рисдиплам'),
        ('c0f2abab-74ca-4048-8936-3f6f0f4e5477', 'Канума', 'Себелипаза альфа'),
        ('41b431c8-2e76-414c-930a-d14850d4d9b7', 'Апбрав', 'Селексипаг'),
        ('86f0da07-f153-4491-95c9-e0ccd09b2568', 'Коселуго', 'Селуметиниб'),
        ('cd3f3d13-5005-4bc1-900c-35d62a21ca9c', 'Гэттестив', 'Тедуглутид'),
        ('bb883c75-c0ee-4016-8afe-7e5a53b488df', '', 'Тригептаноин'),
        ('45d70072-d658-4b01-8016-2f524724d103', 'Финтепла', 'Фенфлурамин'),
        ('bde8046b-99e9-40d1-aa7f-ba6d61edf3e1', null, 'Хайцентра'),
        ('4f553382-d48a-40bd-bc5d-a0b982cb0d50', 'Орфакол, Колбам', 'Холевая кислота'),
        ('023f97aa-52c5-4288-abfb-4d6d26565698', 'Бринейра', 'Церлипоназа альфа'),
        ('e21c118e-fc23-429c-a70f-bdf0c7c515f3', 'Афинитор', 'Эверолимус'),
        ('f5d08b70-acbc-48f2-9bc0-9cc3c610f4c2', '', 'Элапегадемаза'),
        ('b9a1196f-687c-463b-a519-37cc446ad356', 'Трикафта', 'Элексакафтор/Тезакафтор/Ивакафтор + Ивакафтор'),
        ('9898293e-2489-4d06-a438-1a0661dc0294', 'Вимизайм', 'Элосульфаза Альфа'),
        ('d10fb2ff-6bae-4d15-b276-e0a50382807c', 'Розлитрек', 'Энтректиниб'),
        ('4bc258e8-4e68-48d4-9dc0-7828fbf74345', '', 'Этеплирсен');

insert into drug_forms (id, name, report_name, date_registration, drug_id)
values  ('acbe8139-cb97-46cf-a973-21487246e6a7', 'Гранулы', 'Оркамби, гранулы, 125мг+100мг,  №56', '2021-05-12', 'eb7e75f4-d139-453d-b746-9e3ba1addd7c'),
        ('285571b4-f63a-4f20-b3ec-fd12e4271eb3', 'Концентрат для внутривенных инфузий, 100 мг/2 мл', null, null, '4c03bcde-00f5-4632-b7a4-8c7ac8822200'),
        ('d2d46bd9-2126-481b-b705-4380d40437e7', 'Концентрат для приготовления раствора для внутривенного введения, 50 мг/мл', null, null, '4bc258e8-4e68-48d4-9dc0-7828fbf74345'),
        ('a82954df-d166-4cfb-8bf3-fcb0d0225f38', 'Концентрат для приготовления раствора для внутривенного введения, 50 мг/мл', null, null, 'f1b9378f-d0fb-4466-a6d8-06aa85c4ebaf'),
        ('28b5ed16-1ccf-4696-82bc-82035fb3fa6f', 'Концентрат для приготовления раствора для инфузий', 'Вимизайм конц. д/р-ра д/инф. 1 мг/мл 5 мл фл N 1x1 Веттер Фарма-Фертигунг ГмбХ и Ко. КГ/БиоМарин Интернэшнл Лимитед Германия', '2018-11-22', '9898293e-2489-4d06-a438-1a0661dc0294'),
        ('ebbf4047-1c1f-4660-994b-4d8ff6b0e4e0', 'Концентрат для приготовления раствора для инфузий', 'Себелипаза альфа (Канума) 2мг\мл 10 мл №1', '2017-10-31', 'c0f2abab-74ca-4048-8936-3f6f0f4e5477'),
        ('7f9e9eb9-87a4-4916-b446-2c34732484f8', 'Лиофилизат для приготовления концентрата для приготовления раствора для инфузий', 'Алглюкозидаза альфа (Майозайм)  лиофилизат для приготовления концентрата для приготовления раствора для инфузий 50 мг №1 (флаконы)', '2021-02-12', 'f82435fe-cb3a-44aa-aeb3-376e6649b7d4'),
        ('02805102-e617-4a5a-8a11-7ec548735591', 'Лиофилизат для приготовления раствора для подкожного введения', 'Канакинумаб (Иларис) лиофилизат для приготовления раствора для подкожного введения, 150 мг (флаконы)', '2012-01-11', '06b1ff22-e2d7-4325-b6e7-c4c3c2da21a6'),
        ('592ed401-1568-4528-a027-0e10cc130ae7', 'Лиофилизат для приготовления раствора для подкожного введения', 'Тедуглутид (Гэттестив) лиофилизат для приготовления раствора для подкожного введения 5 мг', '2021-06-24', 'cd3f3d13-5005-4bc1-900c-35d62a21ca9c'),
        ('b9f24ef4-3b22-44c1-a151-c5cc793a3c0a', 'Порошок для перорального приема', 'Бетаина ангидрид, порошок для перорального приема,   180г', null, 'b9f18a00-d3b7-4a60-b60a-26f17c5bb6e7'),
        ('c6c35b40-7d97-403c-91a2-1040b1c4f0a4', 'Порошок для приготовления раствора для приема внутрь', 'Эврисди пор. д/р-ра д/приема внутрь 0.75 мг/мл (60 мг) 2 г фл с адапт шп N 1x1 Ф. Хоффманн-Ля Рош Лтд. Швейцария', '2020-11-26', 'e1197846-a3b6-4f21-8a78-df5747a69ca2'),
        ('a376df4b-b3cf-40f2-b77f-d8a10ff2145d', 'Порошок для приема внутрь ', 'Трансларна порошок для приёма внутрь 1000мг пак.-саше 4000 мг №30', '2020-11-24', 'e8891845-01a1-4463-8fb7-47bb39fbc486'),
        ('44b97e94-2144-4c88-89c0-2cc5fbaf0a60', 'Раствор для интратекального введения', 'Спинраза р-р д/интратек. введения 2.4 мг/мл 5 мл фл N 1x1 Веттер Фарма-Фертигунг ГмбХ', '2019-08-16', '7ac5f502-39f8-4fe2-ad99-cf19a34f4855'),
        ('39511713-667b-4648-a4ee-857a8042e7d0', 'Раствор для инфузий', 'Онасемноген Абепарвовек (Золгенсма)', '2021-12-09', '3bbeb86d-71b1-49b0-951c-31d990e7f52b'),
        ('279f1b34-7915-4ff6-b90a-30f02b4e072b', 'Раствор для инфузий', null, '2014-05-06', 'f6175ab0-0aba-44c5-9228-a508ab00372e'),
        ('0b9e0520-a160-44cd-988b-1c5b9de820f4', 'Раствор для инфузий', null, null, '25b0d921-4603-4c1f-9add-252a91a939a7'),
        ('288f94a3-9a9d-4066-9b3d-02b374dbce8e', 'Раствор для инъекций, 2,4 мг/1,5 мл', null, null, 'f5d08b70-acbc-48f2-9bc0-9cc3c610f4c2'),
        ('d869a234-575c-4636-ace1-d92b5d97e930', 'Раствор для подкожного введения', null, '2019-07-19', 'f80372fb-e735-4fbf-b19f-d66991efd93c'),
        ('1d989b0e-84fd-40bc-87d2-c9085cec58b6', 'Раствор для подкожного введения', null, '2021-03-25', 'd156f64e-31ce-40ff-800b-61dabbf67061'),
        ('6519e1f1-940c-43e8-a5c1-e36ac9b16ba1', 'Раствор для подкожного введения', null, '2019-11-20', '65f95d31-9fd0-4313-9be9-2cbacddb2844'),
        ('6b0f5fcf-adaf-404a-b585-59da57ed9e2b', 'Раствор для подкожного введения', null, '2019-01-25', '06b1ff22-e2d7-4325-b6e7-c4c3c2da21a6'),
        ('0c21ab80-0aa7-47be-9e50-0fd350665dcc', 'Раствор для приема внутрь, флакон 30 мл', null, null, '5ed14ee4-4742-4c13-94c2-4a33f520bba1'),
        ('628ffbfa-6a13-4ae9-b9c7-d7f1aa40e4a2', 'Сироп', null, null, '632d328e-9a91-4cb1-ab83-9fd2887e8294'),
        ('c63abfc1-e94c-485e-829f-8419f6b8d045', 'Таблетки', 'МНН Эверолимус (Афинитор), таблетки диспергируемые, 2 мг (Блистер) 10*3 (пачка картонная)', null, 'e21c118e-fc23-429c-a70f-bdf0c7c515f3'),
        ('d3662a06-9379-4833-ba7b-24264c884377', 'Таблетки для перорального применения, 200мг № 60', 'Карглумовая кислота, таблетки для перорального применения, 200мг № 60', null, 'e33152a8-0a3c-40ee-8b19-6ebeeeef2404'),
        ('b830521a-0ad7-418f-88dd-6ffa42cc69f7', 'Таблетки, покрытые пленочной оболочкой', 'Трикафта (Элексакафтор/Тезакафтор/Ивакафтор + Ивакафтор, 100 мг/50мг/75мг + 150мг), таблетки, покрытые пленочной оболочкой, №84', null, 'b9a1196f-687c-463b-a519-37cc446ad356'),
        ('83a3f10f-871c-49fb-937a-29b4b4177662', 'Таблетки, покрытые пленочной оболочкой', 'Оркамби, таблетки, покрытые пленочной оболочкой, 125мг+100 мг, №112', '2020-12-02', 'eb7e75f4-d139-453d-b746-9e3ba1addd7c'),
        ('c8f40f93-7f02-4ba7-ad11-937765808125', 'Шприц для подкожных инъекций', 'Пэгвалиаза, шприц для подкожных инъекций, 20 мг/мл', null, 'b3a31bac-05eb-46a5-85f2-f8327b570d3c');

insert into drug_dozes (id, name,  drug_form_id)
values
    ('c8f40f93-7f02-4ba7-ad11-937765812025','50 мн', '7f9e9eb9-87a4-4916-b446-2c34732484f8'),
    ('c8f40f93-7f02-4ba7-ad11-937765800125','40 мг/мл' ,'d869a234-575c-4636-ace1-d92b5d97e930'),
    ('c8f40f93-7f02-4ba7-ad11-937765800225','100 мг/мл' ,'d869a234-575c-4636-ace1-d92b5d97e930'),
    ('c8f40f93-7f02-4ba7-ad11-937765800325','125 мг', 'a376df4b-b3cf-40f2-b77f-d8a10ff2145d'),
    ('c8f40f93-7f02-4ba7-ad11-937765800425','250 мг', 'a376df4b-b3cf-40f2-b77f-d8a10ff2145d'),
    ('c8f40f93-7f02-4ba7-ad11-937765800525','500 мг', 'a376df4b-b3cf-40f2-b77f-d8a10ff2145d'),
    ('c8f40f93-7f02-4ba7-ad11-937765800625','180 г','b9f24ef4-3b22-44c1-a151-c5cc793a3c0a'),
    ('c8f40f93-7f02-4ba7-ad11-937765800725','10 мг/мл  (1мг-1мл)','0b9e0520-a160-44cd-988b-1c5b9de820f4'),
    ('c8f40f93-7f02-4ba7-ad11-937765800825','50 мг/мл','a82954df-d166-4cfb-8bf3-fcb0d0225f38'),
    ('c8f40f93-7f02-4ba7-ad11-937765800925','1,6 мг/мл','a82954df-d166-4cfb-8bf3-fcb0d0225f38'),
    ('c8f40f93-7f02-4ba7-ad11-937765801125','1000 мг','628ffbfa-6a13-4ae9-b9c7-d7f1aa40e4a2'),
    ('c8f40f93-7f02-4ba7-ad11-937765801225','188 мг+150 мг № 56','83a3f10f-871c-49fb-937a-29b4b4177662'),
    ('c8f40f93-7f02-4ba7-ad11-937765801325','125 мг+100 мг № 112','83a3f10f-871c-49fb-937a-29b4b4177662'),
    ('c8f40f93-7f02-4ba7-ad11-937765801425','125 мг+100 мг № 56','acbe8139-cb97-46cf-a973-21487246e6a7'),
    ('c8f40f93-7f02-4ba7-ad11-937765801525','100 мг/2 мл','acbe8139-cb97-46cf-a973-21487246e6a7'),
    ('c8f40f93-7f02-4ba7-ad11-937765801625','150 мг/мл','6b0f5fcf-adaf-404a-b585-59da57ed9e2b'),
    ('c8f40f93-7f02-4ba7-ad11-937765801725','200 мг/мл','6b0f5fcf-adaf-404a-b585-59da57ed9e2b'),
    ('c8f40f93-7f02-4ba7-ad11-937765801825','150 мг/мл','02805102-e617-4a5a-8a11-7ec548735591'),
    ('c8f40f93-7f02-4ba7-ad11-937765801925','200 мг','d3662a06-9379-4833-ba7b-24264c884377'),
    ('c8f40f93-7f02-4ba7-ad11-937765802025','250 мг/5 мл','0c21ab80-0aa7-47be-9e50-0fd350665dcc'),
    ('c8f40f93-7f02-4ba7-ad11-937765802125','2.4 мг/мл','44b97e94-2144-4c88-89c0-2cc5fbaf0a60'),
    ('c8f40f93-7f02-4ba7-ad11-937765802325','2.4 мг/мл','39511713-667b-4648-a4ee-857a8042e7d0'),
    ('c8f40f93-7f02-4ba7-ad11-937765802425','2,2 мг','c8f40f93-7f02-4ba7-ad11-937765808125'),
    ('c8f40f93-7f02-4ba7-ad11-937765802525','10 мг','c8f40f93-7f02-4ba7-ad11-937765808125'),
    ('c8f40f93-7f02-4ba7-ad11-937765802625','20 мг','c8f40f93-7f02-4ba7-ad11-937765808125'),
    ('c8f40f93-7f02-4ba7-ad11-937765802725','2x10^13 вектор-геномов/мл','c63abfc1-e94c-485e-829f-8419f6b8d045'),
    ('c8f40f93-7f02-4ba7-ad11-937765802825','0.75 мг/мл','c6c35b40-7d97-403c-91a2-1040b1c4f0a4'),
    ('c8f40f93-7f02-4ba7-ad11-937765802925','10 мг','c63abfc1-e94c-485e-829f-8419f6b8d045'),
    ('c8f40f93-7f02-4ba7-ad11-937765803025','5 мг','c63abfc1-e94c-485e-829f-8419f6b8d045'),
    ('c8f40f93-7f02-4ba7-ad11-937765803125','125 мг+200 мг № 112','b830521a-0ad7-418f-88dd-6ffa42cc69f7'),
    ('c8f40f93-7f02-4ba7-ad11-937765803225','100 мг/50мг/75мг + 150мг','b830521a-0ad7-418f-88dd-6ffa42cc69f7'),
    ('c8f40f93-7f02-4ba7-ad11-937765803325','250 мг/5 мл','d2d46bd9-2126-481b-b705-4380d40437e7');



create table drug_recipes
(
    id          uuid      default uuid_generate_v4() not null primary key ,
    drug_id uuid REFERENCES drugs(id) ON UPDATE CASCADE ON DELETE CASCADE,
    drug_form_id uuid REFERENCES drug_forms(id) ON UPDATE CASCADE ON DELETE CASCADE,
    drug_doze_id uuid REFERENCES drug_dozes(id) ON UPDATE CASCADE ON DELETE CASCADE
);

alter table commissions add column
drug_recipe_id uuid REFERENCES drug_recipes(id) ON UPDATE CASCADE ON DELETE CASCADE;