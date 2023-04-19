----------------
-- MKB --
----------------

create table mkb_items
(
    id                   uuid    default uuid_generate_v4() not null primary key,
    code                 varchar,
    name                 varchar,
    comment              varchar,
    range_start          varchar,
    range_end            varchar,
    parent_id            uuid,
    leaf                 boolean default false not null,
    relevant             boolean default true  not null,

    foreign key (parent_id) references mkb_items(id)
);

insert into mkb_items (name) values ('МКБ10');

insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select mc.id, number, mc.name, mc.comment, mc.range_start, mc.range_end, mi.id from mkb_class mc
                                                                                        join mkb_items mi on mi.name = 'МКБ10';
--
--
insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, '-', m.name, m.comment, m.range_start, m.range_end, m1.id from mkb_groups m
                                                                                join mkb_class m1 on m1.id = m.mkb_class_id;
--
insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, '-', m.name, m.comment, m.range_start, m.range_end, m1.id from mkb_sub_group m
                                                                                join mkb_groups m1 on m1.id = m.mkb_group_id;


insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, '-', m.name, m.comment, m.range_start, m.range_end, m1.id from mkb_sub_sub_group m
                                                                                join mkb_sub_group m1 on m1.id = m.mkb_sub_group_id;


insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, m.code, m.name, m.comment, '-', '-',
       coalesce(m.mkb_sub_sub_group_id, m.mkb_sub_group_id, m.mkb_group_id, m.mkb_class_id, (null)) as t
from mkb_diagnosis m;
--
insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, concat_ws('.'::text, m1.code, m.sub_code) AS full_name, m.name, m.comment, '-', '-', m.mkb_diagnosis_id
from mkb_sub_diagnosis m
         join mkb_diagnosis m1 on m1.id = m.mkb_diagnosis_id;
--
alter table patient_diagnosis add column mkb_item_id uuid REFERENCES mkb_items(id) ON UPDATE CASCADE ON DELETE CASCADE;


update patient_diagnosis set mkb_item_id = coalesce(mkb_sub_diagnosis_id, mkb_diagnosis_id, null) where id is not null;

----------
-- USER --
----------

alter table users add email varchar;
create unique index users_email_uindex on users (email);
alter table users add uuid uuid not null default uuid_generate_v4();


----------------
-- RESEARCHES --
----------------

create table researches_pools
(
    id   uuid default uuid_generate_v4() not null primary key,
    name varchar not null
);

create table patients_researches_pools
(
    id   uuid default uuid_generate_v4() not null primary key,
    researches_pool_id uuid not null references researches_pools,
    patient_id  uuid not null references patients
);

alter table register_group rename to researches;
alter table researches rename column register_group_order to item_order;
alter table researches drop constraint register_group_register_id_fkey;

create table patients_researches
(
    id          uuid default uuid_generate_v4() not null primary key,
    research_id uuid not null references researches,
    patient_id  uuid not null references patients,
    item_order int not null default 0
);

alter table register_property rename column register_property_order to item_order;
alter table register_property rename column register_group_id to research_section_id;
alter table register_property rename to questions;

alter table register_property_radio rename column register_property_radio_order to item_order;
alter table register_property_radio rename column register_property_id to question_id;
alter table register_property_radio rename to answers_variants;

alter table register_property_examples rename to question_examples;
alter table question_examples rename column register_property_id to question_id;
alter table question_examples rename column register_property_example_order to item_order;

alter table register_property_measures rename column register_property_measure_order to item_order;
alter table register_property_measures rename column register_property_id to question_id;
alter table register_property_measures rename to question_measures;

alter table register_groups_to_patients rename column register_groups_to_patients_date to item_date;
alter table register_groups_to_patients rename column register_group_id to research_id;
alter table register_groups_to_patients rename to research_results;

alter table register_property_others rename column register_property_others_order to item_order;
alter table register_property_others rename column register_property_id to question_id;
alter table register_property_others rename column register_property_radio_id to answer_variant_id;
alter table register_property_others rename to answer_comments;

alter table register_property_to_patient rename column register_property_radio_id to answer_variant_id;
alter table register_property_to_patient rename column register_property_id to question_id;
alter table register_property_to_patient rename column register_group_to_patient_id to research_result_id;
alter table register_property_to_patient rename column register_property_measure_id to question_measure_id;
alter table register_property_to_patient rename column register_property_variant_id to question_variant_id;
alter table register_property_to_patient rename to answers;

insert into answers_variants (id, name, item_order, question_id)
select id, name, register_property_set_order, register_property_id from register_property_set;

alter table register_property_set_to_patient rename column register_property_set_id to answer_variant_id;
alter table register_property_set_to_patient rename column register_group_to_patient_id to research_result_id;
alter table register_property_set_to_patient drop  constraint "FK_8281318758557dfc2a1fd67f090";
alter table register_property_set_to_patient add constraint "FK_8281318758557dfc2a1fd67f090" foreign key (answer_variant_id) references answers_variants ;
alter table register_property_set_to_patient rename to answer_answer_variants;

update answer_comments set answer_variant_id = register_property_set_id where answer_variant_id is null;
alter table answer_comments drop column register_property_set_id;

create table formulas
(
    id          uuid default uuid_generate_v4() not null primary key,
    name varchar,
    formula varchar,
    research_id uuid not null references researches
);

create table formula_results
(
    id          uuid default uuid_generate_v4() not null primary key,
    name varchar,
    formula_id uuid not null references formulas,
    low_range  numeric,
    high_range  numeric
);





-- alter table researches
--     rename column register_id to research_id;

-- insert into patients_researches
-- select id, register_id, patient_id from register_to_patient;




-- insert into researches
-- select * from register;

-- alter table researches
--     add foreign key (research_id) references researches
--         on delete cascade;




-- alter table question_examples
--     rename column register_property_radio_order to answer_order;

-- drop table register_property_set;
-- drop table register_property_to_register_group;

--
-- alter table register_property_other_to_patient rename column register_property_other_id to answer_comment_id;
-- alter table register_property_other_to_patient rename column register_group_to_patient_id to patients_research_groups_id;
-- alter table register_property_other_to_patient rename to patient_answer_comments;
