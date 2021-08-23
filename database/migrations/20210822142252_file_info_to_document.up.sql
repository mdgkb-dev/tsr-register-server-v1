create table file_info_documents
(
    id uuid not null,
    file_info_id uuid
        constraint file_info_to_document_file_infos_id_fk
        references file_infos (id)
        on update cascade on delete cascade,
    document_id uuid
        constraint file_info_to_document_document_id_fk
        references document
        on update cascade on delete cascade
);

create unique index file_info_to_document_id_uindex
    on file_info_documents (id);

alter table file_info_documents
    add constraint file_info_to_document_pk
        primary key (id);

