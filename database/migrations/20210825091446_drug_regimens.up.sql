create table drug_regimens
(
    id uuid not null
        constraint drug_regimen_pk
        primary key,
    release_form_id uuid
        constraint drug_regimen_release_forms_id_fk
        references release_forms
        on update cascade,
    drug_id uuid
        constraint drug_regimen_drugs_id_fk
        references drugs
        on update cascade
);

create unique index drug_regimen_id_uindex
    on drug_regimens (id);

