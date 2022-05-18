create table insurance_companies
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_e07a506064998c73ab77800afd5"
        primary key,
    name varchar not null
);
