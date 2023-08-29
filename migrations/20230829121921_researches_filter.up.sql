create table researches_domains
(
    id          uuid default uuid_generate_v4() not null primary key,
    research_id uuid not null references researches,
    domain_id uuid not null references domains
);



