create table users
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_cace4a159ff9f2512dd42373760"
        primary key,
    login varchar not null
        constraint "UQ_a62473490b3e4578fd683235c5e"
        unique,
    password varchar not null,
    region varchar,
    human_id uuid
        constraint "REL_456771a1f9cf5a06f562f7d753"
        unique
        constraint "FK_456771a1f9cf5a06f562f7d753a"
        references human
);
