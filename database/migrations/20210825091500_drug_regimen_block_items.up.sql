create table drug_regimen_block_items
(
    id uuid not null
        constraint regimen_item_pk
        primary key,
    drug_regimen_block_id uuid
        constraint drug_regimen_block_items_drug_regimen_block_id_fk
        references drug_regimen_blocks
        on delete cascade,
    days_count integer,
    order_item integer
);

create unique index regimen_item_id_uindex
    on drug_regimen_block_items (id);

