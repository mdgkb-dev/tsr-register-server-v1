create table questions_filters
(
    id uuid default uuid_generate_v4() not null primary key,
    question_id uuid not null references questions
);

create table questions_domains
(
    id uuid default uuid_generate_v4() not null primary key,
    question_id uuid not null references questions,
    domain_id uuid not null references domains
);


insert into questions_domains (id, question_id, domain_id)
values  ('4c5605e7-3942-4209-8c90-8835910488d7', 'fc171666-349d-4c79-86fc-cb6a2bfd2baf', 'b9d7b8a5-d155-4dd5-8040-83c2648f0949');


insert into questions_filters (id, question_id)
values  ('4c5605e7-3942-4209-8c90-8835910488d7', 'fc171666-349d-4c79-86fc-cb6a2bfd2baf');