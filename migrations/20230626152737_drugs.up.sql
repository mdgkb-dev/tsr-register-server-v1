create table drug_dozes
(
    id   uuid default uuid_generate_v4() not null primary key,
    drug_id uuid REFERENCES drugs(id) ON UPDATE CASCADE ON DELETE CASCADE,
    name varchar
);