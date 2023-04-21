create table researches_pools_researches
(
    id   uuid default uuid_generate_v4() not null primary key,
    researches_pool_id uuid not null references researches_pools,
    research_id  uuid not null references researches,
    item_order int not null default 0
);

