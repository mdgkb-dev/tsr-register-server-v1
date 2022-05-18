create table register_property_to_user
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_e886c11085b8f42ae3483df9495"
        primary key,
    register_property_id uuid not null
        constraint "FK_ab3c8fcecaad7908dcd890f9440"
        references register_property,
    user_id uuid not null
        constraint "FK_7d243e83ece8913b37c6e85f915"
        references users
        on update cascade
);

create index "IDX_ab3c8fcecaad7908dcd890f944"
    on register_property_to_user (register_property_id);

create index "IDX_7d243e83ece8913b37c6e85f91"
    on register_property_to_user (user_id);
