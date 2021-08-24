create table document_type_fields
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_504d1a17f1681be11d94673ba31"
        primary key,
    name varchar not null,
    "order" integer,
    type varchar not null,
    document_type_id uuid
        constraint "FK_040777158438fdb7a2ca0d9a3bd"
        references document_types
        on delete cascade
);

