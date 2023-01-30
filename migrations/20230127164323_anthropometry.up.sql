create table chop_scale_questions
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    description varchar,
    item_order integer
);

create table chop_scale_question_scores
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    score integer,
    chop_scale_question_id uuid not null references chop_scale_questions
);

create table chop_scale_tests
(
    id uuid default uuid_generate_v4() not null primary key,
    item_date date,
    patient_id uuid not null references patients
);

create table chop_scale_test_results
(
    id uuid default uuid_generate_v4() not null primary key,
    item_date date,
    chop_scale_test_id uuid not null references chop_scale_tests on delete cascade,
    chop_scale_question_score_id uuid not null references chop_scale_question_scores
);


create table hmfse_scale_questions
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    description varchar,
    item_order integer
);

create table hmfse_scale_question_scores
(
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    score integer,
    hmfse_scale_question_id uuid not null references hmfse_scale_questions
);

create table hmfse_scale_tests
(
    id uuid default uuid_generate_v4() not null primary key,
    item_date date,
    patient_id uuid not null references patients
);

create table hmfse_scale_test_results
(
    id uuid default uuid_generate_v4() not null primary key,
    item_date date,
    hmfse_scale_test_id uuid not null references hmfse_scale_tests on delete cascade,
    hmfse_scale_questions_score_id uuid not null references hmfse_scale_question_scores
);

