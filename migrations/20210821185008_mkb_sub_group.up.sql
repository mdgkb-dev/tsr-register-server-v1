create table mkb_sub_group
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_fb299527aa9a88060a122cefd41"
        primary key,
    name varchar,
    range_start varchar,
    range_end varchar,
    mkb_group_id uuid
        constraint "FK_1309484a375ecdf975815c5095c"
        references mkb_group
        on delete cascade,
    comment varchar,
    leaf boolean default false not null,
    relevant boolean default true not null
);


