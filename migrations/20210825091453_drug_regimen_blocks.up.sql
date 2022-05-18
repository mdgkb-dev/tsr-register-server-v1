create table drug_regimen_blocks
(
    id uuid not null
        constraint drug_regimen_block_pk
        primary key,
    drug_regimen_id uuid
        constraint drug_regimen_block_drug_regimen_id_fk
        references drug_regimens
        on delete cascade,
    infinitely boolean,
    order_item integer
);

create unique index drug_regimen_block_id_uindex
    on drug_regimen_blocks (id);

