do $$
begin
   if not exists (select 1 from pg_type where typname = 'register_queries_type_enum') then
      create type register_queries_type_enum as enum ('plain', 'group', 'aggregate');
   end if;
end $$;

create table register_queries
(
   id uuid default uuid_generate_v4() not null primary key,
   name varchar not null,
   type register_queries_type_enum not null,
   register_id uuid,

   foreign key (register_id) references register (id)
);

create table register_query_to_register_property
(
   id uuid default uuid_generate_v4() not null primary key,
   register_query_id uuid not null,
   register_property_id uuid not null,
   "order" integer,
   is_aggregate boolean default false not null,

   foreign key (register_query_id) references register_queries (id),
   foreign key (register_property_id) references register_property (id)
);