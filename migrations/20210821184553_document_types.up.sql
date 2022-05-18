create table document_types
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_2e1aa55eac1947ddf3221506edb"
        primary key,
    name varchar not null
);
