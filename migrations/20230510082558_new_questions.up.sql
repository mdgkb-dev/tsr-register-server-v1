update questions q
set value_type_id = '6fd80c4c-ece3-479c-94b6-005ccebcfe73'
where q.name = 'Какому критерию анафилаксии (WAO 2020) соответствует данная системная реакция';

insert into selected_answer_variants (id, answer_variant_id,  research_result_id, answer_id)
select a.id, answer_variant_id, research_result_id, a.id from answers a
                                                                  join questions q on q.id = a.question_id
where q.name = 'Какому критерию анафилаксии (WAO 2020) соответствует данная системная реакция'
  and a.answer_variant_id not in ('8b818d9f-0ca6-47e7-a5de-32a8d66e5ddd', 'd31dfab8-9f24-45ff-8a5a-00e50ab188e8');

select * from questions q where q.name = 'Какому критерию анафилаксии (WAO 2020) соответствует данная системная реакция';

-- перенос последующей 1й критерий
insert into  selected_answer_variants (id, answer_variant_id, research_result_id, answer_id)
select a.id, 'ac066c62-ce57-4ae6-a85b-502ccff6df69',research_result_id, a.id from answers a
                                                                                      join questions q on a.question_id = q.id
where q.id = '9df2a216-ded6-45b5-84c1-d928cf86fd59' and answer_variant_id = '8b818d9f-0ca6-47e7-a5de-32a8d66e5ddd';

-- перенос первой 1й критерий
insert into  selected_answer_variants (id, answer_variant_id, research_result_id, answer_id)
select uuid_generate_v4(), 'aa6c886b-cbd1-4619-b5b2-e0526c6aef78',research_result_id, a.id from answers a
                                                                                                    join questions q on a.question_id = q.id
where q.id = '9df2a216-ded6-45b5-84c1-d928cf86fd59' and answer_variant_id = '8b818d9f-0ca6-47e7-a5de-32a8d66e5ddd';

-- перенос последующей 2й критерий
insert into  selected_answer_variants (id, answer_variant_id,  research_result_id, answer_id)
select uuid_generate_v4(), 'c7e8a1b0-e3d2-4784-a295-f1ae20846f9a',research_result_id, a.id from answers a
                                                                                                    join questions q on a.question_id = q.id
where q.id = '9df2a216-ded6-45b5-84c1-d928cf86fd59' and answer_variant_id = '8b818d9f-0ca6-47e7-a5de-32a8d66e5ddd';


-- перенос первой 2й критерий
insert into  selected_answer_variants (id, answer_variant_id,  research_result_id, answer_id)
select uuid_generate_v4(), '9691bd73-1fbc-4f08-b829-c706e9a8daec',research_result_id, a.id from answers a
                                                                                                    join questions q on a.question_id = q.id
where q.id = '9df2a216-ded6-45b5-84c1-d928cf86fd59' and answer_variant_id = '8b818d9f-0ca6-47e7-a5de-32a8d66e5ddd';


update questions set item_order = item_order + 1 where
        research_id in ('e9f2300f-afb7-43e0-93b9-eb110edfa688', 'e9f2300f-afb7-43e0-93b9-eb110edfa689')
                                                   and item_order > 15;

with ins_q as (
    insert into questions (name, item_order, value_type_id,  research_id)
        values
            ('Какому критерию анафилаксии (NIAID/FAAN 2005) соответствует данная системная реакция?', 16, '6fd80c4c-ece3-479c-94b6-005ccebcfe73','e9f2300f-afb7-43e0-93b9-eb110edfa688')
        returning id
)
insert into answer_variants (name, item_order, question_id, score, show_more_questions)
values
    ('1-му',0, (select id from ins_q), 0, false),
    ('2-му',1, (select id from ins_q), 0, false),
    ('3-му',2, (select id from ins_q), 0, false);

with ins_q as (
    insert into questions (name, item_order, value_type_id,  research_id)
        values
            ('Какому критерию анафилаксии (NIAID/FAAN 2005) соответствует данная системная реакция?', 16, '6fd80c4c-ece3-479c-94b6-005ccebcfe73','e9f2300f-afb7-43e0-93b9-eb110edfa689')
        returning id
)
insert into answer_variants (name, item_order, question_id, score, show_more_questions)
values
    ('1-му',0, (select id from ins_q), 0, false),
    ('2-му',1, (select id from ins_q), 0, false),
    ('3-му',2, (select id from ins_q), 0, false);

