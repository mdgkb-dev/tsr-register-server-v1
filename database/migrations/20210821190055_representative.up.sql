create table representative
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_2abe568eacaba9eba605bb231bc"
        primary key,
    human_id uuid
        constraint "REL_705a5ed0d568ad2cec915612ee"
        unique
        constraint "FK_705a5ed0d568ad2cec915612ee8"
        references human
        on update cascade on delete cascade,
    contact_phone varchar,
    contact_email varchar
);

alter table representative owner to mdgkb;

