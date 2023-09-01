alter table document_types
add column required bool default false;

alter table document_types
    add column code varchar;


update document_types
set required = true where id = '69b70665-51b6-48df-ac07-1c1b78673c66';

alter table document_type_fields
    add column required bool default false;

alter table document_type_fields
    add column code varchar;

update document_type_fields
set required = true where id = 'd9e6ad60-8622-407d-9413-38b58eaa9a1b';

update document_type_fields
set code = 'number' where id = 'd9e6ad60-8622-407d-9413-38b58eaa9a1b';

update document_types
set code = 'snils' where id =  '69b70665-51b6-48df-ac07-1c1b78673c66';


alter table document_field_value rename to document_field_values;

UPDATE public.users_accounts SET password = '$2a$10$r3AlsAUhJZASy028tu6GcuitOp6vkL9xKszP1FV/RQ2ZdNagKHKW6' WHERE id = '86777706-b981-4aa9-99ab-652fe67bbeb9';