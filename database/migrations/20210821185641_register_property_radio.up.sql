create table register_property_radio
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_fc6c89a69e6669f3b60f4809bcd"
        primary key,
    name varchar not null,
    register_property_id uuid
        constraint "FK_4d4a161c891e6e45b3590c8a9d9"
        references register_property on delete cascade
);