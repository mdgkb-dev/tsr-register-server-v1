create table mkb_diagnosis
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_3c88e8dbef763c23da8ec58d36c"
        primary key,
    code varchar,
    name varchar,
    comment varchar,
    mkb_class_id uuid
        constraint "FK_80567c2f6c9fff2d263ca20bddc"
        references mkb_class
        on delete cascade,
    mkb_group_id uuid
        constraint "FK_d66713f2aa0d4b9b1d69a4b35e2"
        references mkb_group
        on delete cascade,
    mkb_sub_group_id uuid
        constraint "FK_2c077b0742aafd496ae48c4931d"
        references mkb_sub_group
        on delete cascade,
    mkb_sub_sub_group_id uuid
        constraint "FK_f864e86d0c89abb29479febdc4d"
        references mkb_sub_sub_group
        on delete cascade,
    leaf boolean default false not null,
    relevant boolean default true not null
);
