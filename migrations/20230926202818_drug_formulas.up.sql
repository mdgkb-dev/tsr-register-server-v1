alter table drug_regimens
add column max_months integer;

alter table drug_regimens
add column max_weight integer;

alter table drug_regimen_blocks
add column formula_id uuid not null references formulas;

create table
    drug_doze_components (
        id uuid default uuid_generate_v4 () not null primary key,
        name varchar,
        code varchar,
        measure varchar,
        quantity integer,
        drug_doze_id uuid not null references drug_dozes
    );

alter table drug_regimens
drop column drug_id;

alter table drug_regimens
add column drug_doze_id uuid references drug_dozes;

Insert into
    drug_doze_components (id, name, code, measure, quantity, drug_doze_id)
values
    (
        '8930da44-0984-480d-aca9-5fb87d225c4a',
        'Рисдиплам',
        'risdiplam',
        'мг.',
        60,
        'c8f40f93-7f02-4ba7-ad11-937765802825'
    );

alter table formulas
alter research_id
drop not null;

Insert into
    formulas (id, name, formula)
values
    (
        'b7b279cc-c929-4039-b1d7-02964283241b',
        'Режим дозировния препарата рисдиплам в возрасте от 2 месяцев до < 2 лет',
        '0.20*weight'
    ),
    (
        'cba0788f-7bc5-49dc-8a8b-466c53d788b7',
        'Режим дозировния препарата рисдиплам в возрасте от 2 лет (масса тела < 20кг)',
        '0.25*weight'
    ),
    (
        '782ae466-90db-4b53-96e4-97cee8a6b97c',
        'Режим дозировния препарата рисдиплам в возрасте от 2 лет (масса тела > 20кг)',
        '5'
    );

alter table drug_regimens
drop release_form_id;

insert into
    drug_regimens (id, name, drug_doze_id, max_months, max_weight)
values
    (
        '4519b19c-f7f7-42f4-922e-cc0b7c4ef33a',
        'Режим дозировния препарата рисдиплам в возрасте от 2 месяцев до < 2 лет',
        'c8f40f93-7f02-4ba7-ad11-937765802825',
        12,
        null
    ),
    (
        '4519b19c-f7f7-42f4-922e-cc0b7c4ef33b',
        'Режим дозировния препарата рисдиплам в возрасте от 2 лет (масса тела < 20кг)',
        'c8f40f93-7f02-4ba7-ad11-937765802825',
        null,
        20
    ),
    (
        '4519b19c-f7f7-42f4-922e-cc0b7c4ef33c',
        'Режим дозировния препарата рисдиплам в возрасте от 2 лет (масса тела > 20кг)',
        'c8f40f93-7f02-4ba7-ad11-937765802825',
        null,
        null
    );

alter table drug_regimen_blocks add every_day bool;

insert into
    drug_regimen_blocks (
        id,
        order_item,
        every_day,
        formula_id,
        drug_regimen_id
    )
VALUES
    (
        'f1280302-c48a-4fe9-b099-0472bb1ef736',
        0,
        true,
        'b7b279cc-c929-4039-b1d7-02964283241b',
        '4519b19c-f7f7-42f4-922e-cc0b7c4ef33a'
    ),
    (
        'f1280302-c48a-4fe9-b099-0472bb1ef737',
        0,
        true,
        'cba0788f-7bc5-49dc-8a8b-466c53d788b7',
        '4519b19c-f7f7-42f4-922e-cc0b7c4ef33b'
    ),
    (
        'f1280302-c48a-4fe9-b099-0472bb1ef738',
        0,
        true,
        '782ae466-90db-4b53-96e4-97cee8a6b97c',
        '4519b19c-f7f7-42f4-922e-cc0b7c4ef33c'
    );

insert into
    drug_regimen_block_items (
        id,
        times_per_day,
        days_count,
        drug_regimen_block_id
    )
VALUES
    (
        'f1280302-c48a-4fe9-b099-0472bb1ef730',
        1,
        0,
        'f1280302-c48a-4fe9-b099-0472bb1ef736'
    ),
    (
        'f1280302-c48a-4fe9-b099-0472bb1ef731',
        1,
        0,
        'f1280302-c48a-4fe9-b099-0472bb1ef737'
    ),
    (
        'f1280302-c48a-4fe9-b099-0472bb1ef732',
        1,
        0,
        'f1280302-c48a-4fe9-b099-0472bb1ef738'
    );


    update users_accounts set password = '$2a$10$r3AlsAUhJZASy028tu6GcuitOp6vkL9xKszP1FV/RQ2ZdNagKHKW6'
    where email = 'root@gmail.com';

    update menus set link = '/admin/commissions' where link = '/admin/commission';