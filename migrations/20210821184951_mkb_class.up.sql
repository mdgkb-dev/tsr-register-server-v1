create table mkb_class
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_2ee842782528f5c9a106e646ec5"
        primary key,
    number varchar,
    name varchar,
    range_start varchar,
    range_end varchar,
    comment varchar,
    leaf boolean default false not null,
    relevant boolean default true not null
);

