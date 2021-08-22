create table document_field_value
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_caadab631597b6ad85f1d61f08c"
        primary key,
    value_string varchar,
    value_number integer,
    value_date date,
    document_id uuid not null
        constraint "FK_396b59db04e881d10bb1c315d10"
        references document
        on update cascade on delete cascade,
    document_type_field_id uuid not null
        constraint "FK_b2db3ddf07c32d6a031d009d86f"
        references document_type_fields,
    constraint "UQ_278c4f317fc1505fa2001957d32"
        unique (document_id, document_type_field_id)
);

alter table document_field_value owner to mdgkb;

