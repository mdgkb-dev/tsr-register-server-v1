create table mkb_sub_sub_group
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_ed1344f24086558c8e446f5e9db"
        primary key,
    name varchar,
    range_start varchar,
    range_end varchar,
    comment varchar,
    mkb_sub_group_id uuid
        constraint "FK_277424dd6e1237bbb5ce20c4cc2"
        references mkb_sub_group
        on delete cascade,
    leaf boolean default false not null,
    relevant boolean default true not null
);
