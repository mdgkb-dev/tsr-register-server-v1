insert into drug_doze_components ("id","name","code","measure","quantity","drug_doze_id")
VALUES ('8930da44-0984-480d-aca9-5fb87d225c4b','Нусинерсен','nusinersen','мг.',2.4,'c8f40f93-7f02-4ba7-ad11-937765802125');

INSERT INTO drug_regimens(id, name, drug_doze_id)
values ('4519b19c-f7f7-42f4-922e-cc0b7c4ef33d','Режим дозирования препарата Спинраза','c8f40f93-7f02-4ba7-ad11-937765802125');

INSERT INTO drug_regimens(id, name, drug_doze_id)
values ('4519b19c-f7f7-42f4-922e-cc0b7c4ef33d','Режим дозирования препарата Спинраза','c8f40f93-7f02-4ba7-ad11-937765802125');


ALTER TABLE drug_regimen_blocks add column times_per_day integer;
ALTER TABLE drug_regimen_blocks add column days_count integer;
ALTER TABLE drug_regimen_blocks add column infinitely boolean;
ALTER TABLE drug_regimen_blocks add column every_day boolean;



INSERT INTO formulas (id,name,formula)
VALUES ('782ae466-90db-4b53-96e4-97cee8a6b97d','Режим дозирования препарата Спинраза','12');



INSERT INTO drug_regimen_blocks(id, drug_regimen_id, infinitely, order_item, every_day, times_per_day, days_count, formula_id)
VALUES 
('f1280302-c48a-4fe9-b099-0472bb1ef740', '4519b19c-f7f7-42f4-922e-cc0b7c4ef33d', false, 0, false, 1, 14, '782ae466-90db-4b53-96e4-97cee8a6b97d'), 
('f1280302-c48a-4fe9-b099-0472bb1ef741', '4519b19c-f7f7-42f4-922e-cc0b7c4ef33d', false, 1, false, 1, 14, '782ae466-90db-4b53-96e4-97cee8a6b97d'), 
('f1280302-c48a-4fe9-b099-0472bb1ef742', '4519b19c-f7f7-42f4-922e-cc0b7c4ef33d', false, 2, false, 1, 35, '782ae466-90db-4b53-96e4-97cee8a6b97d'), 
('f1280302-c48a-4fe9-b099-0472bb1ef743', '4519b19c-f7f7-42f4-922e-cc0b7c4ef33d', true, 3, false,  1, 120, '782ae466-90db-4b53-96e4-97cee8a6b97d');


update drug_regimen_blocks
set every_day = true, times_per_day = 1, days_count = 1, infinitely = true
WHERE id in ('f1280302-c48a-4fe9-b099-0472bb1ef736', 'f1280302-c48a-4fe9-b099-0472bb1ef737', 'f1280302-c48a-4fe9-b099-0472bb1ef738');


drop table drug_regimen_block_items;