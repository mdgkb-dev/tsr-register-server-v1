update answers
set patient_id = rr.patient_id
from research_results rr where rr.id = answers.research_result_id;

insert into  questions_filters  (question_id)
select id from questions
where research_id in (
    'c21eb5de-e180-4faa-b1eb-bfd77dddb470',
'e9f2300f-afb7-43e0-93b9-eb110edfa683',
'e9f2300f-afb7-43e0-93b9-eb110edfa689'
    );

insert into  questions_domains(question_id, domain_id)
select id, 'b9d7b8a5-d155-4dd5-8040-83c2648f0949' from questions
where research_id in (
    'c21eb5de-e180-4faa-b1eb-bfd77dddb470',
'e9f2300f-afb7-43e0-93b9-eb110edfa683',
'e9f2300f-afb7-43e0-93b9-eb110edfa689'
    );
