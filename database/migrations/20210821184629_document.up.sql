create table document
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_e57d3357f83f3cdc0acffc3d777"
        primary key,
    document_type_id uuid not null
        constraint "FK_6b439665ef703bf850df3f12134"
        references document_types,
    human_id uuid not null
        constraint "FK_31f5ed20255452b13bfbd9208ae"
        references human
        on update cascade on delete cascade
);
