alter table drug_arrives
    alter column fund_contract_id drop not null;


create table statuses (
    id uuid default uuid_generate_v4() not null primary key,
    name varchar,
    color varchar,
    model varchar
);

insert into statuses
select *
from drug_applications_statuses;

update statuses
    set model = 'drugApplication' where id is not null;

alter table drug_applications drop column drug_application_status_id;
alter table drug_applications add column status_id uuid references statuses;

alter table commissions drop column commission_status_id;
alter table commissions add column status_id uuid references statuses;

alter table drug_arrives add column status_id uuid references statuses;

drop table drug_applications_statuses;
drop table commissions_statuses;
