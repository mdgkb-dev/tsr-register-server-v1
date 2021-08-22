create table register_property_to_register_group
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_39e68b7b43041b49e8ec59c04e1"
        primary key,
    register_property_id uuid not null
        constraint "FK_545f0c4061aba481d84622a3a85"
        references register_property
        on update cascade on delete cascade,
    "order" integer,
    register_group_id uuid not null
        constraint "FK_4e2c97c03a043d9f53587012439"
        references register_group
        on update cascade on delete cascade
);

alter table register_property_to_register_group owner to mdgkb;

create index "IDX_545f0c4061aba481d84622a3a8"
    on register_property_to_register_group (register_property_id);

create index "IDX_4e2c97c03a043d9f5358701243"
    on register_property_to_register_group (register_group_id);

