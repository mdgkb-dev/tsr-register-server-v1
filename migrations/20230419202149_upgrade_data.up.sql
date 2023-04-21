insert into researches_pools
select * from register;

insert into researches_pools_researches (researches_pool_id, research_id, item_order)
select register_id, id, item_order from researches;
alter table researches drop column register_id;
alter table researches drop column item_order;


insert into patients_researches_pools (id, researches_pool_id, patient_id)
select id, register_id, patient_id from register_to_patient;





WITH r AS (
    insert into patients_researches (id, research_id, patient_id)
        select id, research_id, patient_id from research_results
        returning id, research_id, patient_id
)
update research_results
set patient_research_id = r.id
from r
where research_results.id = r.id;

