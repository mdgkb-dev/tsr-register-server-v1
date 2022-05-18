create table insurance_company_to_human
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_af71b25a1e3c78d42ef69648bfe"
        primary key,
    number varchar not null,
    insurance_company_id uuid not null
        constraint "FK_7ab4f0e0b5f313de5aea8acd7e5"
        references insurance_companies,
    human_id uuid
        constraint "FK_9c53c83f8e65f73d388a801c2d3"
        references human
);

create index "IDX_7ab4f0e0b5f313de5aea8acd7e"
    on insurance_company_to_human (insurance_company_id);

create index "IDX_9c53c83f8e65f73d388a801c2d"
    on insurance_company_to_human (human_id);

