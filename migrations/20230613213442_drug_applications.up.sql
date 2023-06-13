alter table drug_applications drop commission_id;
alter table drug_applications add column number varchar;


create table commissions_drug_applications
(
    id   uuid default uuid_generate_v4() not null primary key,
    commission_id uuid REFERENCES commissions(id) ON UPDATE CASCADE ON DELETE CASCADE,
    drug_application_id uuid REFERENCES drug_applications(id) ON UPDATE CASCADE ON DELETE CASCADE
);