delete from answers where id in (
    select answer_id  from selected_answer_variants sav
                               join answer_variants av on sav.answer_variant_id = av.id
                               join questions q on av.question_id = q.id);


alter table selected_answer_variants
    add column  question_id uuid;


update selected_answer_variants set question_id = q.question_id
from (
         select sav.id, av.question_id, research_result_id, true  from selected_answer_variants sav
                                                                           join answer_variants av on sav.answer_variant_id = av.id
                                                                           join questions q on av.question_id = q.id
     ) as q where q.id = selected_answer_variants.id;


update selected_answer_variants sav set answer_id = q.answer_id from
    (select uuid_generate_v4() as answer_id, question_id, research_result_id  from selected_answer_variants sav
     group by question_id, research_result_id) q
where sav.question_id = q.question_id and sav.research_result_id = q.research_result_id;

insert into answers (id, question_id,  filled)
select  distinct sav.answer_id, sav.question_id, true  from selected_answer_variants sav
group by sav.answer_id, sav.question_id;

update answers set research_result_id = q.research_result_id from  (select research_result_id, answer_id, question_id from selected_answer_variants) q
where q.answer_id = answers.id and q.question_id = answers.question_id;

