create table anthropometry
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_671f6f70cea0d7fd71db30be456"
        primary key,
    name varchar not null,
    measure varchar
);


