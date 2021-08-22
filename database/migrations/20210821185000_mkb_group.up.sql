create table mkb_group
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_82457efd3a8b530ddf31fdfb019"
        primary key,
    name varchar,
    range_start varchar,
    range_end varchar,
    comment varchar,
    mkb_class_id uuid
        constraint "FK_789445832fb09e3c0e72de814ca"
        references mkb_class
        on delete cascade,
    leaf boolean default false not null,
    relevant boolean default true not null
);

alter table mkb_group owner to mdgkb;

