create table register
(
    id uuid default uuid_generate_v4() not null
        constraint "PK_14473cc8f2caa81fd19f7648d54"
        primary key,
    name varchar not null
);
