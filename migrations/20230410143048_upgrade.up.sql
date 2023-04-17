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


insert into mkb_items (name)
values ('МКБ10');
insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select mc.id, number, mc.name, mc.comment, mc.range_start, mc.range_end, mi.id from mkb_class mc
                                                                                        join mkb_items mi on mi.name = 'МКБ10';


insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, '-', m.name, m.comment, m.range_start, m.range_end, m1.id from mkb_groups m
                                                                                join mkb_class m1 on m1.id = m.mkb_class_id;

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

insert into mkb_items (id, code, name, comment, range_start, range_end, parent_id)
select m.id, concat_ws('.'::text, m1.code, m.sub_code) AS full_name, m.name, m.comment, '-', '-', m.mkb_diagnosis_id
from mkb_sub_diagnosis m
         join mkb_diagnosis m1 on m1.id = m.mkb_diagnosis_id;

alter table patient_diagnosis add column
    mkb_item_id uuid REFERENCES mkb_items(id) ON UPDATE CASCADE ON DELETE CASCADE;


update patient_diagnosis
set mkb_item_id = coalesce(mkb_sub_diagnosis_id, mkb_diagnosis_id, null)
where id is not null;

alter table users
    add email varchar;

create unique index users_email_uindex
    on users (email);

alter table users
    add uuid uuid not null default uuid_generate_v4();

