ALTER TABLE drug_regimens 
add column months_range numrange;


ALTER TABLE drug_regimens 
add column weight_range numrange;


UPDATE drug_regimens
set months_range = '[2, 24)' WHERE id = '4519b19c-f7f7-42f4-922e-cc0b7c4ef33a';

UPDATE drug_regimens
set months_range = '[24,)' WHERE id = '4519b19c-f7f7-42f4-922e-cc0b7c4ef33b';

UPDATE drug_regimens
set months_range = '[24,)' WHERE id = '4519b19c-f7f7-42f4-922e-cc0b7c4ef33c';


UPDATE drug_regimens
set weight_range = '[0,20)' WHERE id = '4519b19c-f7f7-42f4-922e-cc0b7c4ef33b';

UPDATE drug_regimens
set weight_range = '[20,)' WHERE id = '4519b19c-f7f7-42f4-922e-cc0b7c4ef33c';

update drug_dozes set quantity = 2 WHERE id = 'c8f40f93-7f02-4ba7-ad11-937765802825';



CREATE table drug_needings(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    measures integer ,
    packs integer ,
    calculation varchar,
    weight integer,
    age_in_months integer,

    drug_regimen_id uuid references drug_regimens(id)
);



ALTER table commissions 
add column drug_needing_id uuid references drug_needings(id);