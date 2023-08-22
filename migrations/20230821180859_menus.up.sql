CREATE TABLE menus
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name varchar,
    top boolean,
    hide boolean,
    side boolean,
    menu_order int not null default 0,
    link VARCHAR,
    icon_id uuid ,
    page_id uuid
);

insert into menus (id, name, top, hide, side, menu_order, link, icon_id, page_id)
values  ('408eae9d-043d-4fed-ae88-51f27c508ad1', 'Пациенты', null, null, null, 0, '/admin/patients', null, null),
        ('343a7237-f28b-4f16-99e2-4bdcd6b78f8b', 'Представители', null, null, null, 0, '/admin/representatives', null, null),
        ('94acb304-6651-4b6e-9646-11aeec3910a5', 'Врачебные комиссии/заявки', null, null, null, 0, '/admin/commission', null, null),
        ('e092080c-15c2-43ba-907e-bfd33e08c8eb', 'Заявки в ДЗМ/Фонд добра', null, null, null, 0, '/admin/drug-applications', null, null),
        ('ab0fa211-ba61-40d9-a48b-5724cfaa75ee', 'Склад', null, null, null, 0, '/admin/drug-arrives', null, null),
        ('d24a2240-684b-476b-96f9-ca5363f5b384', 'Лекарства', null, null, null, 0, '/admin/drugs', null, null);