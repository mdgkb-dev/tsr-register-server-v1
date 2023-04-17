create table researches
(
    id   uuid default uuid_generate_v4() not null primary key,
    name varchar not null
);

insert into researches
select * from register;


create table patients_researches
(
    id          uuid default uuid_generate_v4() not null primary key,
    register_id uuid not null references register,
    patient_id  uuid not null references patients
);

insert into patients_researches
select id, register_id, patient_id from register_to_patient;


alter table register_group
    rename column register_id to research_id;

alter table register_group
    rename column register_group_order to item_order;

alter table register_group
drop constraint register_group_register_id_fkey;

alter table register_group
    rename to research_sections;

alter table research_sections
    add foreign key (research_id) references researches
        on delete cascade;


alter table register_property
    rename column register_property_order to item_order;

alter table register_property
    rename column register_group_id to research_section_id;

alter table register_property
    rename to questions;

alter table register_property_radio
    rename column register_property_radio_order to item_order;

alter table register_property_radio
    rename column register_property_id to question_id;

alter table register_property_radio
    rename to answers;

alter table register_property_examples
    rename to question_examples;

alter table question_examples
    rename column register_property_radio_order to answer_order;

alter table question_examples
    rename column register_property_id to question_id;

alter table question_examples
    rename column register_property_example_order to item_order;


alter table register_property_measures
    rename column register_property_measure_order to item_order;

alter table register_property_measures
    rename column register_property_id to question_id;

alter table register_property_measures
    rename to question_measures;

alter table register_groups_to_patients
    rename column register_groups_to_patients_date to item_date;

alter table register_groups_to_patients
    rename column register_group_id to research_section_id;

alter table register_groups_to_patients
    rename to patients_research_sections;


alter table register_property_others
    rename column register_property_others_order to item_order;

alter table register_property_others
    rename column register_property_id to question_id;

alter table register_property_others
    rename column register_property_radio_id to answer_id;

alter table register_property_others
    rename to answer_comments;


alter table register_property_to_patient
    rename column register_property_radio_id to answer_id;

alter table register_property_to_patient
    rename column register_property_id to question_id;

alter table register_property_to_patient
    rename column register_group_to_patient_id to patients_research_groups_id;

alter table register_property_to_patient
    rename column register_property_measure_id to question_measure_id;

alter table register_property_to_patient
    rename to patient_answers;

alter table patient_answers
    rename column register_property_variant_id to question_variant_id;

alter table patient_answers
    rename to patient_questions;

insert into answers (id, name, item_order, question_id)
select id, name, register_property_set_order, register_property_id from register_property_set;

alter table register_property_set_to_patient
    rename column register_property_set_id to answer_id;

alter table register_property_set_to_patient
    rename column register_group_to_patient_id to patients_research_sections_id;

alter table register_property_set_to_patient
drop constraint "FK_8281318758557dfc2a1fd67f090";

alter table register_property_set_to_patient
    add constraint "FK_8281318758557dfc2a1fd67f090"
        foreign key (answer_id) references answers;

alter table register_property_set_to_patient
    rename to patient_answers;


update answer_comments
set answer_id = register_property_set_id
where answer_id is null;

alter table answer_comments
drop column register_property_set_id;

drop table register_property_set;

drop table register_property_to_register_group;


alter table register_property_other_to_patient
    rename column register_property_other_id to answer_comment_id;

alter table register_property_other_to_patient
    rename column register_group_to_patient_id to patients_research_groups_id;

alter table register_property_other_to_patient
    rename to patient_answer_comments;



















