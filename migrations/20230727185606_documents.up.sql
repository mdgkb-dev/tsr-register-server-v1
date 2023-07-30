alter table document rename to documents;

alter table value_type rename to value_types;

alter table document_type_fields add column
value_type_id uuid references value_types;

update document_type_fields
set value_type_id = vt.id from value_types vt
where vt.name = document_type_fields.type::varchar;

alter table file_info_to_document rename to document_file_infos;

UPDATE users SET domain_id = '8669a507-5da0-4603-99b6-3e79e41f3f35' WHERE id = '66f6fe6b-4cbf-4bfa-96de-cdaadb4febbd'