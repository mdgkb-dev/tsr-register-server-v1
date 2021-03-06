    create table patients
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_8dfa510bb29ad31ab2139fbfb99"
        primary key,
    human_id uuid
        constraint "REL_87dd24bef72cee958bbdd799d5"
        unique
        constraint "FK_87dd24bef72cee958bbdd799d56"
        references human
);


create index "IDX_87dd24bef72cee958bbdd799d5"
    on patients (human_id);

