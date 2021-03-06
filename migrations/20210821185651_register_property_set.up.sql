create table register_property_set
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_407419d79a09c3604e958dc21c0"
        primary key,
    name varchar not null,
    register_property_id uuid
        constraint "FK_8f6019b4369c0bc9eef2c65d5a1"
        references register_property on delete cascade,
    register_property_set_order int
);

create index "IDX_8f6019b4369c0bc9eef2c65d5a"
    on register_property_set (register_property_id);
