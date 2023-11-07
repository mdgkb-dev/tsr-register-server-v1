
create table representatives_domains
(
    id          uuid default uuid_generate_v4() not null primary key,
    representative_id uuid not null references representatives,
    domain_id uuid not null references domains
);


insert into representatives_domains(representative_id, domain_id)
SELECT r.id, pd.domain_id from representatives r
join patients_representatives pr on pr.representative_id = r.id
join patients_domains pd on pd.patient_id= pr.patient_id;






