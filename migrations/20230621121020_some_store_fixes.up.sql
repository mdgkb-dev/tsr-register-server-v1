alter table drug_decreases
    alter column drug_arrive_id drop not null;

alter table drug_decreases
    drop constraint drug_decreases_drug_arrive_id_fkey;

