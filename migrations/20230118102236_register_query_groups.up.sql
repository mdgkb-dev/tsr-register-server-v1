create table register_query_groups
(
    id uuid default uuid_generate_v4() not null primary key,
    register_query_id uuid references register_queries(id) on delete cascade,
    register_group_id uuid references register_group(id) on delete cascade,
    item_order integer default 0,
    aggregate_type varchar,
    count_sum boolean,
    count_percentage boolean
);


create table register_query_group_properties
(
    id uuid default uuid_generate_v4() not null primary key,
    register_query_group_id uuid references register_query_groups(id) on delete cascade,
    register_property_id uuid references register_group(id) on delete cascade,
    item_order integer default 0,
    aggregate_type varchar,
    every_radio_set bool
);
